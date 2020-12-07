# BIGIP-CLI

This is a golang library for interacting with the bigip CLI.

Please submit issues and feature requests.

## Usage:
```shell
  bigip [command]
```
**Example**
```shell
$ export BIGIP_ADDRESS=1.1.1.1
$ export BIGIP_USERNAME=foo
$ export BIGIP_PASSWORD=bar
$ bigip dg list
INFO[0000] Initializing BigIP connection                
/Common/ProxyPassvs_https
/Common/aol
/Common/images
/Common/private_net
```

## Available Commands:
```shell
  datagroup      Operations related to datagroup
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