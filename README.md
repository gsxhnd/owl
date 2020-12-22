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
[10]: https://github.com/gsxhnd/owl/workflows/Release/badge.svg?branch=master
[11]: https://img.shields.io/github/v/release/gsxhnd/owl?label=version


# owl

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
```

### 1.1 Cli Usage
```shell
NAME:
   owl - owl

USAGE:
   owl [global options] command [command options] [arguments...]

COMMANDS:
   get       get value by key
   get_keys  get keys by prefix
   put       read file then put value to etcd
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