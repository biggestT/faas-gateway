#! /bin/bash

registry=eu.gcr.io
project=travtech
images=("faas-gateway" "rhymer" "reducer" "track")

for image in "${images[@]}"
do
  docker pull $registry/$project/$image:latest
done

docker exec api_controller docker-compose down 
docker run --rm \
  --name api_controller \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v "$PWD:$PWD" -w="$PWD" \
  docker/compose:1.24.0 up
