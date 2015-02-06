

BUILD_DIR     := ./bin
CFLAGS        := 
MV            := mv
RM            := rm -rf
CP            := cp -rf
CD            := cd
MAKE          := make
MKDIR         := mkdir -p
GO            := go
FILES_SERVICE := $(notdir $(wildcard ./src/service/*))
FILES_TOOLS   := $(notdir $(wildcard ./src/tools/*))
FILES_GOOO    := $(notdir $(wildcard ./src/gooo/*))
PRE_FILES     := $(FILES:%.go=%)
ECHO          := echo

ECHO_DATE     := echo `date +%R:%S`

FUNC_FORRACH  := for i in $1;do $(GO) $2 $3/$$i ;done


.PHONY:  all build deploy service_build tools_build test gooo_test service_test tools_test run clean_pkg clean build_num_print build_num_record

default: build run

all: clean_pkg build test run

build: build_num_print service_build tools_build deploy build_num_record 

deploy:
	@$(MKDIR) $(BUILD_DIR) && $(CD) $(BUILD_DIR) && $(MKDIR) conf static log i18n
	@$(CP) ./conf ./i18n $(BUILD_DIR)
	@$(ECHO_DATE) Deploying...

service_build:
	@$(ECHO_DATE) Building service...
	@for i in $(FILES_SERVICE);do $(GO) install service/$$i; $(ECHO) service/$$i;done

tools_build:
	@$(ECHO_DATE) Building tools...
	@for i in $(FILES_TOOLS);do $(GO) install tools/$$i; $(ECHO) tools/$$i;done

test: gooo_test service_test tools_test
	
gooo_test:
	@$(ECHO_DATE) Testing gooo...
	@for i in $(FILES_GOOO);do $(GO) test gooo/$$i ;done

service_test:
	@$(ECHO_DATE) Testing service...
	@for i in $(FILES_SERVICE);do $(GO) test service/$$i ;done

tools_test:
	@$(ECHO_DATE) Testing tools...
	@for i in $(FILES_TOOLS);do $(GO) test tools/$$i ;done


run: 
	@$(ECHO_DATE) Running...
	@$(CD) $(BUILD_DIR) && ./master

clean_pkg:
	@$(ECHO_DATE) Pkg cleaning...
	@-$(RM) pkg

clean:
	@$(ECHO_DATE) Cleaning...
	@-$(RM) $(BUILD_DIR) pkg

build_num_print: build_num.txt
	@$(ECHO_DATE) The build num $(shell expr `cat $<` + 1)...

build_num_record: build_num.txt
	@$(ECHO) $(shell expr `cat $<` + 1) > $<
	@$(ECHO_DATE) Build num record...

