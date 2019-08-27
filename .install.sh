#!/bin/bash
docker volume create --name postgres_data
docker run --rm --name my_postgres_container -e POSTGRESQL_DATABASE=go-tdd -e POSTGRES_USER=moe -e POSTGRES_PASSWORD=admin -v postgres_data:/var/lib/postgresql/data -p 5432:5432 -d bitnami/postgresql

