#!/bin/bash

# Just to be sure that no instance is running previously
docker container rm -f dogstore-client

# Consumer key generated in subscriptions page
CONSUMER_KEY=aFc9VNGjb2gWuPZegnxuCtdE5d4a
# Consumer secret generated in subscriptions page also
CONSUMER_SECRET=Q5aT0t7TRFToR1SPlHnJWu5NaXAa
BEARER=$(curl https://localhost:9443/oauth2/token --insecure -d "grant_type=client_credentials" -H "Authorization: Basic $(echo -n ${CONSUMER_KEY}:${CONSUMER_SECRET} | base64)" | jq -r '.access_token')

echo $BEARER

docker container run --name dogstore-client -d \
  --network wso-net \
  dogstore-client:0.0.1 /bin/sh -c "/client --url http://wso.sample:8280/dogstore/v1/owners --method GET --times 10 --header \"Accept: application/json\" --header \"Authorization: Bearer $BEARER\""

docker container logs dogstore-client

