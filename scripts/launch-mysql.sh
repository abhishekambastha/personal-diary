#!/bin/bash

docker run -d \
	-p 3306:3306 \
	--name diary \
	-e MYSQL_ROOT_PASSWORD=diary \
	-e MYSQL_ROOT_HOST=% \
	mysql/mysql-server:latest-aarch64
