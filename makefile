


BUILD_DIR    := ./bin
CFLAGS       := 
MV           := mv
RM           := -rm -rf
CP           := -cp
CD           := cd
MAKE         := make
MKDIR        := -mkdir -p
GO 			 := go
FILES_SERV 	 := $(notdir $(wildcard ./src/service/*/*.go))
FILES_TEST   := $(notdir $(wildcard ./src/tools/*/*.go))
PRE_FILES    := $(FILES:%.go=%)
ECHO 		 := echo


.PHONY: install build run clean elses

all: clean_pkg install run

install:
	@$(ECHO) installing...
	@$(MKDIR) $(BUILD_DIR) $(BUILD_DIR)/conf $(BUILD_DIR)/static $(BUILD_DIR)/log $(BUILD_DIR)/i18n
	@$(CP) ./conf/*.json $(BUILD_DIR)/conf
	@$(CP) ./i18n/*.json $(BUILD_DIR)/i18n
	@$(ECHO) $(foreach tar, $(FILES_SERV:%.go=%),`$(GO) install service/$(tar)` ) > /dev/null
	@$(ECHO) $(foreach tar, $(FILES_TEST:%.go=%),`$(GO) install tools/$(tar)` ) & > /dev/null

run: 
	@$(ECHO) running...
	@$(CD) $(BUILD_DIR) && ./master

clean_pkg:
	@$(ECHO) pkg cleaning...
	@$(RM) pkg

clean:clean_pkg
	@$(ECHO) cleaning...
	@$(RM) $(BUILD_DIR)
