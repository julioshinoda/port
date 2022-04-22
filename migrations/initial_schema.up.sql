CREATE TABLE public.port (
	id varchar NOT NULL,
	"name" varchar NULL,
	city varchar NULL,
	country varchar NULL,
	alias _varchar NULL,
	regions _varchar NULL,
	coordinates _float8 NULL,
	province varchar NULL,
	timezone varchar NULL,
	unlocs _varchar NULL,
	code varchar NULL,
	CONSTRAINT port_pk PRIMARY KEY (id)
);
