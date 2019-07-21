@echo off
cd api
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o ../bin/api
cd ../
echo Package success~~~