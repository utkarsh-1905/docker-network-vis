sudo docker run -d -p 6904:3000 --rm --net web --name docker-vis -v /var/run/docker.sock:/var/run/docker.sock docker-vis:latest