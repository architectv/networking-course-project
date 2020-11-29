#!/bin/sh

commonname="yak"
country="RU"
state="Moscow"
locality="Moscow"
organization="Yak Team Ltd"
organizationalunit="IT"
email="root@yak.ru"

openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout nginx.key -out nginx.crt \
    -subj "/C=$country/ST=$state/L=$locality/O=$organization/OU=$organizationalunit/CN=$commonname/emailAddress=$email"
