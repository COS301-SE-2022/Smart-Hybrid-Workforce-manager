FROM nginx:latest

WORKDIR /etc/nginx
COPY ./proxy/nginx.conf ./conf.d/default.conf
EXPOSE 5000
ENTRYPOINT [ "nginx" ]
CMD [ "-g", "daemon off;" ]