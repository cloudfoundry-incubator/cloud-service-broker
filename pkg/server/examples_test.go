// Copyright 2019 the Service Broker Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoogleCloudPlatform/gcp-service-broker/pkg/client"
	"github.com/GoogleCloudPlatform/gcp-service-broker/pkg/providers/builtin"
)

func TestNewExampleHandler(t *testing.T) {
	registry := builtin.BuiltinBrokerRegistry()

	// Validate that the handler returns the correct Content-Type
	handler := NewExampleHandler(registry)
	request := httptest.NewRequest(http.MethodGet, "/examples", nil)
	w := httptest.NewRecorder()

	handler(w, request)

	if w.Code != http.StatusOK {
		t.Errorf("Expected response code: %d got: %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json content type got: %q", contentType)
	}

	body := w.Body.Bytes()
	var allExamples []client.CompleteServiceExample

	err := json.Unmarshal(body, &allExamples)
	if err != nil {
		t.Errorf("Error unmarshalling json data: %q", err)
	}

}
