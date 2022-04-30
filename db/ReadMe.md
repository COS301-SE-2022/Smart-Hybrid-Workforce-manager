# Database configuration

Docker will automatically run **sql/initdb.sh** which will create the schema and stored procedures.

## General Layout

Each schema and its' relating stored procedures should be stored in its own directory.

## Creating a schema file

All schema files should be labelled: `*.schema.*sql*` where \* is any character string.

The naming convention for schema's is as follows: `RELATIONNAME.schema.sql`

## Creating a stored procedure or function

All functions and procedures will be stored in their own respective file under the same directory as the schema file

The naming convention for functions is as follows: `RELATIONNAME.function.FUNCTIONNAME.sql`
