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

// flags is a struct that will grow and contain all of this subcommand's
// settable flags
type flags struct {
	outputFormat string
}

type fact struct {
	Text string `yaml:"Text" json:"Text"`
}

type name struct {
	first string `json:"first"`
	last string `json:"last"`
}

type user struct {
	id string `json:"_id"`
	name name `json:"name"`
}

type catFactsList struct {
	Id string `json:"_id"`
	Text string `json:"text"`
	Type string `json:"type"`
	User user `json:"user"`
	Upvotes string `json:"upvotes"`
	UserUpvoted string `json:"userUpvoted"`
}

type catFactsAPIResponse struct {
	CatFactsList []catFactsList `json:"all"`
}

var opts flags

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "get information on one or more servers",
	Long: `Return information on a server or a list of servers. In the event that
a server is not found, no data will be returned for that server.

Use with the -o flag to control the output's format

Examples:
serverscoop info server1
serverscoop info server1 server2 -o json`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := opts.validate(); err != nil {
			return err
		}

		resp, err := http.Get("https://cat-fact.herokuapp.com/facts")
		if err != nil {
			return err
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var catFacts catFactsAPIResponse
		json.Unmarshal([]byte(bodyBytes), &catFacts)
		fact := randomFact(catFacts.CatFactsList)

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

// Validate will check that the options provided to the subcommand are valid
func (opts *flags) validate() error {
	if opts.outputFormat != "yaml" && opts.outputFormat != "json" {
		return errors.New("--output must be 'yaml' or 'json'")
	}
	return nil
}

// Return a random cat fact from the API response
func randomFact(list []catFactsList) fact {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randomIndex := rand.Intn(len(list))
	return fact{Text: list[randomIndex].Text}
}
