docker run --name citus_standalone -d -p 5432:5432 -e POSTGRES_PASSWORD=password citusdata/citus

docker exec -it citus_standalone bash 

apt-get update -y

apt-get install -y postgresql-15-orafce

psql -U postgres

CREATE EXTENSION orafce;