#!/bin/bash

echo hello world

# variables; customary all caps
NAME=dog

# echo this is 1st args $0
# echo this is 2nd args $1
# echo this is 3rd args $2

# echo $NAME cat
# echo "$NAME cat"

echo $1
if [[ $1 == "dog" ]]; then echo "this is a dog";
else echo "this is not a dog";
fi