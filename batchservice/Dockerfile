FROM golang:1.14.6-alpine3.11
RUN mkdir /batchserviceapp
ADD . /batchserviceapp
WORKDIR /batchserviceapp/target/
RUN go build -o batchservice ../src/batchservice.go 
CMD ["/batchserviceapp/target/batchservice"]

