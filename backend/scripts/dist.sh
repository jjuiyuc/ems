#!/bin/bash

service="weather-worker"

echo "** building services"
go build -o dist/$service cmd/daemon/$service/$service.go

echo "** stopping services"
sudo systemctl stop $service 2>/dev/null

echo "** updating unit files"
cp cmd/daemon/$service/$service.service dist
sudo cp dist/$service.service /etc/systemd/system/$service.service
sudo mkdir -p /opt/derems
sudo cp dist/$service /opt/derems/$service

echo "** copying config"
sudo mkdir -p /opt/derems/config
sudo cp config/derems.yaml /opt/derems/config/derems.yaml

echo "** daemon reload"
sudo systemctl daemon-reload

echo "** starting services"
sudo systemctl start $service
