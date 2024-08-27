FROM postgres:13

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=simpleshop

COPY init.sql /docker-entrypoint-initdb.d/
