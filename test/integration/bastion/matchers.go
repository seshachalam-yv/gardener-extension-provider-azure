// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package bastion

import (
	"net/http"

	"github.com/Azure/go-autorest/autorest"
)

// IsNotFound ignore Azure Not Found Error
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	actual, ok := err.(autorest.DetailedError)
	if !ok {
		return false
	}

	if code, ok := actual.StatusCode.(int); !ok || code != http.StatusNotFound {
		return false
	}

	return true
}
