streamrcv
=====

Makes it easy to host your own livestream, without needing popular platforms.

On a technical level; Accepts incoming multi-tenant authenticated RTMP livestreams, handles transcoding and HLS splitting, and serves the livestream with a simple web frontend.

## Why?

All of us take offline video and share it with impunity, but there's no such option for a livestream. This aims to provide a dead-simple server that does the work needed for regular people to make their own livestream.

Large video platforms often have stifling rules around what they will, and won't, host. They are (reasonably) concerned with losing advertiser revenue, or not being able to grow their audience because of a reputation. They're businesses, focused on growing the business. The user-facing goal of just _serving live video_ is secondary to making money from it. They're under no obligation to serve anyone's video, retain copies of it, or serve users equally.

## Using

Do `make` at the top level - this will build two containers. Afterwards, run the given `docker-compose.yml` to run the project.

## Layout

* `streamrcv` - builds a container that uses [nginx-rtmp](https://github.com/arut/nginx-rtmp-module) to accept rtmp (port 1935), splits HLS (port 1936), and serves the segments and m3u8 files (port 8080).
* `sauth` - builds a container that checks authentication of users trying to publish rtmp.

The top-level `Makefile` will build both projects, and requires only Docker to do so.