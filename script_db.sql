-- Table: public.flights

-- DROP TABLE IF EXISTS public.flights;

CREATE TABLE IF NOT EXISTS public.flights
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    airline_number integer,
    airline_code text COLLATE pg_catalog."default",
    destination text COLLATE pg_catalog."default",
    arrival text COLLATE pg_catalog."default",
    CONSTRAINT flight_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.flights
    OWNER to myuser;

    