#!/bin/bash

app=$1
name=$2

/usr/bin/ffmpeg -re \
    -i rtmp://localhost/${app}/${name} \
    -async 1 -vsync -1 \
    -c:v libx264 -c:a aac -b:v 1920k -b:a 128k -vf "scale=1280:trunc(ow/a/2)*2" -tune zerolatency -preset veryfast -crf 23 -f flv rtmp://localhost/m3u8/${name}_720p \
    -c copy -f flv rtmp://localhost/m3u8/${name}_src &> /var/lib/streamrcv/data/hls/ffmpeg.txt