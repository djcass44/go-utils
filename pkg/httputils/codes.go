/*
 *    Copyright 2020 Django Cass
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

package httputils

import "net/http"

// IsHttpError returns true if a status code is 400 or above
func IsHttpError(code int) bool {
	return code > 399
}

// IsHttpSuccess returns true if a status code is of the 2xx range (200 - 299)
func IsHttpSuccess(code int) bool {
	return code >= http.StatusOK && code <= 299
}
