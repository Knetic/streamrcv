FROM debian:latest

# replace `1.16.1` with the latest version from https://nginx.org/en/download.html
# when updating

RUN apt-get update; \
    apt-get install -y \
    build-essential libpcre3 libpcre3-dev libssl-dev autoconf automake git wget curl zlib1g zlib1g-dev

RUN cd /tmp; \
    git clone https://github.com/arut/nginx-ts-module.git; \
    wget https://nginx.org/download/nginx-1.16.1.tar.gz; \
    git clone https://github.com/sergey-dryabzhinsky/nginx-rtmp-module.git;

RUN cd /tmp; \
    tar -xf nginx-1.16.1.tar.gz;

RUN cd /tmp/nginx-1.16.1/; \
    ./configure --with-http_ssl_module --add-module=../nginx-rtmp-module --with-http_stub_status_module --add-module=../nginx-ts-module; \
    make; \
    make install;

COPY index.html /var/lib/streamrcv/data/index.html
COPY nginx.conf /usr/local/nginx/conf/nginx.conf

EXPOSE 1935
EXPOSE 8080
VOLUME /var/lib/streamrcv/.htpasswd
WORKDIR /var/

CMD ["/usr/local/nginx/sbin/nginx"]