#!/bin/bash
if [ $# -eq 0 ]
  then
    echo "Usage: service_up.sh <service_num>
Service_num: 
	spree:4000      @ 4.4.0 	- 1
	magento2:80/443 @ 2.4.4		- 2
	petstore:8080   @ 1.0.11	- 3"
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
esac

