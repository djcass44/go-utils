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

// Package sliceutils provides useful functions for dealing with slices
// https://gobyexample.com/collection-functions
package sliceutils

import "golang.org/x/exp/constraints"

// Index returns the first index of t or -1 if no match
func Index[T constraints.Ordered](vs []T, t T) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Includes returns true if t is contained within vs
func Includes[T constraints.Ordered](vs []T, t T) bool {
	return Index(vs, t) >= 0
}

// Any returns true if one of the values in vs satisfies f
func Any[T any](vs []T, f func(T) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if every one of the strings in vs satisfies f
func All[T any](vs []T, f func(T) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all the strings in vs that satisfy f
func Filter[T any](vs []T, f func(T) bool) []T {
	vsf := make([]T, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying f to each value in vs
func Map[T any, K any](vs []T, f func(T) K) []K {
	vsm := make([]K, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
