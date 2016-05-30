#!/bin/sh
export dst=dist

gradle clean
rm -r $dst front-emberjs/dist

mkdir -pv $dst
cd app && gradle build && cd -
cd front-emberjs && ember build --env production && cd -

cp -rv app/build/libs/itpkg-*.jar app/config $dst
cp -rv front-emberjs/dist $dst/public

sed -i ':a;N;$!ba;s/>\s*</></g' $dst/public/index.html







