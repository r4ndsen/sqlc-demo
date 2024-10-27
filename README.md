# sqlc-demo

This sampole application will demonstrate how to use [sqlc] and [golang-migrate]

it collects url's and persists them on a local postgres database
and provides a simple web interface to view the collected urls

### Getting Started

* Start the database
    ```sh
    docker-compose up -d
    ```
  
* Generate the go types
    ```sh
    make t
    ```
  
* Run the migrations
    ```sh
    make migrate-up
    ```
    
* Build the application and start the web-server
    ```sh
    make server && \
    open http://localhost:3000
    ```

### Add a new migration

```sh
make migrate-new name=add_users_table
```

you will find two new files:
* `db/migrations/0002_add_users_table.down.sql` and
* `db/migrations/0002_add_users_table.up.sql`

edit them and re-run the type generation and migrations

```sh
make types && \
make migrate-up
```
