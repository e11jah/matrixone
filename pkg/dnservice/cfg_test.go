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

package dnservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	c := &Config{}
	assert.Error(t, c.Validate())
	c.UUID = "dn1"
	assert.NoError(t, c.Validate())
	assert.Equal(t, defaultListenAddress, c.ListenAddress)
	assert.Equal(t, c.ListenAddress, defaultListenAddress)
	assert.Equal(t, c.ServiceAddress, defaultServiceAddress)
	assert.Equal(t, taeStorageBackend, c.Txn.Storage.Backend)
	assert.Equal(t, defaultZombieTimeout, c.Txn.ZombieTimeout.Duration)
	assert.Equal(t, defaultDiscoveryTimeout, c.HAKeeper.DiscoveryTimeout.Duration)
	assert.Equal(t, defaultHeatbeatDuration, c.HAKeeper.HeatbeatDuration.Duration)
	assert.Equal(t, defaultHeatbeatTimeout, c.HAKeeper.HeatbeatTimeout.Duration)
	assert.Equal(t, defaultConnectTimeout, c.LogService.ConnectTimeout.Duration)
}
