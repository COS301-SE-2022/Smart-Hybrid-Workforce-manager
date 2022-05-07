#!/bin/bash
set -e

USER="admin"
DBNAME="arche"
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "Using database $DBNAME"

CONFIG=" dbname=$DBNAME user=$USER"

######### example
for file in ${DIR}/example/*.schema.*sql
do
  psql "$CONFIG" -f "$file"
done

for file in ${DIR}/example/*.function.*sql
do
  psql "$CONFIG" -f "$file"
done

######### user
for file in ${DIR}/user/*.schema.*sql
do
  psql "$CONFIG" -f "$file"
done

for file in ${DIR}/user/*.function.*sql
do
  psql "$CONFIG" -f "$file"
done

######### team
for file in ${DIR}/team/*.schema.*sql
do
  psql "$CONFIG" -f "$file"
done

for file in ${DIR}/team/*.function.*sql
do
  psql "$CONFIG" -f "$file"
done

######### role
for file in ${DIR}/role/*.schema.*sql
do
  psql "$CONFIG" -f "$file"
done

for file in ${DIR}/role/*.function.*sql
do
  psql "$CONFIG" -f "$file"
done

######### permission
for file in ${DIR}/permission/*.schema.*sql
do
  psql "$CONFIG" -f "$file"
done

for file in ${DIR}/permission/*.function.*sql
do
  psql "$CONFIG" -f "$file"
done

######### bookings

######### resources