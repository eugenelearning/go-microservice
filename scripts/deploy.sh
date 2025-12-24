#! /bin/bash

kubectl apply -f ./k8s/cache/secrets.yml
kubectl apply -f ./k8s/cache/service.yml
kubectl apply -f ./k8s/cache/deploy.yml

kubectl apply -f ./k8s/db/volume.yml
kubectl apply -f ./k8s/db/secrets.yml
kubectl apply -f ./k8s/db/service.yml
kubectl apply -f ./k8s/db/deploy.yml

kubectl apply -f ./k8s/api/secrets.yml
kubectl apply -f ./k8s/api/service.yml
kubectl apply -f ./k8s/api/deploy.yml

kubectl apply -f ./k8s/ingress.yml