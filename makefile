# /usr/bin/env make

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



.PHONY:  all build deploy test run clean_pkg clean number_build number_build_record

default: build run

all: clean info test build clean_pkg deploy run

build: number_build build_service build_tools number_build_record deploy

deploy:
	@$(ECHO_DATE) Deploying...
	@$(MKDIR) $(BUILD_DIR) && $(CD) $(BUILD_DIR) && $(MKDIR) conf log i18n db
	@$(CP) ./conf ./i18n $(BUILD_DIR)

build_%:
	@$(ECHO_DATE) Building $(patsubst build_%,%,$@) ...
	@export GOPATH=`pwd` &&\
	 for i in $(notdir $(wildcard ./src/$(patsubst build_%,%,$@)/*));\
	 do\
	 	$(GO) install $(patsubst build_%,%,$@)/$$i;\
	 	$(ECHO) $(patsubst build_%,%,$@)/$$i;\
	 done

test: test_gooo test_service test_tools

test_gooo:
	@$(ECHO_DATE) Testing gooo...
	@export GOPATH=`pwd` && $(GO) test gooo

test_%:
	@$(ECHO_DATE) Testing $(patsubst test_%,%,$@)...
	@export GOPATH=`pwd` &&\
	 for i in $(notdir $(wildcard ./src/$(patsubst test_%,%,$@)/*));\
	 do\
	 	$(GO) test $(patsubst test_%,%,$@)/$$i;\
	 done

run: deploy
	@$(ECHO_DATE) Running...
	@$(CD) $(BUILD_DIR) && ./master

clean_pkg:
	@$(ECHO_DATE) Cleaning pkg ...
	@-$(RM) pkg

clean:
	@$(ECHO_DATE) Cleaning...
	@-$(RM) $(BUILD_DIR)/* pkg


number_build: number_build.txt
	@$(ECHO_DATE) The $(shell expr 1 + `sed -n 1p $<`) build...

number_build_record: number_build.txt
	@$(ECHO_DATE) Build number record...
	@$(ECHO) $(shell expr 1 + `sed -n 1p $<`) > $<

get: get_github.com_go-sql-driver_mysql\
	 get_github.com_mattn_go-sqlite3\
	 get_github.com_jinzhu_gorm\
	 get_github.com_kortem_lingo\
	 get_github.com_lib_pq\
	 get_github.com_go-martini_martini\
	 get_github.com_martini-contrib_render\
	 get_github.com_gorilla_websocket\
	 get_github.com_wzshiming_ffmt
	
get_%:
	@export GOPATH=`pwd` &&\
	var="$(patsubst get_%,%,$@)" &&\
	$(GO) get -u -v $${var//_/\/}

get_else:
	@$(MKDIR) temp
	@$(GIT) clone https://github.com/mrdoob/three.js temp/three.js

info:
	@$(ECHO_DATE) Gooo source `cat ./src/gooo/*.go | wc -l` line...
	@$(ECHO_DATE) Service source `cat ./src/service/*/*.go | wc -l` line...
	@$(ECHO_DATE) Tools source `cat ./src/tools/*/*.go | wc -l` line...
