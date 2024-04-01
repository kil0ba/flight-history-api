## Start a project

```bash
$ make
```

## Start a local DB (docker required)

```bash
$ make compose-up
```

## Create new migration

```bash
$ migrate create -ext sql -dir migrations -seq create_users_table
```

## Run migrations

```bash
$ make migrate-db
```