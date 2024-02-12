test:
    docker run \
        -e POSTGRES_DB=pokedoc \
        -e POSTGRES_USER=admin \
        -e POSTGRES_PASSWORD=password123 \
        -p 5432:5432 \
        postgres:latest

build:
    docker-compose up -d --build postgres
    sleep 3;
    docker-compose up -d --build pokedoc-app

stop:
    docker-compose down

createpokedoc:
    docker build . -t pokedoc-app
