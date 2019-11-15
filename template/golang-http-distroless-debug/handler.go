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
	"io/ioutil"
	"net/http"

	sdk "github.com/openfaas-incubator/go-function-sdk"
	log "github.com/sirupsen/logrus"

	"handler/function"
)

// setHeaders sets response headers from the provided 'http.Header' value.
func setHeaders(w http.ResponseWriter, src http.Header) {
	if src != nil {
		for k, v := range src {
			w.Header()[k] = v
		}
	}
}

// rootHandler handles requests to the root path.
func rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.Body != nil {
		defer req.Body.Close()
	}

	// Read the request body.
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Errorf("Failed to read the request's body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Call the handler.
	r, err := function.Handle(sdk.Request{
		Body:        b,
		Header:      req.Header,
		Host:        req.Host,
		Method:      req.Method,
		QueryString: req.URL.RawQuery,
	})
	if err != nil {
		log.Errorf("Failed to handle request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set response headers, status and body as appropriate, and send it back.
	setHeaders(w, r.Header)
	w.WriteHeader(r.StatusCode)
	if _, err := w.Write(r.Body); err != nil {
		log.Errorf("Failed to write response body: %v", err)
	}
}
