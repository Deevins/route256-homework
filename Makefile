ifeq ($(POSTGRES_SETUP_STRING),)
	POSTGRES_SETUP_STRING := user=test password=test dbname=postgres host=localhost port=5432 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MIGRATION_FOLDER=$(CURDIR)/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_STRING)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_STRING)" down

.PHONY: compose-up
compose-up:
	docker-compose build

.PHONY: compose-rm
compose-rm:
	docker-compose down

.PHONY: unit-test
unit-tests:
	go test .\internal\app\handlers -v

.PHONY: integration-test
unit-tests:
	go test .\tests\ -v

.PHONY: proto-server
proto-server:
#	rm -rf ./internal/pb
#	mkdir -p ./internal/pb
	protoc ./api/proto/server/*.proto \
               --go_out=./internal/pb/server \
               --go-grpc_out=./internal/pb/server \
               --proto_path=.

.PHONY: proto-client
proto-client:
#	rm -rf ./internal/pb
#	mkdir -p ./internal/pb

	protoc ./api/proto/client/*.proto \
               --go_out=./internal/pb/server \
               --go-grpc_out=./internal/pb/server \
               --proto_path=.