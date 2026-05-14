db_login:
	psql ${DB_URL}

db_create_migration:
	migrate create -ext sql -dir migrations -seq ${name}

db_run_migrations:
	migrate -database ${DB_URL} -path migrations up

