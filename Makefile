.PHONY: all $(CMD_TARGETS)

OUTPUT_DIR=bin

CMD_TARGETS = $(notdir $(shell find cmd/* -maxdepth 0 -type d))

all: $(CMD_TARGETS)

#$(CMD_TARGETS): proto
$(CMD_TARGETS):
	export GOOS="linux" && CGO_ENABLED=0 go build -o $(OUTPUT_DIR)/$@ ./cmd/$@

#proto: protocol/*.proto
#	@bash scripts/codegen.sh protocol ./protocol/*.proto protocol/pb/go
#	#@bash scripts/codegen.sh
