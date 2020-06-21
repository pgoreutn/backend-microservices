package main

const envTemplate = `# Basic env configuration

# Service Handlers
USE_HTTP=true
USE_GRPC=true
USE_GRAPHQL=true
USE_KAFKA_CONSUMER=true

HTTP_PORT=8000
GRPC_PORT=8001

BASIC_AUTH_USERNAME=user
BASIC_AUTH_PASS=pass

USE_MONGO=[bool]
MONGODB_HOST_WRITE=[string]
MONGODB_HOST_READ=[string]
MONGODB_DATABASE_NAME=[string]

USE_SQL=[bool]
SQL_DRIVER_NAME=[string]
SQL_DB_READ_HOST=[string]
SQL_DB_READ_USER=[string]
SQL_DB_READ_PASSWORD=[string]
SQL_DB_WRITE_HOST=[string]
SQL_DB_WRITE_USER=[string]
SQL_DB_WRITE_PASSWORD=[string]
SQL_DATABASE_NAME=[string]

USE_REDIS=[bool]
REDIS_READ_HOST=[string]
REDIS_READ_PORT=[string]
REDIS_READ_AUTH=[string]
REDIS_WRITE_HOST=[string]
REDIS_WRITE_PORT=[string]
REDIS_WRITE_AUTH=[string]

USE_RSA_KEY=[bool]

KAFKA_BROKERS=localhost:9092
KAFKA_CLIENT_ID=client_id
KAFKA_CONSUMER_GROUP=[string]


# Additional env

`