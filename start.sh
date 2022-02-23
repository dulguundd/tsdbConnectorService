#!/bin/bash
SERVER_ADDRESS=172.22.2.215 \
SERVER_PORT=9000 \
DB_USER=postgres \
DB_PASSWD=password \
DB_ADDR=172.22.2.215 \
DB_PORT=5432 \
DB_NAME=IoT_db_test \
go run main.go
