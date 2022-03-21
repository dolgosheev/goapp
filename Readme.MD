<h1>Docker build & start app</h1>

<h2>My repo</h2>
<pre>
    docker build . -t v0lshebnick/goapp:0.0.1
    docker push v0lshebnick/goapp:0.0.1

    docker run --rm -p 80:3001 v0lshebnick/goapp:0.0.1
</pre>