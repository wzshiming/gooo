


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

FILES_TEST   := $(notdir $(wildcard ./src/test/*/*.go))
PRE_FILES    := $(FILES:%.go=%)
ECHO 		 := echo


.PHONY: install build run clean elses

all: clean install run

install:
	@$(ECHO) installing...
	@$(MKDIR) $(BUILD_DIR) $(BUILD_DIR)/configs $(BUILD_DIR)/static $(BUILD_DIR)/log
	@$(CP) ./conf/*.json $(BUILD_DIR)/configs
	@$(ECHO) $(foreach tar, $(FILES_SERV:%.go=%),`$(GO) install service/$(tar)` ) > /dev/null
	@$(ECHO) $(foreach tar, $(FILES_TEST:%.go=%),`$(GO) install test/$(tar)` ) > /dev/null

run: 
	@$(ECHO) running...
	@$(CD) $(BUILD_DIR) && ./master

clean:
	@$(ECHO) cleaning...
	@$(RM) $(BUILD_DIR) pkg

