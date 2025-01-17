// Copyright 2019-present Open Networking Foundation.
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

// Package utils implements various gNMI path manipulation facilities.
package utils

import (
	"fmt"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
)

// ToModelName simply joins together model name and version in a consistent way
func ToModelName(name string, version string) string {
	return fmt.Sprintf("%s-%s", name, version)
}

// ToConfigName simply joins together device ID and version in a consistent way
func ToConfigName(deviceID device.ID, version string) string {
	return fmt.Sprintf("%s-%s", deviceID, version)
}
