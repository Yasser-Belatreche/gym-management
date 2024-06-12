#!/bin/bash

set -e

# RabbitMQ
kubectl apply -f ./k8s/rabbitmq/rabbitmq-pv.yaml
kubectl apply -f ./k8s/rabbitmq/rabbitmq-pvc.yaml
kubectl apply -f ./k8s/rabbitmq/rabbitmq-secret.yaml
kubectl apply -f ./k8s/rabbitmq/rabbitmq.yaml

# Service Discovery
kubectl apply -f ./k8s/service-discovery/service-discovery-redis-pv.yaml
kubectl apply -f ./k8s/service-discovery/service-discovery-redis-pvc.yaml
kubectl apply -f ./k8s/service-discovery/service-discovery-secret.yaml
kubectl apply -f ./k8s/service-discovery/service-discovery-redis.yaml
kubectl apply -f ./k8s/service-discovery/service-discovery.yaml

# Api Gateway
kubectl apply -f ./k8s/api-gateway/api-gateway-secret.yaml
kubectl apply -f ./k8s/api-gateway/api-gateway-configmap.yaml
kubectl apply -f ./k8s/api-gateway/api-gateway.yaml

# Auth Service
kubectl apply -f ./k8s/auth-service/auth-service-psql-pv.yaml
kubectl apply -f ./k8s/auth-service/auth-service-psql-pvc.yaml
kubectl apply -f ./k8s/auth-service/auth-service-secret.yaml
kubectl apply -f ./k8s/auth-service/auth-service-configmap.yaml
kubectl apply -f ./k8s/auth-service/auth-service-psql.yaml
#kubectl apply -f ./k8s/auth-service/auth-service.yaml

# Gyms Service
kubectl apply -f ./k8s/gyms-service/gyms-service-psql-pv.yaml
kubectl apply -f ./k8s/gyms-service/gyms-service-psql-pvc.yaml
kubectl apply -f ./k8s/gyms-service/gyms-service-secret.yaml
kubectl apply -f ./k8s/gyms-service/gyms-service-configmap.yaml
kubectl apply -f ./k8s/gyms-service/gyms-service-psql.yaml
#kubectl apply -f ./k8s/gyms-service/gyms-service.yaml

# Memberships Service
kubectl apply -f ./k8s/memberships-service/memberships-service-psql-pv.yaml
kubectl apply -f ./k8s/memberships-service/memberships-service-psql-pvc.yaml
kubectl apply -f ./k8s/memberships-service/memberships-service-secret.yaml
kubectl apply -f ./k8s/memberships-service/memberships-service-configmap.yaml
kubectl apply -f ./k8s/memberships-service/memberships-service-psql.yaml
#kubectl apply -f ./k8s/memberships-service/memberships-service.yaml
