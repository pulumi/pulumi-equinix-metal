[![Actions Status](https://github.com/pulumi/pulumi-equinix-metal/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-equinix-metal/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fequinix-metal.svg)](https://www.npmjs.com/package/@pulumi/equinix-metal)
[![Python version](https://badge.fury.io/py/pulumi-equinix-metal.svg)](https://pypi.org/project/pulumi-equinix-metal)
[![NuGet version](https://badge.fury.io/nu/pulumi.equinixmetal.svg)](https://badge.fury.io/nu/pulumi.equinixmetal)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-equinix-metal/sdk/v2/go)](https://pkg.go.dev/github.com/pulumi/pulumi-equinix-metal/sdk/v2/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-equinix-metal/blob/master/LICENSE)

# Equinix Metal provider

**PLEASE NOTE:** *This provider supercedes the Pulumi Packet provider. There is no direct migration between providers. To perform a migration,
please use the pulumi import command*

The Equinix Metal resource provider for Pulumi lets you use [Equinix Metal](https://metal.equinix.com/) resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).


## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/equinix-metal

or `yarn`:

    $ yarn add @pulumi/equinix-metal

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_equinix_metal

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-equinix-metal/sdk/v3

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.EquinixMetal

## Configuration

The following configuration points are available:

- `equinix-metal:authToken` - (Required) This is your Equinix Metal API Auth token. This can also be specified with the
  `PACKET_AUTH_TOKEN` or `METAL_AUTH_TOKEN` shell environment variable.

## Reference

For further information, please visit [the Equinix Metal provider docs](https://www.pulumi.com/docs/intro/cloud-providers/equinix-metal) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/equinix-metal).
