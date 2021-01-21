@echo off
echo starting...
cd /D "%~dp0"
set GOOS=linux
go build .
echo done!