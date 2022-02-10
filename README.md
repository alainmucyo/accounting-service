# Accounting service

Accounting service is a microservice that is in charge of all financial information – company account balances, transactions, statements, Money push (debit) and pull (credit) is handed over to the payment integrators.

---

## Getting Started

1. git clone this repository & cd to the project directory

## Pre-requisites

* Golang v1.17 or greater
* Git
* Kafka
* VSCode, Goland or even any other code editor of your preferred choice.
* PostgresDB
* Redis

## Installing

* Install [Go](https://go.dev/doc/install) if you don't have it installed.

* Install [git](https://www.digitalocean.com/community/tutorials/how-to-contribute-to-open-source-getting-started-with-git)
  , (optional) if you dont have it installed.

* Install [Kafka](https://kafka.apache.org/) if you don't have it
* Install [postgresDB](https://www.postgresql.org/)

## Run the project

#### Using VSCode

1. Launch VSCode editor
2. Generate a launch json file by clicking on the `Run` button on the right side of your VSCode IDE
3. You can set the `PORT` you want to use in the `.env` file, if you don't set it the api will run on `3000` by default.
4. Congratulations! You have successfully launched Accounting service!

### Launch with Docker

> For this, you need to have [Docker](https://www.docker.com/) installed in your system.

1. Run `docker build -t <image-name> .` to build the docker image
2. Run `docker run -p 3000:3000 <image-name>` to run the image. This will expose port `3000`

### Launch with Docker compose

> For this, you need to have [Docker](https://www.docker.com/) and [Docker compose](https://docs.docker.com/compose/) installed in your system.

1. Run `docker-compose up` to build and run the docker images

### To check if the API is up and running.

Just call this endpoint: `http://localhost:3000/ping` using a GET method It will show a `pong` response.

Find the API docs `/api-docs` to get all API available

## Testing

Run `go test`

## Built With

* Golang v1.17
* [Segmentio Kafka libray](https://github.com/segmentio/kafka-go)
* [Gin Framework](https://github.com/gin-gonic/gin)
* [Redis](https://github.com/go-redis/redis)
* [GORM](https://gorm.io/index.html) ORM

## Authors

* **Alain MUCYO** (https://github.com/alainmucyo)

## Licence

This software is published by `FDI Engineering Team` under the [MIT licence](http://opensource.org/licenses/MIT).

