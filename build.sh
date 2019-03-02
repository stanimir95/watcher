#!/bin/bash
echo "Please input project directory"
read directory
go build -o watch

docker build -t watcher . 

docker run -v $directory:/tmp watcher
