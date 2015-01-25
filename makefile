

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
PRE_FILES     := $(FILES:%.go=%)
ECHO          := echo


.PHONY:  all build deploy service_build tools_build run clean_pkg clean build_num_print build_num_record

all: clean_pkg build run

build: build_num_print service_build deploy
	@$(MAKE) build_num_record tools_build > /dev/null & 

deploy:
	@$(MKDIR) $(BUILD_DIR) && $(CD) $(BUILD_DIR) && $(MKDIR) conf static log i18n
	@$(CP) ./conf ./i18n $(BUILD_DIR)

service_build:
	@$(ECHO) $(foreach tar, $(FILES_SERVICE:%.go=%),`$(GO) install service/$(tar)` ) > /dev/null

tools_build:
	@$(ECHO) $(foreach tar, $(FILES_TOOLS:%.go=%),`$(GO) install tools/$(tar)` ) > /dev/null

run: 
	@$(ECHO) running...
	@$(CD) $(BUILD_DIR) && ./master

clean_pkg:
	@$(ECHO) pkg cleaning...
	@-$(RM) pkg

clean:
	@$(ECHO) cleaning...
	@-$(RM) $(BUILD_DIR) pkg

build_num_print: build_num.txt
	@$(ECHO) building $(shell expr `cat $<` + 1)...
	

build_num_record: build_num.txt
	@$(ECHO) $(shell expr `cat $<` + 1) > $<

