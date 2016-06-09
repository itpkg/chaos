#!/bin/sh
export dst=dist

rm -r $dst front-emberjs/dist

go build -ldflags "-s" -o $dst/demo demo/main.go
upx -o $dst/chaos $dst/demo
rm -v $dst/demo
cp -rv demo/locales $dst

cd front-emberjs && ember build --env production && cd -
cp -rv front-emberjs/dist $dst/public

sed -i ':a;N;$!ba;s/>\s*</></g' $dst/public/index.html
