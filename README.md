# CHAOS(by go)

## Build

```
make
cd dist
./chaos
```

## Deploy
    mv dist /var/www/www.change-me.com
    cd /var/www/www.change-me.com
    ./chaos i
    vi config.toml # don't forget to change domain and database setting
    ./chaos db e # will print sql scripts to crete database and user
    ./chaos ng # will generate nginx.conf file

## Devel

### Env

```
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.4.3
gvm use go1.4.3
GOROOT_BOOTSTRAP=$GOROOT gvm install go1.7beta2
gvm use go1.7beta2 --default

go get -u github.com/nsf/gocode
go get -u golang.org/x/tools/cmd/goimports
go get -u github.com/alecthomas/gometalinter
go get -u github.com/golang/lint/golint
go get -u github.com/rogpeppe/godef

go get -u github.com/Masterminds/glide
cd $GOPATH/src/github.com/Masterminds/glide
make build
mv glide $GOPATH/bin

go get -u github.com/itpkg/chaos
cd $GOPATH/src/github.com/itpkg/chaos
glide install
cd front-react && npm install


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


## Thanks

- <https://github.com/gin-gonic/gin>
- <https://github.com/jinzhu/gorm>
- <https://github.com/urfave/cli>
- <https://github.com/facebookgo/inject>
- <https://github.com/spf13/viper>
- <https://github.com/vmihailenco/msgpack>
