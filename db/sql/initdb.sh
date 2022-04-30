#!/bin/bash
set -e

USER="admin"
DBNAME="arche"
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "Using database $DBNAME"

CONFIG=" dbname=$DBNAME user=$USER"

for file in ${DIR}/example/*.schema.*sql
do
  psql "$CONFIG" -f "$file"
done

for file in ${DIR}/example/*.function.*sql
do
  psql "$CONFIG" -f "$file"
done