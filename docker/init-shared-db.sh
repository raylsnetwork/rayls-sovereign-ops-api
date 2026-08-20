#!/bin/sh
set -e

until pg_isready -h shared-db -U blockscout; do
  echo "waiting for shared-db..."
  sleep 2
done

createdb -h shared-db -U blockscout ops_api  2>/dev/null && echo "created ops_api"  || echo "ops_api already exists"
createdb -h shared-db -U blockscout raylzdb  2>/dev/null && echo "created raylzdb" || echo "raylzdb already exists"
