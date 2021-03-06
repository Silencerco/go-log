//
// Copyright 2016-2017 Pedro Salgado
// Copyright 2015 Rakuten Marketing LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package mock

import (
	"github.com/steenzout/go-log"

	"github.com/stretchr/testify/mock"
)

// LogFilter mock implementation of the gol.LogFilter interface.
type LogFilter struct {
	mock.Mock
}

// Filter performs a filter check on the given message.
// Returns whether or not a given message should be filtered.
func (m *LogFilter) Filter(msg *gol.LogMessage) bool {
	args := m.Mock.Called(msg)

	return args.Get(0).(bool)
}

var _ gol.LogFilter = (*LogFilter)(nil)
