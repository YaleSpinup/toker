# toker

A simple token management tool.

---
[![goreleaser](https://github.com/YaleSpinup/toker/actions/workflows/releaser.yml/badge.svg)](https://github.com/YaleSpinup/toker/actions/workflows/releaser.yml)
[![tests](https://github.com/YaleSpinup/toker/actions/workflows/tests.yaml/badge.svg)](https://github.com/YaleSpinup/toker/actions/workflows/tests.yaml)
## Usage

Note that for all commands below you will need to replace `toker` with `go run main.go` unless you've compiled the code into a binary named `toker`. You can get pre-compiled binaries from https://github.com/YaleSpinup/toker/releases.

```
A very basic cli tool for dealing with UUID tokens.

Usage:
  toker [command]

Available Commands:
  compare     Compare a password (UUID) with a bcrypt hash
  hash        Hash a password
  help        Help about any command
  new         Generate tokens

Flags:
      --config string   config file (default is $HOME/.toker.yaml)
  -h, --help            help for toker

Use "toker [command] --help" for more information about a command.
```

### Generate some example uuid tokens

#### Generate just UUIDs

`toker new`

#### Generate UUIDs and bcrypt hashes

`toker new -e`

### Hash a token

`toker hash 0fcb1d28-ef49-486b-8645-df9cbc943a04`

### Compare a token with a hash

`toker compare 0fcb1d28-ef49-486b-8645-df9cbc943a04 '$2a$10$mVRImM1rNCS0EVBDMNicyOSfFLLgOnv284evL.8NnM/M3Pm9ZwdlC'`

## Author

* E Camden Fisher <camden.fisher@yale.edu>
* Brandon Tassone <brandon.tassone@yale.edu>


## License

GNU Affero General Public License v3.0 (GNU AGPLv3)  
Copyright (c) 2019 Yale University
