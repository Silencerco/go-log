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
	field "github.com/steenzout/go-log/fields/severity"
	filter "github.com/steenzout/go-log/filters/severity"
	"github.com/steenzout/go-log/formatters"
	"github.com/steenzout/go-log/loggers/simple"
)

var txtFmt = &formatters.Text{}
var errorLog gol.Logger = simple.New(filter.New(field.Error), txtFmt, os.Stderr)

func main() {
	// this will be written to stderr
	errorLog.Send(gol.NewEmergency("message", "system is down"))
	errorLog.Send(gol.NewAlert("message", "failed to write to disk"))
	errorLog.Send(gol.NewCritical("message", "high server load"))
	errorLog.Send(gol.NewError("message", "invalid number format"))

	// this will not be written anywhere
	errorLog.Send(gol.NewWarning("message", "performance close to 1s threshold"))
	errorLog.Send(gol.NewNotice("message", "failed to communicate with monitoring service"))
	errorLog.Send(gol.NewInfo("message", "requested processed in 250ms"))
	errorLog.Send(gol.NewDebug("debug", "var x = 10"))
}
