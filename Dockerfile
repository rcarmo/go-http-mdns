FROM gliderlabs/alpine
MAINTAINER rcarmo
WORKDIR /app
ADD web-linux /app/
ADD static /app/static
ADD views /app/views
EXPOSE 8000
CMD ["/app/web-linux"]