export DRIVER=mysql
export CONNECTION_STRING=user:123456@tcp(localhost:3306)/account_control_db?parseTime=true
export MIGRATION_DIR="./infra/db/migrations"
export COVERAGE_PATH=./coverage.out

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migrate-up:
	goose --dir=$(MIGRATION_DIR) $(DRIVER) $(CONNECTION_STRING) up

# make add-migration -e MIGRATION_NAME=create_table
add-migration:
ifeq ($(strip $(MIGRATION_NAME)),)
	@echo "Invalid value for variable MIGRATION_NAME."
	@echo "Try the command: $ make add-migration -e MIGRATION_NAME=<migration_name>"
	exit 1
else
	goose -dir ./infra/db/migrations create $(MIGRATION_NAME) sql
endif

ping-local-db:
	goose mysql "$(CONNECTION_STRING)" status

test:
	go test ./... -cover -coverprofile=$(COVERAGE_PATH) && go tool cover -html=$(COVERAGE_PATH)
