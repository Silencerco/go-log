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

package main

import (
	"os"

	"github.com/steenzout/go-log"
	"github.com/steenzout/go-log/formatters"
	"github.com/steenzout/go-log/loggers/simple"
)

var txtFmt = &formatters.Text{}
var log gol.Logger = simple.New(nil, txtFmt, os.Stdout)
var errorLog gol.Logger = simple.New(nil, txtFmt, os.Stderr)

func main() {
	log.Send(&gol.LogMessage{"message": "written to log"})

	errorLog.Send(&gol.LogMessage{"message": "written to error log"})
}
