// Wiregost - Golang Exploitation Framework
// Copyright © 2020 Para
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package completers

import (
	"strings"

	"github.com/maxlandon/wiregost/client/commands"
)

// AutoCompleter is the autocompletion engine
type UserCompleter struct {
	Command *commands.Command
}

// Do is the completion function triggered at each line
func (hc *UserCompleter) Do(ctx *commands.ShellContext, line []rune, pos int) (options [][]rune, offset int) {

	// Complete command args
	splitLine := strings.Split(string(line), " ")
	line = trimSpaceLeft([]rune(splitLine[len(splitLine)-1]))

	for _, arg := range hc.Command.Args {
		search := arg.Name
		if !hasPrefix(line, []rune(search)) {
			sLine, sOffset := doInternal(line, pos, len(line), []rune(search+"="))
			options = append(options, sLine...)
			offset = sOffset
		} else {
			words := strings.Split(string(line), "=")
			argInput := lastString(words)

			// For some arguments, the split results in a last empty item.
			if words[len(words)-1] == "" {
				argInput = words[0]
			}

			// All boolean values
			if arg.Type == "boolean" {
				for _, search := range []string{"true ", "false "} {
					offset = 0
					if strings.HasPrefix(search, argInput) {
						options = append(options, []rune(search[len(argInput):]))
						offset = len(argInput)
					}
				}
				return
			}
		}
	}
	return options, offset
}
