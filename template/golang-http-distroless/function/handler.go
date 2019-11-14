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

package function

import (
	"net/http"

	sdk "github.com/openfaas-incubator/go-function-sdk"
)

// Handle handles the incoming request by responding with "200 OK" and the original bytes.
func Handle(req sdk.Request) (sdk.Response, error) {
	return sdk.Response{
		Body:       req.Body,
		StatusCode: http.StatusOK,
	}, nil
}
