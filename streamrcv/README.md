streamrcv
====

Receives RTMP video, splits it into fragments, generates HLS master/variant, and serves the streams statically.

## How to Use

The project must be run as a docker container. The following line;

```
docker build -t streamrcv .
docker run -p 1935:1935 -p 8080:8080 streamrcv
```

Will start up an RTMP listener on the standard RTMP port (1935) and a basic HTTP server on 8080, which provides a simple front-end and serves the actual livestream to users.

It is highly recommended to mount `/var/lib/streamrcv/data/hls` and `/var/lib/streamrcv/data/rec` to your local drive. `hls` will contain the fragments and master/variant files, and is automatically trimmed. `rec` contains unedited video as it was sent to the server.

## References

https://github.com/arut/nginx-rtmp-module/wiki/Directives