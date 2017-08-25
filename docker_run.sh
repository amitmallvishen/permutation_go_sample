CGO_ENABLED=0 GOOS=linux go build -o my_service  -a -installsuffix cgo -v services.go

docker rmi -f my_docker

docker build -t my_docker -f Docker_conf .

docker images

docker ps -a

#start the docker 

docker run -p 8080:8080 --name my_docker_ -d my_docker
