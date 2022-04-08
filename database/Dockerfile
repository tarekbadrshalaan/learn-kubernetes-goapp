# set the base docker image for subsequent instructions.
FROM postgres

# specify the log level of database container as environment variable(used in both in build time and in runtime).
ENV POSTGRES_LOG_LEVEL=error

COPY script.sql /docker-entrypoint-initdb.d/

# docker build -t goapp-database .  
# docker run --rm -it -p 5000:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres goapp-database
# docker run --rm -it -p 5000:5432 goapp-database
# psql -d goapp -U postgres