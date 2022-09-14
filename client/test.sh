#!/bin/bash

# Just to be sure that no instance is running previously
docker container rm -f dogstore-client

docker container run --name dogstore-client -d \
  --network wso-net \
  dogstore-client:0.0.1 "/client --url http://wso.sample:8280/dogstore/v1/owners --method GET --times 3000 --header \"Accept: application/json\""
