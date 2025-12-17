### Applying migration command

```bash
 migrate -path db/migrations \
  -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" \
  up
```


### Creating migration command


```bash
migrate create -ext sql -dir db/migrations -seq fix_users_id
```


### Generate The interfaces for application 

```bash
sqlc migrate
```