// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var datagroupListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "Operations related to datagroup",
	Long:    ``,
	Run:     listDataGroup,
}

func listDataGroup(cmd *cobra.Command, args []string) {
	log.Debugf("Retrieving DataGroups: %s")
	client, err := Client()
	if err != nil {
		er("Problem in DataGroupList")
	}

	dgs, err := client.InternalDataGroups()
	if err != nil {
		er(err)
	}

	for _, datagroup := range dgs.DataGroups {
		fmt.Println(datagroup.FullPath)
	}
}
func init() {
	// Adding Commands
	datagroupCmd.AddCommand(datagroupListCmd)
}
