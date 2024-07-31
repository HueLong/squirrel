FROM php:8.1.0-fpm-alpine

ENV TIMEZONE Asia/Shanghai

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
  && mkdir -p /project/admin \
  && chmod 777 -R /project/admin
WORKDIR /project/admin

RUN apk update && apk upgrade
RUN  curl -L -o /tmp/composer.phar https://mirrors.aliyun.com/composer/composer.phar  \
  && mv /tmp/composer.phar /usr/local/bin/composer  \
  && chmod +x  /usr/local/bin/composer  \
  && composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/ \
  && apk add --update --no-cache libzip-dev tzdata oniguruma-dev autoconf g++  pcre-dev libtool make openssl-dev \
  freetype-dev libjpeg-turbo-dev libmcrypt-dev libpng-dev zlib-dev 
RUN docker-php-source extract  
RUN docker-php-ext-install -j"$(getconf _NPROCESSORS_ONLN)" iconv 
RUN docker-php-ext-configure bcmath --enable-bcmath 
RUN docker-php-ext-install pdo  \
  && docker-php-ext-install bcmath  \
  && docker-php-ext-install pdo_mysql  \
  && docker-php-ext-install zip \
  && ln -snf /usr/share/zoneinfo/$TIMEZONE /etc/localtime \
  && echo $TIMEZONE > /etc/timezone \
  && pecl install redis-5.3.4 \
  && docker-php-ext-enable redis 
RUN pecl channel-update pecl.php.net
RUN pecl install xdebug-3.1.2 \
  && docker-php-ext-enable xdebug \
  && echo -e "xdebug.mode=debug\nxdebug.client_host=host.docker.internal\nxdebug.idekey=PHPSTORM" \
  >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini
COPY mongodb-1.12.0.tgz /mongodb-1.12.0.tgz
RUN pecl install /mongodb-1.12.0.tgz \
  && docker-php-ext-enable mongodb \ 
  && docker-php-source delete \
  && rm -rf /tmp/pecl-memcache-4.0.4 php7.zip \
  && apk del autoconf g++ libtool make pcre-dev
