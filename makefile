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



.PHONY:  all build makefile test run clean_pkg clean number_build number_build_record

default: build run

all: clean info test build clean_pkg run

build: number_build build_ number_build_record 

build_: shutdown
	@$(GO) install -v rego/...
	@$(GO) install -v service/...

test:
	@$(GO) test rego/...
	@$(GO) test service/...

test_v:
	@$(GO) test -v rego/...
	@$(GO) test -v service/...

run: shutdown
	@$(ECHO_DATE) Running...
	@$(CD) $(BUILD_DIR) && ./master 

shutdown: 
	@$(ECHO_DATE) Shutdown...
	@-$(CD) $(BUILD_DIR) && ./shutdown

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

get: get_github.com_go-sql-driver_mysql\
	 get_github.com_mattn_go-sqlite3\
	 get_github.com_jinzhu_gorm\
	 get_github.com_kortem_lingo\
	 get_github.com_go-martini_martini\
	 get_github.com_martini-contrib_render\
	 get_github.com_martini-contrib_sessions\
	 get_github.com_gorilla_websocket\
	 get_github.com_wzshiming_ffmt\
	 get_github.com_mitchellh_colorstring
	
get_%:
	@export GOPATH=`pwd` &&\
	var="$(patsubst get_%,%,$@)" &&\
	$(GO) get -u -v $${var//_/\/}

get_else:
	@$(MKDIR) temp
	@$(GIT) clone https://github.com/mrdoob/three.js temp/three.js

info:
	@wc ./src/rego/*.go ./src/rego/*/*.go ./src/rego/*/*/*.go
