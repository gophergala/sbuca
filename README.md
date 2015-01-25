# sbuca

Simple But Useful Certificate Authority


## Installation

    go get github.com/gophergala/sbuca


## Quick Start

Let's tried the hosted sbuca server

Generate a Rsa Key

    sbuca genkey > test.key

Generate a Certificate Request

    subca gencsr --key test.key > test.csr

Submit the Certificate Request to the hosted server and get the Certificate

    subca submitcsr --host try.sbuca.com:8600 > test.crt

In case you want to get the Certificate in another computer, you can add `--format id`, then the output will become the id (serial number) of the Certificate

    subca submitcsr --host try.sbuca.com:8600 --format id 

Then you can get the certificate in another computer

    subca getcrt --host try.sbuca.com:8600 [ID] > test.crt

To get CA's certificate

    subca getcacrt --host try.sbuca.com:8600 > ca.crt


## Usage

### Run the CA server

To run a CA server, you can use

    sbuca server

It'll generate ca/ca.srl, ca/ca.key, and ca/ca.crt if needed.
The server listens to `0.0.0.0:8600` by default.


If you want to generate the key & certiricate by your own:

    mkdir ca certs
    echo 01 > ca/ca.srl
    openssl genrsa -out ca/ca.key 2048
    openssl req -x509 -new -key ca/ca.key -out ca/ca.crt
    sbuca server


### Generate a RSA Key

    sbuca genkey > test.key

This command is same as

    openssl genrsa -out test.key 2048


### Generate a Certification Request

    sbuca gencsr --key test.key > test.csr

This command is same as

    openssl req -new -key test.key -out test.csr


### Submit the Certification Request to the CA

By default, it'll output the signed Certificate to STDIN 

    sbuca submitcsr --host localhost:8600 test.csr > test.crt

If you want to get the ID instead, you can add `--format id`

    sbuca submitcsr --host localhost:8600 --format id test.csr > test.crt

We can use this `id` to get the certificate in another computer

In case you want to use curl to submit the csr, it'll output a JSON by default

    curl localhost:8600/certificates -XPOST --data-urlencode csr@test.csr

If you want to download the Certificate instead of the JSON, you can add `?format=file`

    curl localhost:8600/certificates -XPOST --data-urlencode csr@test.csr > test.crt


### Get the Certificate by ID

    sbuca getcrt --host localhost:8600 [ID] > test.crt

You can also use curl to get the Certificate (I use ID=20 as example)

    curl localhost:8600/certificates/20?format=file > test.crt

Congrat, `myserver.crt` is the signed certification


### Get CA's Certificate

    sbuca getcacrt --host localhost:8600 > ca.crt

You can also use curl to get it

    curl localhost:8600/ca/certificate?format=file > ca.crt


## TODO

1. password protection
2. admin fuctions: delete, delete all, get all, reset, ...
3. Dockerfile
