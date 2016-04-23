@echo off
set path=;
set path=%path%;d:\lib\nodejs
set path=%path%;d:\lib\ruby2\bin
set path=%path%;d:\lib\python27
set path=%path%;d:\lib\python33
set path=%path%;d:\lib\php-5.3
set path=%path%;d:\lib\mongodb\bin
set path=%path%;d:\lib\mingw64\bin
set path=%path%;d:\lib\mingw64\msys\1.0\bin
set path=%path%;d:\lib\cygwin64\bin
set path=%path%;d:\lib\Apache2\bin
set path=%path%;d:\lib\Git\bin
set path=%path%;d:\lib\go64\bin
set path=%path%;C:\Users\zsm\AppData\Roaming\npm
set GOROOT=d:\lib\go64
set GOPATH=%CD%
set PWD=GOPATH
rem set GOOS=linux
rem set GOARCH=amd64
set GOOS=windows
rem set GOARCH=386
rem set GOOS=linux
rem start powershell -command zsh 
start mintty -e zsh