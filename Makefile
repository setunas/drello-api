mysql-dsn = "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_TCP_HOST}:${DB_PORT})/${DB_NAME}"
migrate-cmd = migrate -path db/migrations -database $(mysql-dsn)

db-migrate:
	$(migrate-cmd) up 1

db-migrate-all:
	$(migrate-cmd) up

db-rollback:
	$(migrate-cmd) down 1

db-rollback-all:
	$(migrate-cmd) down 