// Copyright 2021 Matrix Origin
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

package sum

import "github.com/matrixorigin/matrixone/pkg/container/types"

type intSum struct {
	cnt int64
	sum int64
	typ types.Type
}

type uintSum struct {
	cnt int64
	sum uint64
	typ types.Type
}

type floatSum struct {
	cnt int64
	sum float64
	typ types.Type
}

type sumCount struct {
	cnt int64
	sum float64
	typ types.Type
}

type intSumCount struct {
	cnt int64
	sum int64
	typ types.Type
}

type uintSumCount struct {
	cnt int64
	sum uint64
	typ types.Type
}

type floatSumCount struct {
	cnt int64
	sum float64
	typ types.Type
}