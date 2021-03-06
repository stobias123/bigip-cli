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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	address  string
	username string
	password string
	output   string

	rootCmd = &cobra.Command{
		Use:   "bigip",
		Short: "The bigip cli is a binary for interacting with F5 Bigip Appliances",
		Long:  `This is a golang library for interacting with the bigip CLI.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bigip-cli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "The address of the BigIP appliance you'd like to connect to. BIGIP_ADDRESS")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "table", "BigIP output format (json,text)")
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "BigIP username. BIGIP_USERNAME")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "BigIP password. BIGIP_PASSWORD")

	// Viper Bindings
	viper.BindPFlag("address", rootCmd.PersistentFlags().Lookup("address"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))

	rootCmd.AddCommand(versionCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bigip-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bigip-cli")
	}

	viper.SetEnvPrefix("BIGIP")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	// This forces credential set.
	// We don't set persistent flags as required because of this...
	// https://github.com/spf13/viper/issues/397
	if _, err := Client(); err != nil {
		er("Set your creds.")
	}
	log.SetLevel(log.WarnLevel)
}

func er(msg interface{}) {
	log.Error(msg)
	rootCmd.Usage()
	os.Exit(1)
}
