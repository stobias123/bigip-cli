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

var (
	dataGroupName string

	getDataGroupCmd = &cobra.Command{
		Use: "get",
		//Aliases: []string{"ls", "l"},
		Short: "Get a specific Datagroup",
		Long:  ``,
		Run:   getDataGroup,
	}
)

func getDataGroup(cmd *cobra.Command, args []string) {
	log.Debugf("Retrieving DataGroup: %s", dataGroupName)
	client, err := Client()
	if err != nil {
		er("Problem in getDataGroup")
	}
	dg, err := client.GetInternalDataGroup(dataGroupName)
	if err != nil {
		er(err)
	}

	if output == "json" {
		dgJson, err := dg.MarshalJSON()
		if err != nil {
			er("[GetDataGroup] Problem w/ marshalJSON")
		}
		fmt.Printf("%s\n", dgJson)
	} else {
		fmt.Println(dg.FullPath)
		for _, record := range dg.Records {
			fmt.Printf("\t %s --> %s\n", record.Name, record.Data)
		}

	}
}

func init() {
	// Flags - Maybe split this later if we get to big.
	// Adding Commands
	getDataGroupCmd.Flags().StringVar(&dataGroupName, "name", "", "DataGroup name")
	cobra.MarkFlagRequired(getDataGroupCmd.Flags(), "name")
	datagroupCmd.AddCommand(getDataGroupCmd)
}
