# docker-application-practice

## init setting

* make docker image
<pre>
  cd main
  docker build --rm -t web .
</pre>
* check docker image list
<pre>
  docker images
</pre>
* start container
<pre>
  docker run -p 3030:8080 --name="test" -d web
</pre>
* test application
<pre>
  localhost:3030/World
</pre>
* stop container
<pre>
  docker stop test
</pre>
## use docker-compose

* run 4 application port 3031 - 3034
* 1 nginx port 5000
<pre>
  docker-compose up -d
</pre>
* stop compose
<pre>
  docker-compose down
</pre>
