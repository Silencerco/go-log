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
	"fmt"
	"os"

	"github.com/steenzout/go-log"
	"github.com/steenzout/go-log/formatters"

	field_severity "github.com/steenzout/go-log/fields/severity"
	filter_severity "github.com/steenzout/go-log/filters/severity"
	logger_simple "github.com/steenzout/go-log/loggers/simple"
	manager_simple "github.com/steenzout/go-log/managers/simple"
)

// Log holds the application LogManager instance.
var Log gol.LoggerManager

func init() {
	fmt.Println("init():start")
	Log = manager_simple.New()

	f := filter_severity.New(field_severity.Info)
	formatter := formatters.Text{}
	logger := logger_simple.New(f, formatter, os.Stdout)
	Log.Register("main", logger)

	channel := make(chan *gol.LogMessage, 10)
	Log.Run(channel)
	Log.Send(gol.NewInfo("message", "main.Log has been configured"))
	channel <- gol.NewInfo("message", "this message was sent directly to the log manager channel")

	fmt.Println("init():end")
}
