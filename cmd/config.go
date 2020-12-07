package cmd

import (
	"fmt"

	"github.com/scottdware/go-bigip"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config represents the data required for creating a Bigip Client
type Config struct {
	Address        string
	Port           string
	Username       string
	Password       string
	LoginReference string
	ConfigOptions  *bigip.ConfigOptions
}

// Client returns a bigip client for use in REST operations
func Client() (*bigip.BigIP, error) {
	c := Config{
		Address:  fmt.Sprintf("%s", viper.Get("address")),
		Username: fmt.Sprintf("%s", viper.Get("username")),
		Password: fmt.Sprintf("%s", viper.Get("password")),
	}
	if c.Address != "" && c.Username != "" && c.Password != "" {
		//log.Infof("Initializing BigIP connection")
		var client *bigip.BigIP
		var err error
		if c.LoginReference != "" {
			client, err = bigip.NewTokenSession(c.Address, c.Username, c.Password, c.LoginReference, c.ConfigOptions)
			if err != nil {
				log.Errorf("[ERROR] Error creating New Token Session %s ", err)
				return nil, err
			}

		} else {
			client = bigip.NewSession(c.Address, c.Username, c.Password, c.ConfigOptions)
		}
		err = c.validateConnection(client)
		if err == nil {
			return client, nil
		}
		return nil, err
	}
	return nil, fmt.Errorf("BigIP provider requires address, username and password")
}

func (c *Config) validateConnection(client *bigip.BigIP) error {
	t, err := client.SelfIPs()
	if err != nil {
		log.Errorf("[ERROR] Connection to BigIP device could not have been validated: %v ", err)
		return err
	}

	if t == nil {
		log.Warnf("[WARN] Could not validate connection to BigIP")
		return nil
	}
	return nil
}
