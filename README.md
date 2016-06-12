# CHAOS

## Build

```
make
cd dist
./chaos
```

## Devel

### Env

```
go get -u github.com/nsf/gocode
go get -u golang.org/x/tools/cmd/goimports
go get -u github.com/alecthomas/gometalinter
go get -u github.com/kardianos/govendor
go get -u github.com/golang/lint/golint
go get -u github.com/rogpeppe/godef

go get -u github.com/itpkg/chaos
cd $GOPATH/src/github.com/itpkg/chaos
govendor sync
sudo npm install -g ember-cli@2.6
```

### Editor(Atom)

```
apm install seti-ui seti-syntax
```

#### Plugs

- go-plus
- git-plus
- react
- atom-beautify
- atom-typescript
- autosave: enabled is true

## Notes

### govendor

```
govendor init # Create the "vendor" folder and the "vendor.json" file.
govendor list # List and filter existing dependencies and packages.
govendor fetch golang.org/x/oauth2 # Like "go get" but copies dependencies into a "vendor" folder.
```

## Thanks

- <https://github.com/gin-gonic/gin>
- <https://github.com/jinzhu/gorm>
- <https://github.com/urfave/cli>
- <https://github.com/facebookgo/inject>
- <https://github.com/spf13/viper>
- <https://github.com/vmihailenco/msgpack>
