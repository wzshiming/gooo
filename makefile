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
ECHO          := echo
ECHO_DATE     := echo `date +%R:%S`



.PHONY:  all build deploy test run clean_pkg clean number_build number_build_record

default: build deploy

all: clean info test build clean_pkg deploy run

build: number_build build_service build_tools number_build_record 

deploy:
	@$(ECHO_DATE) Deploying...
	@$(MKDIR) $(BUILD_DIR) && $(CD) $(BUILD_DIR) && $(MKDIR) conf static log i18n
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

run: 
	@$(ECHO_DATE) Running...
	@$(CD) $(BUILD_DIR) && ./master

clean_pkg:
	@$(ECHO_DATE) Cleaning pkg ...
	@-$(RM) pkg

clean:
	@$(ECHO_DATE) Cleaning...
	@-$(RM) $(BUILD_DIR) pkg


number_build: number_build.txt
	@$(ECHO_DATE) The $(shell expr 1 + `sed -n 1p $<`) build...

number_build_record: number_build.txt
	@$(ECHO_DATE) Build number record...
	@$(ECHO) $(shell expr 1 + `sed -n 1p $<`) > $<

get:
	@export GOPATH=`pwd` &&\
	$(ECHO_DATE) Getting github.com/go-sql-driver/mysql ...;\
	$(GO) get github.com/go-sql-driver/mysql &&\
	$(ECHO_DATE) Getting github.com/jinzhu/gorm ...;\
	$(GO) get github.com/jinzhu/gorm &&\
	$(ECHO_DATE) Getting github.com/kortem/lingo ...;\
	$(GO) get github.com/kortem/lingo &&\
	$(ECHO_DATE) Getting github.com/lib/pq ...;\
	$(GO) get github.com/lib/pq

info:
	@$(ECHO_DATE) Gooo source `cat ./src/gooo/*.go | wc -l` line...
	@$(ECHO_DATE) Service source `cat ./src/service/*/*.go | wc -l` line...
	@$(ECHO_DATE) Tools source `cat ./src/tools/*/*.go | wc -l` line...
