# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

.PHONY: help create-migration migrate-up migrate-down migrate-force

##@ Migration Commands

help: ## Show help
	@echo "\n\033[1mAvailable commands:\033[0m\n"
	@awk 'BEGIN {FS = ":.*##";} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

create-migration: ## Create an empty migration file
	@read -p "Enter the migration name (snake_case): " SEQ; \
	docker run --rm -v $(CURDIR)/database/migrations:/migrations migrate/migrate \
		create -ext sql -dir /migrations -seq $${SEQ}

migrate-up: ## Apply all up migrations
	docker run --rm -v $(CURDIR)/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PWD)@$(POSTGRES_URL):$(POSTGRES_PORT)" up

migrate-down: ## Roll back the last N migrations (default: 1)
	@read -p "Number of migrations to roll back (default: 1): " NUM; NUM=$${NUM:-1}; \
	docker run --rm -v $(CURDIR)/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PWD)@$(POSTGRES_URL):$(POSTGRES_PORT)" down $${NUM}

migrate-force: ## Force migration to specific version
	@read -p "Enter the version to force: " VERSION; \
	docker run --rm -v $(CURDIR)/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PWD)@$(POSTGRES_URL):$(POSTGRES_PORT)" force $${VERSION}
