FROM ubuntu
MAINTAINER Jitang Zheng <jitang.zheng@gmail.com>

ENV LANG en_US.UTF-8
ENV GOPATH /opt/go
ENV PATH $GOPATH/bin:$PATH
ENV PROJECT itpkg/chaos
ENV WWW_HOME /var/www/localhost

#LOCALES
RUN for l in en_US zh_CN zh_TW; do echo "$l.UTF-8 UTF-8" >> /etc/locale.gen; done
RUN locale-gen
RUN update-locale LANG=$LANG

RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y openssh-server nginx redis-server postgresql elasticsearch \
  golang-go git npm nodejs \
  vim net-tools bash-completion
RUN apt-get clean
RUN ln -s /usr/bin/nodejs /usr/bin/node

#CLONE CODE
RUN go get -d github.com/$PROJECT/demo
RUN cd $GOPATH/src/github.com/$PROJECT/front-react && npm install

#SSH
RUN echo 'root:root' | chpasswd
RUN sed -i 's/PermitRootLogin without-password/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed -i 's/prohibit-password/yes/' /etc/ssh/sshd_config
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd


#POSTGRESQL
WORKDIR /etc/postgresql/9.5/main
RUN echo "listen_addresses='*'" >> postgresql.conf
RUN sed -i 's/peer/trust/' pg_hba.conf
RUN sed -i 's/md5/trust/' pg_hba.conf

#DISABLE ipv6
RUN for i in all default lo; do echo "net.ipv6.conf.$i.disable_ipv6 = 1" >> /etc/sysctl.conf; done

#BUILD
WORKDIR $GOPATH/src/github.com/$PROJECT
RUN cd demo \
  && go build -ldflags "-s" -o $WWW_HOME/chaos main.go \
  && cp -r locales $WWW_HOME
RUN cd front-react \
  && npm run build \
  && mv build $WWW_HOME/public

EXPOSE 22 80 443 5432 6379
VOLUME "/etc/postgresql" "/var/log/postgresql" "/var/lib/postgresql" $WWW_HOME

COPY docker-entrypoint.sh /entrypoint.sh
ENTRYPOINT "/entrypoint.sh"
