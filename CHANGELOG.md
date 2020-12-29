# Changelog

## v1.5.1
- [feat]: add `DeleteRemote` method
- [feat]: add `DeleteRemote` cmd
- [test]: add unit test

## v1.5.0
- [refactor]: merge utils.go
- [fix]: etcd dep error
- [fix]: make release ldflags error
- [doc]: update README.md

## v1.4.1
- [fix] ignore filepath set or not
- [fix] `exists` method

## v1.4.0
- [feat] add cmd-tool

## v1.3.0
- [feat] add `GetString`
- [feat] add `GetInt`
- [feat] add `GetFloat64`
- [feat] add `GetBool`
- [feat] add `GetStringSlice`
- [feat] add `GetIntSlice`
- [feat] add `GetStringMap`
- [feat] add `GetStringMapString`
- [feat] add `GetInt64`

## v1.2.0

- [chore] update etcd client version to v3.3.25
- [chore] add cast dep
- [fix] change `SetConfig` to `SetRemoteConfig`
- [fix] change `SetAddr` to `SetRemoteAddr`
- [fix] change `Put` to `PutRemote`
- [fix] change `Get` to `GetRemote`
- [fix] change `GetKeys` to `GetRemoteKeys`
- [feat] add `ReadConf`, `SetConfName` and `AddCOnfPath` for load local yaml file
- [feat] add `ReadInConf`
- [test] add owl test

## v1.1.0

- [feat] add get keys

