FROM postgres:alpine

COPY ./schema.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]