docker run --name oracle -d -p 1521:1521 -e ORACLE_PASSWORD=password gvenzl/oracle-xe:21.3.0

docker exec -it oracle sqlplus
system root