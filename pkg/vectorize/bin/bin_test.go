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

package bin

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/stretchr/testify/require"
)

func TestCountBitLenForInt(t *testing.T) {
	// count bits for unsigned int
	// eg: 0(0), 1(1), 2(10), 3(11)
	require.Equal(t, int64(1+1+2+2), Uint8BitLen([]uint8{0, 1, 2, 3}))
	require.Equal(t, int64(3+3), Uint16BitLen([]uint16{4, 5}))
	require.Equal(t, int64(3+3), Uint32BitLen([]uint32{6, 7}))
	require.Equal(t, int64(4+4), Uint64BitLen([]uint64{8, 9}))

	// count bits for signed int
	// eg: -1(0xffffffffffffffff, 64bits)
	// 127(1111111)
	// -128(1111111111111111111111111111111111111111111111111111111110000000, 64bits)
	require.Equal(t, int64(64+7+64), Int8BitLen([]int8{-1, 127, -128}))
	require.Equal(t, int64(5+5), Int16BitLen([]int16{17, 18}))
	require.Equal(t, int64(7+8), Int32BitLen([]int32{100, 200}))
	require.Equal(t, int64(64+64+24+30), Int64BitLen([]int64{-1e7, -1e9, 1e7, 1e9}))
}

func TestCountBitLenForFloat(t *testing.T) {
	// count bits for float
	// eg: 0.2(0), 1.8(1), 2.99(10), 3.14(11)
	require.Equal(t, int64(1+1+2+2), Float32BitLen([]float32{.2, 1.8, 2.99, 3.14}))
	require.Equal(t, int64(64+7+64), Float32BitLen([]float32{-1.99, 127.99, -128.89}))
	require.Equal(t, int64(7+8), Float32BitLen([]float32{100.99, 200.99}))

	require.Equal(t, int64(64+64+24+30), Float64BitLen([]float64{-1e7, -1e9, 1e7, 1e9}))
	// Phi=1.61...(1), E=2.7(10), Pi=3.14(11)
	require.Equal(t, int64(1+2+2), Float64BitLen([]float64{float64(math.Phi), float64(math.E), float64(math.Pi)}))
}

func TestUnsignedIntToBinary(t *testing.T) {
	cases := []uint64{0, 1, 2, 3}
	want := "011011"
	bytesNeed := Uint64BitLen(cases)
	ret := &types.Bytes{
		Data:    make([]byte, bytesNeed),
		Lengths: make([]uint32, len(cases)),
		Offsets: make([]uint32, len(cases)),
	}
	ret = Uint64ToBinary(cases, ret)

	require.Equal(t, uint32(1), ret.Lengths[0])
	require.Equal(t, uint32(0), ret.Offsets[0])

	require.Equal(t, uint32(1), ret.Lengths[1])
	require.Equal(t, uint32(1), ret.Offsets[1])

	require.Equal(t, uint32(2), ret.Lengths[2])
	require.Equal(t, uint32(1+1), ret.Offsets[2])

	require.Equal(t, uint32(2), ret.Lengths[3])
	require.Equal(t, uint32(1+1+2), ret.Offsets[3])

	require.Equal(t, []byte(want), ret.Data)
}

func TestIntToBinary(t *testing.T) {
	cases := []int64{-1, 127, -128}
	want := "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111110000000"
	bytesNeed := Int64BitLen(cases)
	ret := &types.Bytes{
		Data:    make([]byte, bytesNeed),
		Lengths: make([]uint32, len(cases)),
		Offsets: make([]uint32, len(cases)),
	}
	ret = Int64ToBinary(cases, ret)
	fmt.Println("bin:", strconv.FormatInt(cases[0], 2))
	require.Equal(t, uint32(64), ret.Lengths[0])
	require.Equal(t, uint32(0), ret.Offsets[0])

	require.Equal(t, uint32(7), ret.Lengths[1])
	require.Equal(t, uint32(64), ret.Offsets[1])

	require.Equal(t, uint32(64), ret.Lengths[2])
	require.Equal(t, uint32(64+7), ret.Offsets[2])

	require.Equal(t, []byte(want), ret.Data)
}

func TestFloatToBinary(t *testing.T) {
	cases := []float64{float64(math.Phi), float64(math.E), float64(math.Pi)}
	want := "11011"
	bytesNeed := Float64BitLen(cases)
	ret := &types.Bytes{
		Data:    make([]byte, bytesNeed),
		Lengths: make([]uint32, len(cases)),
		Offsets: make([]uint32, len(cases)),
	}
	ret = Float64ToBinary(cases, ret)

	require.Equal(t, uint32(1), ret.Lengths[0])
	require.Equal(t, uint32(0), ret.Offsets[0])

	require.Equal(t, uint32(2), ret.Lengths[1])
	require.Equal(t, uint32(1), ret.Offsets[1])

	require.Equal(t, uint32(2), ret.Lengths[2])
	require.Equal(t, uint32(1+2), ret.Offsets[2])

	require.Equal(t, []byte(want), ret.Data)
}
