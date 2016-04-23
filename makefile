#! /usr/bin/env make -f

BUILD_DIR     := ./bin
CFLAGS        := 
MV            := mv
RM            := rm -rf
CP            := cp -rf
CD            := cd
MAKE          := make
MKDIR         := mkdir -p
GO            := if [ -z $$GOPATH ]; then export GOPATH=`pwd` ; export GOBIN=`pwd`/bin ;fi && go
GIT           := git
ECHO          := echo
ECHO_DATE     := echo `date +%R:%S`
FLAGS         := -ldflags "-H windowsgui"
REM           := 1>/dev/null 2>&1

.PHONY:  all build makefile test run clean_pkg clean_bin clean number_build number_build_record

default: build run

all: clean info test build clean_pkg run

build: number_build dev number_build_record  


dev: clean_bin
	@$(ECHO_DATE) Dev...
	@$(GO) install -v github.com/wzshiming/server/...
	@$(GO) install -v service/...

prod: clean_bin
	@$(ECHO_DATE) Prod...
	@$(GO) install -v github.com/wzshiming/server/...
	@$(GO) install -v service/...

test:
	@$(GO) test github.com/wzshiming/...
	@$(GO) test service/...

test_v:
	@$(GO) test -v github.com/wzshiming/...
	@$(GO) test -v service/...

run: shutdown
	@$(ECHO_DATE) Dev Running...
	@$(CD) $(BUILD_DIR) && ./master -server server_dev.json

run_prod:
	@$(ECHO_DATE) Prod Running...
	@$(CD) $(BUILD_DIR) && ./master 

shut: shutdown

shutdown: 
	@$(ECHO_DATE) Shutdown...
	@-$(CD) $(BUILD_DIR) $(REM) && ./shutdown $(REM) && ./shutdown -server server_dev.json $(REM)

clean_pkg:
	@$(ECHO_DATE) Cleaning pkg ...
	@-$(RM) pkg $(REM)
	
clean_bin: shutdown
	@$(ECHO_DATE) Cleaning bin ...
	@-$(RM) $(BUILD_DIR) $(REM)

clean: 
	@$(ECHO_DATE) Cleaning...
	@-$(MAKE) clean_pkg clean_bin $(REM)

number_build: number_build.txt
	@$(GO) version
	@-$(ECHO_DATE) The $(shell expr 1 + `sed -n 1p $<`) build...

number_build_record: number_build.txt
	@$(ECHO_DATE) Build number record...
	@-$(ECHO) $(shell expr 1 + `sed -n 1p $<`) > $<

get: 
	@$(GO) get -v service/...
