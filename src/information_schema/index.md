# Information Schema

- table_catalog:	def
- table_schma:		database name
- 

## Table Exists

	select count(1)
	from information_schema.tables
	where table_name = 'table_name' and table_catalog = 'db_name';
