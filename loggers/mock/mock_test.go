//
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

package mock_test

import (
	"github.com/mediaFORGE/gol/loggers/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// LoggerTestSuite test suite for the github.com/mediaFORGE/gol/loggers/mock package.
type LoggerTestSuite struct {
	suite.Suite
}

func (s *LoggerTestSuite) TestNew() {
	l := mock.New()
	assert.NotNil(s.T(), l)
}
