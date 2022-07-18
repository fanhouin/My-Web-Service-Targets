# My-Web-Service-Targets

## Services
```bash
spree:4000      @ 4.4.0     - 1
magento2:80/443 @ 2.4.4     - 2
petstore:8080   @ 1.0.11    - 3
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