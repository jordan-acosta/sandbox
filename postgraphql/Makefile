# Makefile

all:
	@echo ""
	@echo "postgres-start: Start local postgres"
	@echo "postgres-stop: Stop local postgres"
	@echo "postgres-migrate: Migrate postgres to latest schema."
	@echo "api-run: Run postgraphql API."
	@echo "run: Start up the whole stack."
	@echo ""

.PHONY: postgres-start postgres-stop postgres-migrate api-run run

postgres-start:
	pg_ctl -D ./pgsql/data/ start -l postgres-log

postgres-stop:
	pg_ctl -D ./pgsql/data/ stop

postgres_migrate:
	migrate -database postgres://localhost:5432/postgres up

api-run:
	node --experimental-modules index.mjs

run:
	#$(MAKE) -j2 postgres-start run-api
