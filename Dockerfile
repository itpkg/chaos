FROM ubuntu
MAINTAINER Jitang Zheng <jitang.zheng@gmail.com>

ENV GOPATH /opt/go
ENV PATH $GOPATH/bin:$PATH
ENV PROJECT itpkg/chaos
WORKDIR /var/www

RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y ssh nginx redis-server postgresql \
  golang-go git npm nodejs \
  vim
RUN apt-get clean

RUN go get -d github.com/$PROJECT/demo
RUN cd $GOPATH/src/github.com/$PROJECT/demo \
  && go build -ldflags "-s" -o $WORKDIR/chaos main.go \
  && cp -r locales $WORKDIR/
RUN ln -s /usr/bin/nodejs /usr/bin/node
RUN cd $GOPATH/src/github.com/$PROJECT/front-react \
  && npm install \
  && npm run build && mv build $WORKDIR/public

RUN rm -r $GOPATH
#RUN for i in ssh redis postgresql; do systemctl enable $i; done

#RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/9.3/main/pg_hba.conf
#RUN echo "listen_addresses='*'" >> /etc/postgresql/9.3/main/postgresql.conf

EXPOSE 443 5432 6379
VOLUME ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql", "/var/www" ]
CMD ["/sbin/init"]
