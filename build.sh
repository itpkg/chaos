#!/bin/sh
export dst=dist

rm -r $dst front-emberjs/dist

go build -ldflags "-s" -o $dst/api demo/main.go
cp -rv demo/locales $dst

cd front-emberjs && ember build --env production && cd -
cp -rv front-emberjs/dist $dst/public

sed -i ':a;N;$!ba;s/>\s*</></g' $dst/public/index.html
