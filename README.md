# My-Web-Service-Targets

## Services
```bash
spree:4000      @ 4.4.0     - 1
magento2:80/443 @ 2.4.4     - 2
petstore:8001   @ 1.0.11    - 3
onos:8002   	@ 2.5.8		- 4
redmine:8003    @ 5.0.2		- 5
```
## Setup
```bash
# for spree
cd spree
./setup.sh

# for magento2 self build
cd magento2/self_build
docker build -t <your_name>/magento_test . # and follow the comments to complete the setting
```

## Usage
```bash
# run the service
service_up.sh <service_num>

# stop the service
service_down.sh <service_num>
```