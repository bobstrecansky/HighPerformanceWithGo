FROM mysql:latest
ENV MYSQL_DATABASE=queues \
    MYSQL_ROOT_PASSWORD=admin
ADD schema.sql /docker-entrypoint-initdb.d
EXPOSE 3306
