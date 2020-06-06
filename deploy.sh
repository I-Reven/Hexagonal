#!/bin/bash

echo Please select your deployment state? local, test, live
read state

if [ $state == "live" ]; then
  echo Please select your deployment state? up, down, build
  read state

  if [ $state == "build" ]; then
    docker-compose up -d --build
    docker push koushamad/hexagonal-core:latest
    docker push koushamad/hexagonal-grafana:latest

    docker push koushamad/hexagonal-icustomer:latest

    docker-compose config > docker-compose-deploy.yaml && kompose convert -f docker-compose-deploy.yaml --out ./k8s
    docker-compose down
    rm -rvf docker-compose-deploy.yaml
  fi

  if [ $state == "up" ]; then
    kubectl create -f ./k8s --save-config
  fi

  if [ $state == "down" ]; then
    kubectl delete -f ./k8s
  fi
fi

if [ $state == "test" ]; then
  docker-compose down
  docker-compose up -d mongodb elassandra redis rabbitmq core
fi

if [ $state == "local" ]; then
  docker-compose up --build  --remove-orphans
fi



