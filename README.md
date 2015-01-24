# sbuca

Simple But Useful Certificate Authority

## Installation

    go install github.com/gophergala/sbuca


## Usage

### Run the server

    mkdir ca certs
    sbuca server


### Issure Certificate

First, generate a certificate request

    openssl req -new -keyout myserver.key -out myserver.csr

And then send the csr by the restful api

    curl localhost:3000/certificates -XPOST --data-urlencode csr@myserver.csr > myserver.crt

Congrat, `myserver.crt` is the signed certification


## TODO

1. mkdir ca certs automatically if no exists
2. make the return format consistent (all json by default, use ?format=file to return a file
3. deploy a try server, try.sbuca.com, so that we can use the api with hosting our own CA
