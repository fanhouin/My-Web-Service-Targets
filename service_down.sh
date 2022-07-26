#!/bin/bash
if [ $# -eq 0 ]
  then
    echo "Usage: service_down.sh <service_num>
Service_num: 
	petstore:8001   @ 1.0.11	- 1
	spree:4000      @ 4.4.0 	- 2
	magento2:80/443 @ 2.4.4		- 3
	gitlab:80/443 	@ 15.1.2	- 4
	onos:8002   	@ 2.5.8		- 5
	redmine:8003    @ 5.0.2		- 6"
	exit 
fi

case $1 in
1)
	cd petstore && \
	docker-compose down
	;;	
2)
	cd spree/spree_starter-main &&\
	docker-compose down
	;;
3)
	cd magento2 && \
	docker-compose down
	;;
4) 
	cd gitlab	&& \
	docker-compose down
	;;
5) 
	cd onos	&& \
	docker-compose down
	;;
6) 
	cd redmine	&& \
	docker-compose down
	;;

esac

