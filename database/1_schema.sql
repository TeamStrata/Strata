-- DESCRIPTION
-- This script will setup the schema for the strata database.
-- 
-- EXAMPLE USAGE
-- psql.exe -U <postgres username> -d strata -f .\database\schema.sql

DO
$do$
BEGIN
	-- Users table
	IF EXISTS (
		SELECT 1
		FROM information_schema.tables
		WHERE table_schema = 'public'
		AND table_name = 'users'
	) THEN
		DROP TABLE public.users;
	END IF;

	CREATE TABLE users (
	user_id SERIAL PRIMARY KEY,
	user_name TEXT UNIQUE NOT NULL,
	password_hash TEXT NOT NULL 
	);
END
$do$