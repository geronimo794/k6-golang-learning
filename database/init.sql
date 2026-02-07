CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- public.table_test definition

-- Drop table

-- DROP TABLE public.table_test;

CREATE TABLE public.table_test (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	data_test varchar(100) DEFAULT NULL::character varying NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	updated_at timestamptz DEFAULT now() NOT NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT table_test_pkey PRIMARY KEY (id)
);