COMPOSE_APP = docker compose -f compose.yml

.PHONY: app-build app-up app-down app-shell app-logs gen-proto \
        client-build client-up client-down client-shell client-logs client-rail

build:
	@echo ">>> Building APP image…"
	$(COMPOSE_APP) build

up:
	@echo ">>> Starting APP container…"
	$(COMPOSE_APP) up -d

down:
	@echo ">>> Stopping APP container…"
	$(COMPOSE_APP) down

shell:
	@echo ">>> Opening shell into APP container…"
	$(COMPOSE_APP) exec app bash

logs:
	@echo ">>> Tailing logs for APP…"
	$(COMPOSE_APP) logs -f

gen-proto:
	@echo ">>> Generating Go code from .proto files in container…"
	$(COMPOSE_APP) exec app bash -lc '\
	 protoc \
       -I proto \
       --go_out=gen/pb --go_opt=paths=source_relative \
       --go-grpc_out=gen/pb --go-grpc_opt=paths=source_relative \
       proto/*.proto \
	'