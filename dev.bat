@echo off
cd api
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go install
cd ../