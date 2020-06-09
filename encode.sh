#!/bin/bash

app=$1
name=$2

echo "encoding for ${app}/${name}" >> /var/lib/streamrcv/data/hls/wup.txt

/usr/bin/ffmpeg -re \
    -i rtmp://localhost/${app}/${name} \
    -async 1 -vsync -1 \
    -c:v libx264 -c:a libvo_aacenc -b:v 1024k -b:a 128k -vf "scale=960:trunc(ow/a/2)*2" -tune zerolatency -preset veryfast -crf 23 -f flv rtmp://localhost/m3u8/${name}_sd \
    -c:v libx264 -c:a libvo_aacenc -b:v 1920k -b:a 128k -vf "scale=1280:trunc(ow/a/2)*2" -tune zerolatency -preset veryfast -crf 23 -f flv rtmp://localhost/m3u8/${name}_720p \
    -c copy -f flv rtmp://localhost/m3u8/${name}_src &> /var/lib/streamrcv/data/hls/ffmpeg.txt