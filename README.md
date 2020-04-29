[![Build Status](https://travis-ci.com/gertjvr/pulumi-bitbucket.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/gertjvr/pulumi-bitbucket)

# Bitbucket provider

The Bitbucket resource provider for Pulumi lets you use Bitbucket resources in your infrastructure
programs. To use this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/bitbucket

or `yarn`:

    $ yarn add @pulumi/bitbucket

### Python

To use from Python, install using `pip`:

    $ pip install pulumi-bitbucket

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-bitbucket/sdk/go/...

## Configuration

The following configuration points are available:

- `bitbucket:username` - (Required) Your username used to connect to bitbucket. You can also set this via the environment variable. `BITBUCKET_USERNAME`
- `bitbucket:password` - (Required) Your password used to connect to bitbucket. You can also set this via the environment variable. `BITBUCKET_PASSWORD`

## Reference

- NodeJS: https://www.pulumi.com/docs/reference/pkg/nodejs/pulumi/bitbucket/
- Python: https://www.pulumi.com/docs/reference/pkg/python/pulumi_bitbucket/
- Go: https://godoc.org/github.com/pulumi/pulumi-bitbucket/sdk/go/bitbucket
