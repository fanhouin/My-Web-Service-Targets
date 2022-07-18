#!/bin/bash
wget https://github.com/spree/spree_starter/archive/main.zip &&
unzip main.zip &&
rm main.zip &&
cd spree_starter-main &&
bin/setup