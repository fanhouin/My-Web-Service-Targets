#!/bin/bash
if [ $# -eq 0 ]
  then
    echo "Usage: service_up.sh <service_num>
Service_num: 
	petstore:8001   @ 1.0.11	- 1
	spree:4000      @ 4.5.0 	- 2
	magento2:80/443 @ 2.4.4		- 3
	gitlab:80/443 	@ 15.1.2	- 4
	onos:8002   	@ 2.5.8		- 5
	redmine:8003    @ 5.0.2		- 6
	realworld:8004  @ 728cb04	- 7"
	exit 
fi

case $1 in
1)
	cd petstore && \
	docker-compose up -d
	;;	
2)
	cd spree/spree_starter-4.5.0 &&\
	docker-compose up -d
	;;
3)
	cd magento2 && \
	docker-compose up -d
	;;
4)
	cd gitlab	&& \
	docker-compose up -d
	;;
5)
	cd onos	&& \
	docker-compose up -d
	;;
6)
	cd redmine	&& \
	docker-compose up -d
	;;
7)
	cd realworld && \
	docker-compose up -d
	;;

esac

