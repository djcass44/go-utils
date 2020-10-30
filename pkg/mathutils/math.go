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

package mathutils

// Min returns the smallest of the 2 given integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the largest of the 2 given integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinFloat64 returns the smallest of the 2 given float64's
func MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// MaxFloat64 returns the largest of the 2 given float64's
func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
