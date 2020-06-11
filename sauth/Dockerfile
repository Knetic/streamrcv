FROM debian:9

WORKDIR /var/lib/sauth
EXPOSE 1930

COPY ./.output/sauth /usr/local/bin/sauth

CMD ["/usr/local/bin/sauth"]