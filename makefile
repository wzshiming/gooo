


BUILD_DIR    := ./bin
CFLAGS       := 
MV           := mv
RM           := -rm -rf
CP           := -cp
CD           := cd
MAKE         := make
MKDIR        := -mkdir -p
GO 			 := go
FILES 		 := $(notdir $(wildcard ./src/*/*.go))
PRE_FILES    := $(FILES:%.go=%)
ECHO 		 := echo


.PHONY: install build run clean elses

all: clean install run

install: elses
	@$(ECHO) installing...
	@$(MKDIR) $(BUILD_DIR) $(BUILD_DIR)/configs $(BUILD_DIR)/static $(BUILD_DIR)/log
	@$(CP) ./conf/*.json $(BUILD_DIR)/configs
	@$(ECHO) $(foreach tar, $(PRE_FILES),`$(GO) install $(tar)` ) > /dev/null


run: 
	@$(ECHO) running...
	@$(CD) $(BUILD_DIR) && ./master

clean:
	@$(ECHO) cleaning...
	@$(RM) $(BUILD_DIR) pkg

