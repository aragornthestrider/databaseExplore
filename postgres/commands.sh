docker run -d --name postgres --network testnetwork -p 5432:5432 -h postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres postgres:15.3
