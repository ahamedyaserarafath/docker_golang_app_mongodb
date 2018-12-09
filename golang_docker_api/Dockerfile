## Simple docker file which will copy the current directory
## and download the respective golang and its package necessary to run the server
## finally it will run the simple perzis api server
FROM golang:1.8

WORKDIR /go/src/golang_api
COPY golang_api /go/src/golang_api/

RUN go-wrapper download   
RUN go-wrapper install    
RUN cd /go/src/golang_api/ && go build

CMD ["/go/src/golang_api/golang_api"]
