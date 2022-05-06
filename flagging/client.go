/*
 *    Copyright 2022 Django Cass
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package flagging

import "github.com/Unleash/unleash-client-go/v3"

var defaultClient *unleash.Client

// Initialize will specify the options to be used by the default client.
func Initialize(options ...unleash.ConfigOption) (err error) {
	defaultClient, err = unleash.NewClient(options...)
	return
}

// IsEnabled queries the default client whether the specified feature is enabled or not.
func IsEnabled(feature string, options ...unleash.FeatureOption) bool {
	if defaultClient == nil {
		return false
	}
	return unleash.IsEnabled(feature, options...)
}
