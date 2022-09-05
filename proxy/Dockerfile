FROM nginx:latest

WORKDIR /etc/nginx
COPY ./proxy/nginx.conf ./conf.d/default.conf
EXPOSE 8080
ENTRYPOINT [ "nginx" ]
CMD [ "-g", "daemon off;" ]