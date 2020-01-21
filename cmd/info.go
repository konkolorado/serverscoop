/*
Copyright © 2020 Uriel Mandujano <uriel.mandujano14@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// Options is a struct that will grow and contain all of this subcommand's
// settable options
type Options struct {
	outputFormat string
}

// Validate will check that the options provided to the subcommand are valid
func (opts *Options) Validate() error {
	if opts.outputFormat != "yaml" && opts.outputFormat != "json" {
		return errors.New("--output must be 'yaml' or 'json'")
	}
	return nil
}

var opts Options

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "get information on one or more servers",
	Long: `Return information on a server or a list of servers. In the event that
a server is not found, no data will be returned for that server.

Use with the -o flag to control the output's format

Examples:
serverscoop info server1
serverscoop info server1 server2 -o json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("info called with", opts.outputFormat)

		if err := opts.Validate(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "yaml",
		"Output format (yaml|json)")
}
