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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
 	"io/ioutil"
	"net/http"

 	"github.com/ghodss/yaml"
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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires at least 1 arg(s), only received %d",
												len(args))
		}
		if err := opts.Validate(); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		payload, err := json.Marshal(map[string]string{
		    "url": args[0],
		})
		if err != nil {
			return err
		}

		resp, err := http.Post("https://cleanuri.com/api/v1/shorten", "application/json",
													bytes.NewBuffer(payload))
		if err != nil {
			return err
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if opts.outputFormat == "json" {
			fmt.Println(string(bodyBytes))
		}
		if opts.outputFormat == "yaml" {
			respBytes := []byte(bodyBytes)
			y, _ := yaml.JSONToYAML(respBytes)
			fmt.Println(y)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "yaml",
		"Output format (yaml|json)")
}
