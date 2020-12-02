include .env.test.local # considered only for run-test-local-only
export

up:
	docker-compose up -d db
	curl https://s3.amazonaws.com/careers-picpay/users.csv.gz --output data/users.csv.gz
	gunzip -f -k data/users.csv.gz > data/users.csv
	docker exec parsimonious-users-api_db_1 bash -c 'while !</dev/tcp/db/5432; do sleep 1; done;'
	docker exec parsimonious-users-api_db_1 psql -U postgres -f /opt/data/migration.sql
	docker exec parsimonious-users-api_db_1 psql -U postgres -f /opt/data/import-data.sql
	docker-compose up -d app

down:
	docker-compose down

run-adminer:
	docker-compose up -d adminer

prepare-test:
	docker-compose up -d test-db
	docker exec parsimonious-users-api_test-db_1 bash -c 'while !</dev/tcp/test-db/5432; do sleep 1; done;'
	docker exec parsimonious-users-api_test-db_1 psql -U postgres -f /opt/data/migration.sql
	docker exec parsimonious-users-api_test-db_1 psql -U postgres -f /opt/data/import-data.sql

run-test-only:
	docker-compose up --build test

run-test: prepare-test run-test-only

run-test-local-only:
	go test -v -count=1 ./...

run-test-local: prepare-test run-test-local-only

	