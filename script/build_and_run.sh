#!/bin/bash

echo "Building the docker image..."
docker build -t wordlesolver .

echo "Running wordlesolver..."
docker run -it --rm wordlesolver
