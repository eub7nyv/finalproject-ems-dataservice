# finalproject-dataservice
 
# Golang dataservice example
Build out payer dataservice in golang to show structure and containerization.

## Build the code
From the root dir:

`./scripts/build.sh ./dataservice`

## Build docker image, put it in the same network as Postgres
`docker build -t 'dataservice:latest' . --network=kafka

## Run it...
`docker run --rm dataservice:latest <action> <initial value> <values...>`


## Ingore below lines for now - TODO later.

# Based on...
# This is meant to be an idiomatic go implementation of [the java version](https://github.com/sxie2008/finalproject-ems-dataservice/tree/feature-Shelly-dataservice)

# License
This is licensed under Apache-2.0.


# local run command: 
go install -v dataservice/app.go dataservice/payer.go  dataservice/app_test.go dataservice/payer_test.go


go build -v dataservice/app.go dataservice/payer.go  dataservice/app_test.go dataservice/payer_test.go dataservice/db.go dataservice/pharmacy.go dataservice/handler.go dataservice/pharmacy_test.go

./app



## Postgres Log to see if server is OK. if below exists, run command to resolve it
OXC02XJ0BMJGH8:fullstack  ****** $ tail -n 10 /usr/local/var/log/postgres.log
2020-07-22 19:06:24.180 MST [11997] FATAL:  lock file "postmaster.pid" already exists
2020-07-22 19:06:24.180 MST [11997] HINT:  Is another postmaster (PID 817) running in data directory "/usr/local/var/postgres"?

rm -f /usr/local/var/postgres/postmaster.pid
