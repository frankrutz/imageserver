# imageserver


Golang webserver for the marketing department

# Installation on Ubuntu 16.04
apt-get update

apt-get install golang

apt-get install imagemagick imagemagick-doc 


# To run the webserver:
mkdir tmp
go build swiss.go
nohup ./swiss &
