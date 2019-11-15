// Copyright 2019 Form3 Financial Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"net/http"
	"os"
	"time"
)

const (
	bindAddr           = ":8082"
	defaultTimeout     = 10 * time.Second
	readTimeoutEnvvar  = "READ_TIMEOUT"
	writeTimeoutEnvvar = "WRITE_TIMEOUT"
)

func main() {
	// Handle SIGINT and SIGTERM so we can gracefully shutdown.
	c := setupSignalHandler()

	// Create a new HTTP server with the specified (or default) read/write timeouts.
	m := http.NewServeMux()
	m.HandleFunc("/", rootHandler)
	s := &http.Server{
		Addr:         bindAddr,
		Handler:      m,
		ReadTimeout:  parseDurationOrDefault(os.Getenv(readTimeoutEnvvar), defaultTimeout),
		WriteTimeout: parseDurationOrDefault(os.Getenv(writeTimeoutEnvvar), defaultTimeout),
	}

	// Start the HTTP server.
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Wait for SIGINT or SIGTERM and gracefully shutdown the HTTP server.
	<-c
	ctx, fn := context.WithTimeout(context.Background(), defaultTimeout)
	defer fn()
	if err := s.Shutdown(ctx); err != nil {
		panic(err)
	}
}
