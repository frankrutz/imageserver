# imageserver
Golang webserver for the marketing department

Tested and used on Google Cloud


## Installation on Ubuntu 16.04

webserver must allow http / https traffic

apt-get update

apt-get install golang

apt-get install imagemagick imagemagick-doc 

git clone https://github.com/frankrutz/imageserver

## To run the webserver:

cd imageserver

chmod u+x printimage.sh

mkdir tmp

go build swiss.go

nohup ./swiss &

