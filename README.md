# Go Juice Inventory

This runs a simple HTTP server that allows the user to work with inventory at a juice shop.

## Getting Started

### Building the project
Use the following command to build the project
```bash
docker-compose build --pull juice-inventory
```

### Running the project
Use the following command to run the project
```bash
docker-compose up
```

### Making requests
By default, the service runs at localhost:8090

#### Show all inventory
```bash
curl localhost:8090/products
```

#### Delete an item
To delete an item, you must specify its ID as a path parameter
```bash
curl -X DELETE localhost:8090/products/{id}
```

### Debugging the project
In order to effectively debug, you can run the project without docker. If the project has been run before, the postgres
database can be run as a standalone container like so (assumes that go-juice-inventory_postgres_data is available
via `docker volume ls`):
```bash
docker run -d \
    --name inventory-postgres \
    -e POSTGRES_PASSWORD=changeme \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v go-juice-inventory_postgres_data:/var/lib/postgresql/data \
    -p 5432:5432/tcp \
    --expose=5432 \
    postgres:13.3-alpine
```

Then, you can `go run` the program:
```bash
POSTGRES_USER=postgres \
    POSTGRES_PASSWORD=changeme \
    POSTGRES_DB=postgres \
    POSTGRES_HOST=localhost\
    POSTGRES_PORT=5432 \
    go run main.go
```
If using a debugger such as [delve](https://github.com/go-delve/delve), be sure to invoke the debugger as well:
```bash
POSTGRES_USER=postgres \
    POSTGRES_PASSWORD=changeme \
    POSTGRES_DB=postgres \
    POSTGRES_HOST=localhost\
    POSTGRES_PORT=5432 \
    dlv debug main.go
```
You'll also need to copy the `juices.txt` file over to `src/juices.txt` because that's where it exists in the container:
```bash
mkdir -p src/
cp juices.txt src/juices.txt
```
