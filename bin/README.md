# owl

## Install

### Binary release

### go install
```bash
go install github.com/gsxhnd/owl/bin/owl
```

## Usage
```bash
usage: owl COMMAND [arg...]

commands:
   get  retrieve the value of a key
   put  set the value of a key
   version show version

```

### get
```bash
usage: owl get [flags] [arg...]
flags:
   -e, --endpoint string   etcd endpoint (default "http://127.0.0.1")
arg:
   the key what you want value at the etcd
```

### put
```bash
usage: owl put [flags] [arg...]
example:
    owl put /conf/test.yaml ../mock/test.yaml
flags:
    -e, --endpoint string   etcd endpoint (default "http://127.0.0.1")
arg:
    the key what you want value at the etcd
```

## Example
### get
```bash
owl get -e 127.0.0.1:2379 /conf/test.yaml
```
### put
```bash
owl put -e local_dev:2379 /conf/test.yaml ./mock/test.yaml
```