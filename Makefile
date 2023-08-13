GOHOSTOS:=$(shell go env GOHOSTOS)

ifeq ($(GOHOSTOS), windows)
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	ADMIN_RPC_PROTO_FILES=$(shell $(Git_Bash) -c "find services/admin/proto -name *.proto")
else
	ADMIN_RPC_PROTO_FILES=$(shell find services/admin/proto -name *.proto)
endif

define generate
	protoc --proto_path=./services/$(1)/proto --proto_path=./services/$(1)/proto/message  --go_out=./services/$(1) --go-grpc_out=./services/$(1)  $(2)
endef

.PHONY: admin
admin:
	$(call generate,$@, $(ADMIN_RPC_PROTO_FILES))
.PHONY: admin_sql
admin_sql:
	goctl model mysql ddl --src ./data/sql/admin.sql --dir ./services/admin/internal/model

.PHONY: wire
wire:
	cd services/admin && wire

.PHONY: env
env:
	cd deploy/docker && docker-compose up -d
.PHONY: admin_build
admin_build:
	cd services/admin && go build -o server admin.go wire_gen.go



