streamrcv
=====

Accepts incoming multi-tenant authenticated RTMP livestreams, handles transcoding and HLS splitting, and serves the livestream with a simple web frontend.

Layout
====

* `streamrcv` - builds a container that uses [nginx-rtmp](https://github.com/arut/nginx-rtmp-module) to accept rtmp (port 1935), splits HLS (port 1936), and serves the segments and m3u8 files (port 8080).
* `sauth` - builds a container that checks authentication of users trying to publish rtmp.

The top-level `Makefile` will build both projects, and requires only Docker to do so.