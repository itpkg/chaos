#!/bin/sh

for s in postgresql redis-server nginx ssh
do
  /etc/init.d/$s start
done
#/etc/init.d/ssh start -D
cd /var/www/localhost && ./chaos s
