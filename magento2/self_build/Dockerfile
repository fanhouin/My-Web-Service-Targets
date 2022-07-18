FROM ubuntu:20.04

RUN apt-get update && \
    apt-get install -y \
        apt-transport-https \
        software-properties-common \
        zip \
        unzip \
        curl \
        wget \
        vim 

# install java && elasticsearch
RUN apt-get install -y openjdk-8-jdk && \
    apt -y install gnupg2 && \
    wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | apt-key add - && \
    echo "deb https://artifacts.elastic.co/packages/oss-7.x/apt stable main" | tee  /etc/apt/sources.list.d/elastic-7.x.list && \
    apt update && \
    apt install -y elasticsearch-oss=7.9.1
    # vim /etc/elasticsearch/elasticsearch.yml
    # and change the following lines:
        # network.host: 192.168.0.1 -> network.host: localhost
    # service elasticsearch restart

# install apache2
RUN apt-get install -y apache2 && \
    a2enmod rewrite
    # vim etc/apache2/sites-available/000-default.conf
    # add this line:
        # <Directory "/var/www/html">
        #     AllowOverride All
        # </Directory>
    # service apache2 restart
 

# install php-fpm & php-cli
RUN add-apt-repository ppa:ondrej/php && \
    apt-get -y install php7.2-fpm php7.2-cli
    # vim /etc/php/7.2/fpm/php.ini
    # vim /etc/php/7.2/cli/php.ini
    # edit the line:
        # memory_limit = 2G
        # max_execution_time = 1800
        # zlib.output_compression = On
    # service php7.2-fpm restart
    
# install php7.4
RUN apt-get -y install php7.4 \
    php7.4-pdo  \
    php7.4-mysqlnd  \
    php7.4-opcache \
    php7.4-xml  \
    php7.4-gd \
    php7.4-mysql \
    php7.4-intl \
    php7.4-mbstring \
    php7.4-bcmath \
    php7.4-json \
    php7.4-iconv \
    php7.4-soap \
    php7.4-zip \
    php7.4-curl

RUN apt install -y mysql-server && \ 
    ss -tap | grep mysql && \
    service mysql restart
    # mysql -u root -p 
    # create database magento; 
    # create user 'magento'@'localhost' IDENTIFIED BY 'magento'; 
    # GRANT ALL ON magento.* TO 'magento'@'localhost';
    # flush privileges;
    # exit 

RUN mkdir /var/www/html/magento2

# vim /etc/apache2/sites-available/magento2.conf
# <VirtualHost *:80>
#  	ServerAdmin admin@domain.com
#  	ServerName magento2.local
#  	ServerAlias magento2.local
#  	DocumentRoot /var/www/html/magento2/pub
#  	<Directory /var/www/html/magento2/>
#     	Options Indexes FollowSymLinks MultiViews
#     	AllowOverride All
#     	Order allow,deny
#     	allow from all
#  	</Directory>
#  	ErrorLog ${APACHE_LOG_DIR}/magento_error.log
#  	CustomLog ${APACHE_LOG_DIR}/magento_access.log combined
# </VirtualHost>


# vim /etc/hosts
# 127.0.0.1   magento2.local

# php bin/magento setup:install \
# --base-url=http://magento2.local/ \
# --db-host=localhost \
# --db-name=magento \
# --db-user=magento \
# --db-password=magento \
# --backend-frontname=admin \
# --admin-firstname=admin \
# --admin-lastname=admin \
# --admin-email=admin@admin.com \
# --admin-user=admin \
# --admin-password=admin123 \
# --language=en_US \
# --currency=USD \
# --timezone=America/Chicago \
# --use-rewrites=1 \
# --search-engine=elasticsearch7 \
# --elasticsearch-host=localhost \
# --elasticsearch-port=9200

# chmod -R 755 /var/www/html/magento2/

# php bin/magento setup:upgrade
# php bin/magento setup:di:compile
# php bin/magento setup:static-content:deploy -f
# php bin/magento indexer:reindex
# php bin/magento cache:clean
# php bin/magento cache:flush
# chmod -R 777 pub/ var/ generated/

# php bin/magento deploy:mode:set developer
# php bin/magento cache:flush

# vim app/etc/env.php
# 'host' => 'localhost' -> 'host' => '127.0.0.1'