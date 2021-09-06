db-migrate:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:4306)/drello-dev" up 1

db-migrate-all:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:4306)/drello-dev" up

db-rollback:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:4306)/drello-dev" down 1

db-rollback-all:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:4306)/drello-dev" down