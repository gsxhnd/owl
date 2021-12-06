[![GoDoc][1]][2]
[![license][3]][4]
![Test][5]
[![Coverage Status][6]][7]
[![Go Report Card][8]][9]
![Release][10]
![GitHub release (latest by date)][11]

[1]: https://godoc.org/github.com/gsxhnd/owl?status.svg
[2]: https://pkg.go.dev/github.com/gsxhnd/owl
[3]: https://img.shields.io/github/license/gsxhnd/owl
[4]: https://opensource.org/licenses/MIT
[5]: https://github.com/gsxhnd/owl/workflows/Test/badge.svg
[6]: https://coveralls.io/repos/github/gsxhnd/owl/badge.svg
[7]: https://coveralls.io/github/gsxhnd/owl
[8]: https://goreportcard.com/badge/github.com/gsxhnd/owl
[9]: https://goreportcard.com/report/github.com/gsxhnd/owl
[10]: https://github.com/gsxhnd/owl/workflows/Release/badge.svg
[11]: https://img.shields.io/github/v/release/gsxhnd/owl?label=version
[12]: https://github.com/spf13/viper

# owl

Go yaml configuration lib,can get config value from yaml file or etcd cluster like [viper][12] lib.

Owl also support binary cli, help you store yaml file's content into etcd cluster.

## 1. Install as cli

```shell
# macos
wget https://github.com/gsxhnd/owl/releases/latest/download/owl-darwin-amd64 -O owl
# linux x64
wget https://github.com/gsxhnd/owl/releases/latest/download/owl-linux-amd64 -O owl
# windows x64
wget https://github.com/gsxhnd/owl/releases/latest/download/owl-windows-amd64.exe -O owl.exe

chmod +x /usr/local/bin/owl

## show version
owl version

owl version:  1.5.1
owl commit:  6840df54b9a566acb78b36195fd9e826bb04d6cf
owl tree state:  clean
owl build date:  2020-12-29T15:49:14+0800
go version:  go1.15.6
go compiler:  gc
platform:  darwin/amd64
```

### 1.1 CLI Usage

```shell
NAME:
   owl - owl

USAGE:
   owl [global options] command [command options] [arguments...]

COMMANDS:
   get       get value by key
   get_keys  get keys by prefix
   put       read file then put value to etcd
   delete    delete value by key
   version   show version
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --endpoint value, -e value  (default: "http://127.0.0.1:2379")
   --help, -h                  show help (default: false)
```

## 2. Add as lib

```shell
go get -u github.com/gsxhnd/owl
```

### Putting Values into owl

#### Reading Config Files

Owl requires minimal configuration so it knows where to look for config files. Owl just supports YAML file.

Examples:

```go
owl.SetConfName("test.yaml")
owl.AddConfPath("./mock/")
err := owl.ReadConf()
if err != nil { // Handle errors reading the config file
  panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```

### Getting Values From

In owl, there are a few ways to get a value depending on the value’s type. The following functions and methods exist:

- `Get(key string) : interface{}`
- `GetBool(key string) : bool`
- `GetFloat64(key string) : float64`
- `GetInt(key string) : int`
- `GetIntSlice(key string) : []int`
- `GetString(key string) : string`
- `GetStringMap(key string) : map[string]interface{}`
- `GetStringMapString(key string) : map[string]string`
- `GetStringSlice(key string) : []string`
- `GetAll() : map[string]interface{}`
