#!/bin/bash
set -e

USER="admin"
DBNAME="arche"
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "Using database $DBNAME"

CONFIG=" dbname=$DBNAME user=$USER"

######### resource
for file in ${DIR}/resource/*.schema.*sql
do
  psql "$CONFIG" -f "$file"
done

for file in ${DIR}/resource/*.function.*sql
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

######### booking
for file in ${DIR}/booking/*.schema.*sql
do
  psql "$CONFIG" -f "$file"
done

for file in ${DIR}/booking/*.function.*sql
do
  psql "$CONFIG" -f "$file"
done

######### mock
for file in ${DIR}/mock/*.sql
do
  psql "$CONFIG" -f "$file"
done
