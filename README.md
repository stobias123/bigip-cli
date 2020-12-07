# BIGIP-CLI

This is a golang library for interacting with the bigip CLI.

## Usage:
```shell
  bigip [command]
```

## Available Commands:
```shell
  help           Help about any command
  version        Gets the version of bigip cli
  virtualservers Actions related to virtualservers
```

## Flags:
```shell
  -a, --address string    The address of the BigIP appliance you'd like to connect to.
      --config string     config file (default is $HOME/.bigip-cli.yaml)
  -h, --help              help for bigip
  -p, --password string   BigIP password
  -u, --username string   BigIP username
```