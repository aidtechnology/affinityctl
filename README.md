# Affinityctl

[![Build Status](https://github.com/aidtechnology/affinityctl/workflows/ci/badge.svg?branch=master)](https://github.com/aidtechnology/affinityctl/actions)
[![Version](https://img.shields.io/github/tag/aidtechnology/affinityctl.svg)](https://github.com/aidtechnology/affinityctl/releases)
[![Software License](https://img.shields.io/badge/license-BSD3-red.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/aidtechnology/affinityctl?style=flat)](https://goreportcard.com/report/github.com/aidtechnology/affinityctl)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0-ff69b4.svg)](.github/CODE_OF_CONDUCT.md)

Affinity gateway integration tools.

## Javascript

Located in the `js` folder. Available for NodeJS and browser environments.

```javascript
affinity.GetMaterial().then((response) => {
  console.log("your new pin is: " + response.pin);
});
```

## GoLang client

Available on the `"github.com/aidtechnology/affinityctl/client"` package.

```go
sdk, _ := client.New(nil)
pin, _ := sdk.DID.GetMaterial()
log.Printf("your new pin is: %s", pin)
```

## CLI

This repository also provide the CLI-based `affinityctl` tool to
facilitate integration with the existing gateway services.

```text
Usage:
  affinityctl [command]

Available Commands:
  completion  Generate auto-completion scripts for common shells
  create      Create a new DID
  gateway     Start a gateway instance
  help        Help about any command
  info        Display information on a given identifier
  list        List identifiers generated
  resolve     Retrieve the DID document for a given identifier
  vc          Verifiable credential operations
  version     Show version information
```

The tool also provide basic commands for verifiable credentials
operations.

```text
Verifiable credential operations

Usage:
  affinityctl vc [command]

Available Commands:
  issue       Issue a new verifiable credential
  list        Authenticate and list stored credentials
  store       Store an existing VC in the user's wallet
  verify      Verify an existing credential

Flags:
  -h, --help   help for vc

Global Flags:
  --config string   config file

Use "affinityctl vc [command] --help" for more information about a command.
```
