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
	"io"

	"github.com/stretchr/testify/mock"
)

// Writer mock implementation of the io.Writer interface.
type Writer struct {
	mock.Mock
}

// Write writes the given bytes.
// Returns the number of written bytes and error.
func (m *Writer) Write(p []byte) (n int, err error) {
	args := m.Mock.Called(p)

	n = args.Get(0).(int)

	if args.Get(1) == nil {
		err = nil
	} else {
		err = args.Get(1).(error)
	}
	return
}

var _ io.Writer = (*Writer)(nil)
