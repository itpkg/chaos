# CHAOS(by go)

## Run(by docker)

```
docker build -t chaos .
docker run --rm -P -p 2222:22 -p 8080:80 chaos
firefox http://localhost:8080
ssh -p 2222 root@localhost # password is "root"
```

## Devel
### Install go
```
sudo pacman -S go go-tools
echo "GOPATH=/opt/go\nPATH=\$GOPATH/bin:\$PATH\nexport GOPATH PATH" >> ~/.zshrc
```
NEED RE-LOGIN FIRST
### Clone code

```
go get -u github.com/nsf/gocode
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
- atom-beautify
- atom-typescript
- language-docker
- linter-eslint
- language-vue
- react
- standard-formatter
- autosave: enabled is true

## Notes

### Docker

```
docker stop $(docker ps -l -q) # stop the most recent container
docker kill $(docker ps -q) # Kill all running containers
docker rm $(docker ps -a -q) # Delete all stopped containers (including data-only containers)
docker rmi $(docker images -q) # Delete ALL images
```

## Thanks

- <https://docs.docker.com/>
- <https://github.com/gin-gonic/gin>
- <https://github.com/jinzhu/gorm>
- <https://github.com/urfave/cli>
- <https://github.com/facebookgo/inject>
- <https://github.com/spf13/viper>
- <https://github.com/vmihailenco/msgpack>
