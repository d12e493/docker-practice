# docker-application-practice

## init setting

* make docker image
<pre>
  cd main
  docker build --rm -t web .
</pre>
* check docker image list

  docker images

* start container

  docker run -p 3030:8080 --name="test" -d web

* test application

  localhost:3030/World

* stop container

  docker stop test

## use docker-compose

* run 4 application port 3031 - 3034

  docker-compose up -d
