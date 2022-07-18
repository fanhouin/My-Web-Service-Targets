#!/bin/bash
if [ $# -eq 0 ]
  then
    echo "Usage: service_up.sh <service_num>
Service_num: 
	spree:4000      @ 4.4.0 	- 1
	magento2:80/443 @ 2.4.4		- 2
	petstore:8001   @ 1.0.11	- 3
	onos:8002   	@ 2.5.8		- 4
	redmine:8003    @ 5.0.2		- 5"
	exit 
fi

case $1 in
1)
	cd spree/spree_starter-main &&\
	docker-compose down
	;;
2)
	cd magento2 && \
	docker-compose down
	;;
3)
	cd petstore && \
	docker-compose down
	;;	
4) 
	cd onos	&& \
	docker-compose down
	;;
5) 
	cd redmine	&& \
	docker-compose down
	;;
esac

