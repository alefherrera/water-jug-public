docker-build:
	docker build . -t water-jug:1.0

docker-run:
    docker run -p 8080:8080 water-jug:1.0