// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package network

import (
	"encoding/json"

	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

func mutateNetworkConfig(newObj, _ *extensionsv1alpha1.Network) error {
	extensionswebhook.LogMutation(logger, "Network", newObj.Namespace, newObj.Name)

	var (
		networkConfig map[string]interface{}
		//backendNone   = calicov1alpha1.None
		err error
	)

	if newObj.Spec.ProviderConfig != nil {
		networkConfig, err = decodeNetworkConfig(newObj.Spec.ProviderConfig)
		if err != nil {
			return err
		}
	} else {
		networkConfig = map[string]interface{}{"kind": ""}
	}

	networkConfig["backend"] = "none"
	modifiedJSON, err := json.Marshal(networkConfig)
	if err != nil {
		return err
	}

	newObj.Spec.ProviderConfig = &runtime.RawExtension{
		Raw: modifiedJSON,
	}

	return nil
}

func decodeNetworkConfig(network *runtime.RawExtension) (map[string]interface{}, error) {
	var networkConfig map[string]interface{}
	if network == nil || network.Raw == nil {
		return map[string]interface{}{}, nil
	}
	if err := json.Unmarshal(network.Raw, &networkConfig); err != nil {
		return nil, err
	}
	return networkConfig, nil
}
