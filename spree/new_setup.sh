#!/bin/bash
# http://localhost:4000/admin
# default user: your@email.com , password: spree123

wget https://github.com/spree/spree_starter/archive/refs/tags/v4.5.0.zip &&
unzip v4.5.0.zip &&
rm v4.5.0.zip &&
cd spree_starter-4.5.0 && 
bin/setup
