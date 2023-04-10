#!/bin/bash
goc run . &
sleep 3 
cd coverage
go run coverage.go


# lsof -i -P -n | grep goc | grep LISTEN > goc.ip
# while true; do
#   sleep 60
# done