1. cd to this directory
1. `docker compose up` for postgres
1. `docker exec -it postgres bash`
1. `psql -U admin golang_postgres`
1. make sure that uuid-ossp extension is installed by `select * from pg_available_extensions;`
1. if not create it, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
1. run the migration using migrate tool [here](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate), run after that, `migrate -path db/migrations -database "postgresql://admin:password123@localhost:6500/golang_postgres?sslmode=disable" -verbose up`
1. run `sqlc generate`
1. use air to run the server `alias air='$(go env GOPATH)/bin/air'` then run `air`