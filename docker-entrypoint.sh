#!/bin/sh

for s in postgresql redis-server nginx ssh
do
  /etc/init.d/$s start
done

cd /var/www/localhost
if [ ! -f config.toml ]
then
  ./chaos i -e production  
  ./chaos db n
  ./chaos db m
  ./chaos db s
  ./chaos ng
  ln -sfn etc/nginx/sites-enabled/localhost.conf /etc/nginx/sites-enabled/default
  nginx -s reload
fi
./chaos s
#/etc/init.d/ssh start -D
