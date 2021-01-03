#!/bin/sh

go build -o rsa main.go

# Measure system time, user time, real time and memory usage
# when running a simple program that generates an RSA key

echo "--------"

# with 1024 bits
./xtime.sh ./rsa --bits 1024

echo "--------"


# with 4096 bits
./xtime.sh ./rsa --bits 4096

echo "--------"