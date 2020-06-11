FROM debian:9

WORKDIR /var/lib/sauth
EXPOSE 8080

COPY ./.output/sauth /usr/local/bin/sauth

CMD ["/usr/local/bin/sauth"]