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
	"strings"

	"github.com/steenzout/go-log"
	"github.com/steenzout/go-log/fields"
	field_severity "github.com/steenzout/go-log/fields/severity"
	"github.com/steenzout/go-log/loggers/simple"

	"github.com/fatih/color"
)

// Custom struct for a custom formatter.
type Custom struct{}

// Format formats the log message.
func (f Custom) Format(msg *gol.LogMessage) (string, error) {
	lmsg := msg.FieldLength()
	buffer := make([]string, lmsg, lmsg)

	i := 0
	for k, v := range *msg {
		if k != fields.Severity && k != fields.Timestamp {
			buffer[i] = fmt.Sprintf("%s:'%s'", k, v)
			i += 1
		}
	}

	t, _ := msg.Timestamp()

	if severity, err := msg.Severity(); err != nil {
		return fmt.Sprintf("%s UNKNOWN %s\n", t.String(), strings.Join(buffer, " ")), nil
	} else {
		switch severity >= field_severity.Error {
		case true:
			return fmt.Sprintf("%s %s %s\n", t.String(), color.RedString("%s", severity), strings.Join(buffer, " ")), nil
		default:
			return fmt.Sprintf("%s %s %s\n", t.String(), severity, strings.Join(buffer, " ")), nil
		}
	}
}

var _ gol.LogFormatter = (*Custom)(nil)

var logger gol.Logger = simple.New(nil, &Custom{}, os.Stdout)

func main() {
	// this will be written to stderr
	logger.Send(gol.NewEmergency("message", "system is down"))
	logger.Send(gol.NewAlert("message", "failed to write to disk"))
	logger.Send(gol.NewCritical("message", "high server load"))
	logger.Send(gol.NewError("message", "invalid number format"))

	// this will not be written anywhere
	logger.Send(gol.NewWarning("message", "performance close to 1s threshold"))
	logger.Send(gol.NewNotice("message", "failed to communicate with monitoring service"))
	logger.Send(gol.NewInfo("message", "requested processed in 250ms"))
	logger.Send(gol.NewDebug("debug", "var x = 10"))
}
