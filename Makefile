dst=dist

build:
	make $(dst)/chaos
	make react




$(dst)/chaos:
	go build -ldflags "-s -X main.version=`git describe --long --tags`" -o $@.1 demo/main.go
	upx -o $@ $@.1
	-rm -v $@.1
	-cp -rv demo/locales $(dst)


react:
	cd front-react && npm run build
	-cp -rv front-react/build $(dst)/public

clean:
	-rm -rv $(dst)
