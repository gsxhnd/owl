# Changelog

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

## v1.0.0

- [refactor] change sync pool

## v0.4.0

- [chore] change travis-ci to github action

## v0.3.0

- [fixed] fix etcd watch
- [feat] delete backend dir, instead use owl.go
- [feat] use owl.go as third-party lib
