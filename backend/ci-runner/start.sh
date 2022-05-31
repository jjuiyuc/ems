#!/bin/bash

mkdir /var/run/mysqld
# --skip-name-resolve solves the problem of client connect() taking too long
mysqld --user=root --skip-name-resolve &

while [ ! -S /var/run/mysqld/mysqld.sock ]; do
  echo "waiting for mysql to finish starting"
  sleep 1
done
echo "mysql online"

mysqladmin create der_ems

mysql -u root -e "CREATE USER 'ubiik'@'%' IDENTIFIED BY 'testing123';"
mysql -u root -e "GRANT ALL PRIVILEGES ON *.* TO 'ubiik'@'%' WITH GRANT OPTION;"

/opt/kafka/bin/zookeeper-server-start.sh /opt/kafka/config/zookeeper.properties >/tmp/zookeeper.log 2>&1 &

ok=1
while [ $ok -ne 0 ]; do
  echo "waiting for zookeeper to finish starting"
  /opt/kafka/bin/zookeeper-shell.sh localhost:2181 stat / >/dev/null
  ok=$?
done
echo "zookeeper online"

if [ "$1" = "-block" ]; then
  /opt/kafka/bin/kafka-server-start.sh /opt/kafka/config/server.properties
else
  /opt/kafka/bin/kafka-server-start.sh /opt/kafka/config/server.properties >/tmp/kafka.log 2>&1 &
  # wait for kafka to finish starting
  ok=1
  while [ $ok -ne 0 ]; do
    echo "waiting for kafka to finish starting"
    /opt/kafka/bin/zookeeper-shell.sh localhost:2181 ls /brokers/ids/0 >/dev/null
    ok=$?
  done
fi
echo "kafka online"
