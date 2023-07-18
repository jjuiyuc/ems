FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y make git wget sed curl python apt-utils gcc pkg-config software-properties-common

# kafka
RUN wget -O kafka.tgz https://archive.apache.org/dist/kafka/3.0.0/kafka_2.13-3.0.0.tgz
RUN mkdir /opt/kafka
RUN tar -xzvf kafka.tgz --strip 1 -C /opt/kafka

# install ssh client
RUN apt-get update -y && apt-get install openssh-client -y

# install python
RUN apt-get update -y && apt-get -y install python3 python3-pip python3-dev

# install ansible
RUN pip3 install ansible

# mysql 8.0
RUN wget https://repo.mysql.com/mysql-apt-config_0.8.13-1_all.deb
RUN dpkg -i mysql-apt-config_0.8.13-1_all.deb
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 467B942D3A79BD29

RUN apt-get update && apt-get install -y mysql-server openjdk-8-jre-headless

# node.js
RUN curl -sL https://deb.nodesource.com/setup_14.x -o nodesource_setup.sh
RUN bash nodesource_setup.sh
RUN apt install nodejs

# yarn
RUN npm install yarn -g

# jest-junit
RUN yarn add jest-junit -g

# go 1.20
RUN wget https://dl.google.com/go/go1.20.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz

RUN sed -i 's?"$?:/usr/local/go/bin"?' /etc/environment
ENV PATH "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/usr/local/go/bin:/usr/local/google-cloud-sdk/bin"

# sqlboiler
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
ENV PATH "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/usr/local/go/bin:/usr/local/google-cloud-sdk/bin:/root/go/bin"

RUN go install github.com/t-yuki/gocover-cobertura@latest

# start
ADD start.sh /
ADD server.properties /opt/kafka/config/server.properties
