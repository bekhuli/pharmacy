MIGRATE=go run cmd/migrate/main.go

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Usage: make migrate-force version=<version_number>"; \
		exit 1; \
	fi
	$(MIGRATE) force $(version)

migration:
	@if [ -z "$(word 2, $(MAKECMDGOALS))" ]; then \
		echo "Usage: make new-migration <name>"; \
	else \
		migrate create -ext sql -dir migrations -seq $(word 2, $(MAKECMDGOALS)); \
	fi

%:
	@: