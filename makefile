#!/usr/bin/env make -f

BUILD_DIR     := ./bin
CFLAGS        := 
MV            := mv
RM            := rm -rf
CP            := cp -rf
CD            := cd
MAKE          := make
MKDIR         := mkdir -p
GO            := go
GIT           := git
ECHO          := echo
ECHO_DATE     := echo `date +%R:%S`
FLAGS         := -ldflags "-H windowsgui"

.PHONY:  all build makefile test run clean_pkg clean number_build number_build_record

default: build run

all: clean info test build clean_pkg run

coffee:
	@$(ECHO_DATE) Coffee...
	@coffee -c -j ./static/web/public/c.js ./src/ygo_frame/*.coffee

build: number_build dev number_build_record 

dev: shutdown coffee
	@$(ECHO_DATE) Dev...
	@$(GO) install -v github.com/wzshiming/server/...
	@$(GO) install -v service/...
	@$(MV) ./bin/resoud ./client/bin/resoud

prod: shutdown coffee
	@$(ECHO_DATE) Prod...
	@$(GO) install $(FLAGS) service/resoud
	@$(MV) ./bin/resoud ./client/bin/resoud
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

shutdown: 
	@$(ECHO_DATE) Shutdown...
	@-$(CD) $(BUILD_DIR) && ./shutdown && ./shutdown -server server_dev.json

clean_pkg:
	@$(ECHO_DATE) Cleaning pkg ...
	@-$(RM) pkg

clean: shutdown
	@$(ECHO_DATE) Cleaning...
	@-$(RM) $(BUILD_DIR)/* pkg


number_build: number_build.txt
	@$(GO) version
	@$(ECHO_DATE) The $(shell expr 1 + `sed -n 1p $<`) build...

number_build_record: number_build.txt
	@$(ECHO_DATE) Build number record...
	@$(ECHO) $(shell expr 1 + `sed -n 1p $<`) > $<

get: 
	@$(GO) get -v service/...
	
get_%:
	@export GOPATH=`pwd` &&\
	var="$(patsubst get_%,%,$@)" &&\
	$(GO) get -u -v $${var//_/\/}

get_else:
	@$(MKDIR) temp
	@$(GIT) clone https://github.com/mrdoob/three.js temp/three.js


