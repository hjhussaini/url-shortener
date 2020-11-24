#/bin/bash

DATABASE_SERVER=localhost:9042 \
DATABASE_KEYSPACE=shortlink_db \
CACHE_SERVER=localhost:6379 \
HOST=localhost \
PORT=6100 \
./read-server
