

  flyway:
    image: flyway/flyway:6.4.1-alpine
    environment:
      - FLYWAY_EDITION=community
    command: -url=jdbc:postgresql://db:5432/flyway_teste -schemas=flyway -user=flyway -password=flyway -connectRetries=10 migrate


Error:__Database__is__uninitialized__and__superuser__password__is__not__specified.
        You__must__specify__POSTGRES_PASSWORD__to__a__non-empty__value__for__the
        superuser.__For__example,__"-e__POSTGRES_PASSWORD=password"__on__"docker__run".

        You__may__also__use__"POSTGRES_HOST_AUTH_METHOD=trust"__to__allow__all
        connections__without__a__password.__This__is__*not*__recommended.

        See__PostgreSQL__documentation__about__"trust":
0https://www.postgresql.org/docs/current/auth-trust.html







The files belonging to this database system will be owned by user "postgres".
This user must also own the server process.

The database cluster will be initialized with locale "en_US.utf8".
The default database encoding has accordingly been set to "UTF8".
The default text search configuration will be set to "english".

Data page checksums are disabled.

fixing permissions on existing directory /var/lib/postgresql/data ... ok
creating subdirectories ... ok
selecting dynamic shared memory implementation ... posix
selecting default max_connections ... 100
selecting default shared_buffers ... 128MB
selecting default time zone ... UTC
creating configuration files ... ok
running bootstrap script ... ok
performing post-bootstrap initialization ... 
2022-04-14 03:10:39.142 UTC [31] FATAL:  zero-length delimited identifier at or near """" at character 12
2022-04-14 03:10:39.142 UTC [31] STATEMENT:  ALTER USER ""flyway"" WITH PASSWORD E'"flyway"';
	
child process exited with exit code 1
initdb: removing contents of data directory "/var/lib/postgresql/data"