docker-compose rm
docker image rm icon-bsc_goloop
rm -rf work/*
#docker-compose -f docker-compose.yml -f docker-compose.provision.yml up -d   --force-recreate && docker-compose stop
#docker inspect iconbridge_src -f '{{ json .State.Health.Log }}' | jq .
docker-compose build
