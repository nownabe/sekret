Sekret
======

[![GitHub release](https://img.shields.io/github/release/nownabe/sekret.svg?style=popout)](https://github.com/nownabe/sekret/releases)
[![License](https://img.shields.io/github/license/nownabe/sekret.svg?style=popout)](https://github.com/nownabe/sekret/blob/master/LICENSE.txt)
[![Build Status](https://travis-ci.org/nownabe/sekret.svg?branch=master)](https://travis-ci.org/nownabe/sekret)
[![Go Report Card](https://goreportcard.com/badge/github.com/nownabe/sekret)](https://goreportcard.com/report/github.com/nownabe/sekret)
[![codecov](https://codecov.io/gh/nownabe/sekret/branch/master/graph/badge.svg)](https://codecov.io/gh/nownabe/sekret)

Sekret is a tool to edit encrypted Kubernetes Secrets YAML as plain text.

# Installation

```bash
go get github.com/nownabe/sekret
```

Or download binary from [GitHub releases](https://github.com/nownabe/sekret/releases)

# Usage

```bash
$ sekret --help
NAME:
   sekret - Work with encrypted Kubernetes Secrets

USAGE:
   sekret [global options] command [command options] [arguments...]

VERSION:
   1.0.1

COMMANDS:
     edit          Edit secrets as plain text
     encrypt, enc  Encrypt file
     decrypt, dec  Decrypt encrypted file
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --key value    Encryption key (16 or 32 bytes) [$ENCRYPTION_KEY]
   --help, -h     show help
   --version, -v  print the version
```

## Examples

### Encrypt and Decrypt

```bash
$ export ENCRYPTION_KEY=$YOUR_ENCRYPTION_KEY
$ sekret enc secret.yaml > secret.yaml.enc
$ file secret.yaml*
secret.yaml:     ASCII text
secret.yaml.enc: data
$ sekret dec secret.yaml.enc
apiVersion: v1
kind: Secret
metadata:
  namespace: my-namespace
  name: my-secret
data:
  apikey: dGhpcyBpcyBhcGkga2V5
```

# Development

## Release

```bash
tools/release 1.0.0
```
