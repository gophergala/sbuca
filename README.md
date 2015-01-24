# sbuca

Simple But Useful Certificate Authority


## Quick Start

Let's tried the hosted sbuca server

Generate rsakey & certificate request

    openssl genrsa -out server.key
    openssl req -new -key server.key -out server.csr

Send a Post request to the sbuca server, you'll get the signed cert file

    curl try.sbuca.com:3000/certificates -XPOST --data-urlencode csr@server.csr > server.crt


## Installation

    go get github.com/gophergala/sbuca

## Usage

### Init the server

    mkdir ca certs
    openssl genrsa -out ca/ca.key 2048
    openssl req -x509 -new -key ca/ca.key -out ca/ca.crt

### Run the server

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
