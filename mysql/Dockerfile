FROM mysql:8.2

ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes
ENV MYSQL_DATABASE=shisha-log

RUN echo "USE mysql;" > /docker-entrypoint-initdb.d/timezones.sql &&  mysql_tzinfo_to_sql /usr/share/zoneinfo >> /docker-entrypoint-initdb.d/timezones.sql

EXPOSE 3306

COPY my.cnf /etc/mysql/my.cnf