/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2018 Roland Singer [roland.singer@deserbit.com]
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package cmd

import (
	"fmt"

	"github.com/spencerdodd/grumble"
)

var App = grumble.New(&grumble.Config{
	Name:        "foo",
	Description: "test for command switching",
})
var CommandSet1 = []*grumble.Command{
	&grumble.Command{
		Name:      "command1",
		Help:      "command set 1 command",
	},
}
var CommandSet2 = []*grumble.Command{
	&grumble.Command{
		Name:      "command2",
		Help:      "command set 2 command",
	},
}


func setCommandSet1(a *grumble.App) {
	fmt.Printf("[*] switching to command set 1 (%v commands)\n", len(CommandSet1))
	a.RemoveCommand("command2")
	for i := 0; i < len(CommandSet1); i++ {
		App.AddCommand(CommandSet1[i])
	}
}

func setCommandSet2(a *grumble.App) {
	fmt.Printf("[*] switching to command set 2\n", len(CommandSet2))
	a.RemoveCommand("command1")
	for i := 0; i < len(CommandSet2); i++ {
		App.AddCommand(CommandSet2[i])
	}
}

func init() {
	switchCmd1 := &grumble.Command{
		Name:     "switch1",
		Help:     "switch to commmand set 1",
		Run: func(c *grumble.Context) error {
			setCommandSet1(App)
			return nil
		},
	}
	App.AddCommand(switchCmd1)

	switchCmd2 := &grumble.Command{
		Name:     "switch2",
		Help:     "switch to commmand set 2",
		Run: func(c *grumble.Context) error {
			setCommandSet2(App)
			return nil
		},
	}
	App.AddCommand(switchCmd2)
}
