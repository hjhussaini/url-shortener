#/bin/bash

DATABASE_SERVER=localhost:9042 \
DATABASE_KEYSPACE=key_db \
KEY_GENERATION_INTERVAL=1000 \
CACHE_SERVER=localhost:6379 \
HOST=localhost \
PORT=5100 \
./keygen-server
