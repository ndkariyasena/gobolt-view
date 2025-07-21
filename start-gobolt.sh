#! /bin/bash

cd cmd/server || exit 1
go run main.go start && return 
cd web || exit 1
npm run start && return