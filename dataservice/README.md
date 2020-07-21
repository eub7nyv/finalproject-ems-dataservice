# finalproject-dataservice
 
# Golang dataservice example
Build out payer dataservice in golang to show structure and containerization.

## Build the code
From the root dir:

`./scripts/build.sh ./dataservice`

## Build docker image, put it in the same network as Postgres
`docker build -t 'dataservice:latest' . --network=my-net

## Run it...
`docker run --rm dataservice:latest <action> <initial value> <values...>`


## Ingore below lines for now - TODO later.

# Based on...
# This is meant to be an idiomatic go implementation of [the java version](https://github.com/sxie2008/finalproject-ems-dataservice/tree/feature-Shelly-dataservice)

# License
This is licensed under Apache-2.0.