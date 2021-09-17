mysql-dsn = "mysql://root:password@tcp(127.0.0.1:4306)/drello-dev"
migrate-cmd = migrate -path db/migrations -database $(mysql-dsn)

db-migrate:
	$(migrate-cmd) up 1

db-migrate-all:
	$(migrate-cmd) up

db-rollback:
	$(migrate-cmd) down 1

db-rollback-all:
	$(migrate-cmd) down 