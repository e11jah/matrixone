// Copyright 2021 - 2022 Matrix Origin
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

package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileServices(t *testing.T) {
	fs := newFileServices(t, 3, 2)
	require.NotNil(t, fs.getDNLocalFileService(0))
	require.NotNil(t, fs.getDNLocalFileService(1))
	require.NotNil(t, fs.getDNLocalFileService(2))
	require.Nil(t, fs.getDNLocalFileService(3))
	require.NotNil(t, fs.getCNLocalFileService(0))
	require.NotNil(t, fs.getCNLocalFileService(1))
	require.Nil(t, fs.getCNLocalFileService(2))
}
