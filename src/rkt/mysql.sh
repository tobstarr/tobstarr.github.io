#!/bin/bash

## this should be some ExecStartPre
sudo mkdir -p /var/run/mysqld /opt/rkt/mysql
sudo chown 999 /var/run/mysqld

## this should be ExecStart
sudo rkt run --set-env=MYSQL_ALLOW_EMPTY_PASSWORD=true --volume volume-var-lib-mysql,kind=host,source=/opt/rkt/mysql --volume socket,kind=host,source=/var/run/mysqld --mount=volume=socket,target=/run/mysqld --insecure-skip-verify docker://mysql:5.6.27
