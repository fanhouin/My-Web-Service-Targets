#!/bin/bash
# http://localhost:4000/admin
# default user: your@email.com , password: spree123

wget https://github.com/spree/spree_starter/archive/main.zip &&
unzip main.zip &&
rm main.zip &&
cd spree_starter-main &&
bin/setup