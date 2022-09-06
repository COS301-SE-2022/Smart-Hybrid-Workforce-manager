FROM nginx:latest

WORKDIR /etc/nginx
RUN mkdir ./ssl
COPY ./proxy/nginx.conf ./conf.d/default.conf
COPY ./proxy/Arche-Server.cert ./ssl/Arche-Server.cert
COPY ./proxy/Arche-Server.key ./ssl/Arche-Server.key
EXPOSE 8080
ENTRYPOINT [ "nginx" ]
CMD [ "-g", "daemon off;" ]