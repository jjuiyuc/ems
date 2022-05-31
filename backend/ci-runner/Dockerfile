FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y make git wget sed curl python apt-utils gcc pkg-config software-properties-common

# kafka
RUN wget -O kafka.tgz https://archive.apache.org/dist/kafka/3.0.0/kafka_2.13-3.0.0.tgz
RUN mkdir /opt/kafka
RUN tar -xzvf kafka.tgz --strip 1 -C /opt/kafka

# mysql 8.0
RUN wget https://repo.mysql.com/mysql-apt-config_0.8.13-1_all.deb
RUN dpkg -i mysql-apt-config_0.8.13-1_all.deb
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 467B942D3A79BD29
RUN apt-get update && apt-get install -y mysql-server

# go 1.18
RUN wget https://dl.google.com/go/go1.18.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz

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