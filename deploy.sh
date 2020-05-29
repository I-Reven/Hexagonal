#!/bin/bash

echo Please select your deployment state? local, test, live, down

read state

if [ $state == "live" ]; then
docker-compose config > docker-compose-deploy.yaml && kompose convert -f docker-compose-deploy.yaml --out ./k8s
kubectl create -f ./k8s --save-config
#kubectl apply -f ./k8s
fi

if [ $state == "test" ]; then
docker-compose up -D mongodb elassandra redis rabbitmq
fi

if [ $state == "local" ]; then
docker-compose up -build
fi

if [ $state == "down" ]; then
kubectl delete -f ./k8s
docker-compose down
fi



