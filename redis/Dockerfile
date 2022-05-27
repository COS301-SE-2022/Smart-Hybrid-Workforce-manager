# COPIED FROM https://hub.docker.com/_/redis/
FROM redis:6.2.7

COPY redis.conf /usr/local/etc/redis/redis.conf

CMD [ "arche-redis", "/usr/local/etc/redis/redis.conf" ]
EXPOSE 6379