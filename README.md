# micro
go micro project

docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service

docker run -e MICRO_REGISTRY=mdns consignment-cli
