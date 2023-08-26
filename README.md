## Postman Documentation API
- https://documenter.getpostman.com/view/26480274/2s9Y5YR2r5

## How to Run
### Run Postgres in Docker
- Pull postgres docker image 
```
docker pull postgres
```
- Build and run container
```
docker run --name my-postgres -e POSTGRES_USER=adminpostgres -e POSTGRES_PASSWORD=87654321 -e POSTGRES_DB=my_db -p 5433:5432 -d postgres
```

### Run Kafka in Docker
- Run `misc/docker-compose.yml`
```
docker-compose -f docker-compose.yml up
```
- Go into kafka `/bin/` folder in docker
```
docker exec -it kafka /bin/sh
```
and then
```
cd kafka_2.13-2.8.1/bin
```
- Create Kafka topics for the project inside the `/bin/` folder, `input-harga`, `topup`, and `buyback`
```
kafka-topics.sh --create --topic <topic-name> --bootstrap-server localhost:9092
```
