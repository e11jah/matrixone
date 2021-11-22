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

package max

import (
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/aggregation"
	"github.com/matrixorigin/matrixone/pkg/vectorize/max"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func NewInt32(typ types.Type) *int32Max {
	return &int32Max{typ: typ}
}

func (a *int32Max) Reset() {
	a.v = 0
	a.cnt = 0
}

func (a *int32Max) Type() types.Type {
	return a.typ
}

func (a *int32Max) Dup() aggregation.Aggregation {
	return &int32Max{typ: a.typ}
}

func (a *int32Max) Fill(sels []int64, vec *vector.Vector) error {
	if n := len(sels); n > 0 {
		v := max.Int32MaxSels(vec.Col.([]int32), sels)
		if a.cnt == 0 || v > a.v {
			a.v = v
		}
		a.cnt += int64(n - vec.Nsp.FilterCount(sels))
	} else {
		v := max.Int32Max(vec.Col.([]int32))
		if a.cnt == 0 || v > a.v {
			a.v = v
		}
		a.cnt += int64(vec.Length() - vec.Nsp.Length())
	}
	return nil
}

func (a *int32Max) Eval() interface{} {
	if a.cnt == 0 {
		return nil
	}
	return a.v
}

func (a *int32Max) EvalCopy(proc *process.Process) (*vector.Vector, error) {
	data, err := proc.Alloc(4)
	if err != nil {
		return nil, err
	}
	vec := vector.New(a.typ)
	vs := encoding.DecodeInt32Slice(data[:4])
	vs[0] = a.v
	if a.cnt == 0 {
		vec.Nsp.Add(0)
	}
	vec.Col = vs
	vec.Data = data
	return vec, nil
}