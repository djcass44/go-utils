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

const (
	statusCodeGroupInfo      = 100
	statusCodeGroupSuccess   = 200
	statusCodeGroupRedirect  = 300
	statusCodeGroupClientErr = 400
	statusCodeGroupServerErr = 500
)

var (
	// IsHttpInformation returns true if a status code is of the 1xx range (100 - 199).
	IsHttpInformation = statusGroupEqual(statusCodeGroupInfo)

	// IsHttpSuccess returns true if a status code is of the 2xx range (200 - 299).
	IsHttpSuccess = statusGroupEqual(statusCodeGroupSuccess)

	// IsHttpRedirect returns true if a status code is of the 3xx range (300 - 399).
	IsHttpRedirect = statusGroupEqual(statusCodeGroupRedirect)

	// IsHttpClientError returns true if a status code is of the 4xx range (400 - 499).
	IsHttpClientError = statusGroupEqual(statusCodeGroupClientErr)

	// IsHttpServerError returns true if a status code is of the 5xx range (500 - 599).
	IsHttpServerError = statusGroupEqual(statusCodeGroupServerErr)
)

// IsHttpError returns true if a status code is 4xx or 5xx.
func IsHttpError(code int) bool {
	return IsHttpClientError(code) || IsHttpServerError(code)
}

// Returns a function that checks if a given status code belongs to a status code group.
func statusGroupEqual(groupCode int) func(int) bool {
	// statusCodeGroup returns a status code group (1xx, 2xx, ...) for a given status code.
	statusCodeGroup := func(code int) int {
		// Status code group is indicated by the first digit, ie. a result of integer division is enough.
		firstDigit := code / 100

		return firstDigit * 100
	}

	return func(code int) bool {
		return statusCodeGroup(code) == groupCode
	}
}
