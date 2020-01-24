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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
  "time"

  "gopkg.in/yaml.v2"
	"github.com/spf13/cobra"
)

// Options is a struct that will grow and contain all of this subcommand's
// settable options
type Options struct {
	outputFormat string
}

type Fact struct {
	Text string `yaml:"Text"`
}

type Name struct {
	First string `json:"first"`
	Last string `json:"last"`
}

type User struct {
	Id string `json:"_id"`
	Name Name `json:"name"`
}

type CatFactsList struct {
	Id string `json:"_id"`
	Text string `json:"text"`
	Type string `json:"type"`
	User User `json:"user"`
	Upvotes string `json:"upvotes"`
	UserUpvoted string `json:"userUpvoted"`
}

type CatFactsAPIResponse struct {
	CatFactsList []CatFactsList `json:"all"`
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
		resp, err := http.Get("https://cat-fact.herokuapp.com/facts")
		if err != nil {
			return err
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var catFacts CatFactsAPIResponse
		json.Unmarshal([]byte(bodyBytes), &catFacts)

		rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
		randomFactIndex := rand.Intn(len(catFacts.CatFactsList))
		fact := Fact{Text: catFacts.CatFactsList[randomFactIndex].Text}

		if opts.outputFormat == "json" {
			factJSON, err := json.Marshal(fact)
			if err != nil {
				return err
			}
			fmt.Println(string(factJSON))
		}
		if opts.outputFormat == "yaml" {
			factYAML, err := yaml.Marshal(fact)
			if err != nil {
				return err
			}
			fmt.Print(string(factYAML))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "yaml",
		"Output format (yaml|json)")
}
