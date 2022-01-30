mysql-dsn = "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_TCP_HOST}:${DB_PORT})/${DB_NAME}"
test-mysql-dsn = "mysql://${TEST_DB_USER}:${TEST_DB_PASS}@tcp(${TEST_DB_TCP_HOST}:${TEST_DB_PORT})/${TEST_DB_NAME}"
migrate-cmd = ${GOPATH}/bin/migrate -path db/migrations -database

db-migrate:
	$(migrate-cmd) $(mysql-dsn) up 1

db-migrate-all:
	$(migrate-cmd) $(mysql-dsn) up

db-rollback:
	$(migrate-cmd) $(mysql-dsn) down 1

db-rollback-all:
	$(migrate-cmd) $(mysql-dsn) down 

db-migrate-all-test:
	$(migrate-cmd) $(test-mysql-dsn) up

db-rollback-all-test:
	$(migrate-cmd) $(test-mysql-dsn) down