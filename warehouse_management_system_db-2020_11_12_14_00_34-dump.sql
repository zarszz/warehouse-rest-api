--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0 (Debian 13.0-1.pgdg100+1)
-- Dumped by pg_dump version 13.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: warehouse; Type: DATABASE; Schema: -; Owner: ucok
--

CREATE DATABASE warehouse WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE warehouse OWNER TO ucok;

\connect warehouse

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: ucok
--

CREATE TABLE public.categories (
    id character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    category text
);


ALTER TABLE public.categories OWNER TO ucok;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: ucok
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.categories_id_seq OWNER TO ucok;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ucok
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: items; Type: TABLE; Schema: public; Owner: ucok
--

CREATE TABLE public.items (
    id text NOT NULL,
    item_name text,
    category_id character varying,
    warehouse_id text,
    owner_id text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    rack_id text,
    description text,
    room_id text
);


ALTER TABLE public.items OWNER TO ucok;

--
-- Name: racks; Type: TABLE; Schema: public; Owner: ucok
--

CREATE TABLE public.racks (
    id character varying NOT NULL,
    room_id character varying NOT NULL,
    name character varying NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.racks OWNER TO ucok;

--
-- Name: rooms; Type: TABLE; Schema: public; Owner: ucok
--

CREATE TABLE public.rooms (
    id character varying NOT NULL,
    warehouse_id character varying,
    name character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone
);


ALTER TABLE public.rooms OWNER TO ucok;

--
-- Name: user_addresses; Type: TABLE; Schema: public; Owner: ucok
--

CREATE TABLE public.user_addresses (
    address_id character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id character varying,
    address text,
    city text,
    state text,
    country text,
    postal_code text
);


ALTER TABLE public.user_addresses OWNER TO ucok;

--
-- Name: users; Type: TABLE; Schema: public; Owner: ucok
--

CREATE TABLE public.users (
    id character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    is_admin boolean,
    email text,
    password text,
    first_name text,
    last_name text,
    gender text,
    date_of_birth timestamp with time zone
);


ALTER TABLE public.users OWNER TO ucok;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: ucok
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO ucok;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ucok
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: warehouse_addresses; Type: TABLE; Schema: public; Owner: ucok
--

CREATE TABLE public.warehouse_addresses (
    id character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    warehouse_id text,
    address text,
    province text,
    city text,
    country text,
    postal_code text,
    state text
);


ALTER TABLE public.warehouse_addresses OWNER TO ucok;

--
-- Name: warehouse_addresses_id_seq; Type: SEQUENCE; Schema: public; Owner: ucok
--

CREATE SEQUENCE public.warehouse_addresses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.warehouse_addresses_id_seq OWNER TO ucok;

--
-- Name: warehouse_addresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ucok
--

ALTER SEQUENCE public.warehouse_addresses_id_seq OWNED BY public.warehouse_addresses.id;


--
-- Name: warehouses; Type: TABLE; Schema: public; Owner: ucok
--

CREATE TABLE public.warehouses (
    id text NOT NULL,
    warehouse_name text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.warehouses OWNER TO ucok;

--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: ucok
--

COPY public.categories (id, created_at, updated_at, deleted_at, category) FROM stdin;
0	2020-10-27 00:27:51.001181+00	2020-10-27 00:27:51.001181+00	\N	barang random 1
c54ce48bc146b576f1be1a53cfc4bf2da439585e	2020-10-27 00:30:39.424685+00	2020-10-27 00:30:39.424685+00	\N	barang random 1
3cd9e28d95d31270e106eadf3bade3b8d0d77d0c	2020-10-30 07:39:08.849303+00	2020-10-30 07:39:08.849303+00	\N	barang random 1
4d67ac62e320ea5a2a811ca3f8310ca82d6a325d	2020-10-30 07:39:10.123499+00	2020-10-30 07:39:10.123499+00	\N	barang random 1
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: ucok
--

COPY public.items (id, item_name, category_id, warehouse_id, owner_id, created_at, updated_at, deleted_at, rack_id, description, room_id) FROM stdin;
6cf10dabeda5912d92102daa499b97bab27cb298	Ucok Monitor 2	254cb40ec505023784557766048f7184f244fbe0	c00e76db68a780053d5c2490ec6d6a9d120c3eaf	157cebb033f4362eeff6f2cc16ab198e6914a2ed	2020-11-11 01:39:08.136197+00	2020-11-11 01:39:08.136197+00	\N	4b0b32c962af5452163c0503d595b9394a8e4194	ucok monitor 2 description	254cb40ec505023784557766048f7184f244fbe0
2a7141da2fe01ceb9c25ce058c65ba04ac0ac5c6	Ucok Monitor 2	254cb40ec505023784557766048f7184f244fbe0	c00e76db68a780053d5c2490ec6d6a9d120c3eaf	157cebb033f4362eeff6f2cc16ab198e6914a2ed	2020-11-08 07:34:01.6705+00	2020-11-08 07:34:01.6705+00	\N	4b0b32c962af5452163c0503d595b9394a8e4194	ucok monitor 2 description	254cb40ec505023784557766048f7184f244fbe0
6d6c5cdde4e95c10ab769675bc7543a2a266022a	Ucok Monitor 2	254cb40ec505023784557766048f7184f244fbe0	c00e76db68a780053d5c2490ec6d6a9d120c3eaf	157cebb033f4362eeff6f2cc16ab198e6914a2ed	2020-11-11 01:39:08.14307+00	2020-11-11 01:39:08.14307+00	\N	4b0b32c962af5452163c0503d595b9394a8e4194	ucok monitor 2 description	254cb40ec505023784557766048f7184f244fbe0
\.


--
-- Data for Name: racks; Type: TABLE DATA; Schema: public; Owner: ucok
--

COPY public.racks (id, room_id, name, created_at, updated_at, deleted_at) FROM stdin;
4b0b32c962af5452163c0503d595b9394a8e4194	254cb40ec505023784557766048f7184f244fbe0	Rack 1 at Ucok room 001	2020-11-08 03:35:10.877802+00	2020-11-08 03:35:10.877802+00	\N
39de3985d04303cac303adf95467c2719764cb43	254cb40ec505023784557766048f7184f244fbe0	Rack 2	2020-11-08 03:35:34.832695+00	2020-11-08 03:35:34.832696+00	\N
fef11cf5917c2f5bb7222deae91ba09ea0b1672a	254cb40ec505023784557766048f7184f244fbe0	Rack 3	2020-11-08 03:35:39.906845+00	2020-11-08 03:35:39.906846+00	\N
da343f4f2ee4cbaadccb719c7c1008faad9a8cd9	254cb40ec505023784557766048f7184f244fbe0	Rack 4	2020-11-08 03:35:43.550421+00	2020-11-08 03:35:43.550422+00	\N
\.


--
-- Data for Name: rooms; Type: TABLE DATA; Schema: public; Owner: ucok
--

COPY public.rooms (id, warehouse_id, name, created_at, updated_at) FROM stdin;
254cb40ec505023784557766048f7184f244fbe0	c00e76db68a780053d5c2490ec6d6a9d120c3eaf	room 001 AT UCOK WAREHOUSE 001	2020-11-08 03:27:55.650787+00	2020-11-08 03:27:55.650787+00
c1d3a04db1c040a77327a741b02ef8659c03f3b9	c00e76db68a780053d5c2490ec6d6a9d120c3eaf	room 002 AT UCOK WAREHOUSE 001	2020-11-08 03:27:59.417838+00	2020-11-08 03:27:59.417838+00
f7799fe9a5c1fc4e12ff8cf89d1e73cc7fdb1091	c00e76db68a780053d5c2490ec6d6a9d120c3eaf	room 003 AT UCOK WAREHOUSE 001	2020-11-08 03:28:02.629487+00	2020-11-08 03:28:02.629487+00
7fa483c0436401ea924c4b25290a24a789d6c763	c00e76db68a780053d5c2490ec6d6a9d120c3eaf	room 004 AT UCOK WAREHOUSE 001	2020-11-08 03:28:06.894854+00	2020-11-08 03:28:06.894854+00
\.


--
-- Data for Name: user_addresses; Type: TABLE DATA; Schema: public; Owner: ucok
--

COPY public.user_addresses (address_id, created_at, updated_at, deleted_at, user_id, address, city, state, country, postal_code) FROM stdin;
2e2b7bd75ffb7574a9fa246096be0464b2a81e43	2020-10-30 02:36:01.325141+00	2020-10-30 02:36:01.325141+00	\N	e3d7af444d9329eff7607bd3aa2e81acfb5ff340	cipaku	nagayawa	ciendog	isekai	40323
6b818d8e6f7a9ea25e0f04db7d913c0e09079c27	2020-10-30 07:37:04.117651+00	2020-10-30 07:37:04.117651+00	\N	157cebb033f4362eeff6f2cc16ab198e6914a2ed	cipaku	nagayawa	ciendog	isekai	40323
b7a6b7079ace10ea2403b35ea1d32cf5d0c44025	2020-10-30 07:38:18.44394+00	2020-10-30 07:38:18.443941+00	\N	010f11e24c8d239fc119ef6a278e0bea6eeb4827	cipaku	nagayawa	ciendog	isekai	40323
c3e99cc40167b03d21f21812245026c1a3cbf794	2020-11-08 03:30:05.440292+00	2020-11-08 03:30:05.440292+00	\N	6019dfaf0e98d950ae614b206d9fc13ed1b92343	cipaku	nagayawa	ciendog	isekai	40323
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: ucok
--

COPY public.users (id, created_at, updated_at, deleted_at, is_admin, email, password, first_name, last_name, gender, date_of_birth) FROM stdin;
1114738d76dac7f49a8954d87a493be4b98cfdef	2020-10-27 01:59:56.578755+00	2020-10-27 01:59:56.578756+00	\N	\N	ucok@email.com	password	ucok 12	ganteng	M	1990-02-10 00:00:00+00
157cebb033f4362eeff6f2cc16ab198e6914a2ed	2020-10-30 07:36:06.845812+00	2020-10-30 07:37:41.679689+00	\N	\N	udin@email.com3	password	udin gt update	ganteng gt update	M	1999-01-20 00:00:00+00
010f11e24c8d239fc119ef6a278e0bea6eeb4827	2020-10-30 07:38:03.241243+00	2020-10-30 07:38:03.241244+00	\N	\N	ucok@email.com	password	ucok 12	ganteng	M	1990-02-10 00:00:00+00
6019dfaf0e98d950ae614b206d9fc13ed1b92343	2020-11-08 03:29:52.696262+00	2020-11-08 03:29:52.696263+00	\N	\N	ucok@email.com	password	ucok 12	ganteng	M	1990-02-10 00:00:00+00
\.


--
-- Data for Name: warehouse_addresses; Type: TABLE DATA; Schema: public; Owner: ucok
--

COPY public.warehouse_addresses (id, created_at, updated_at, deleted_at, warehouse_id, address, province, city, country, postal_code, state) FROM stdin;
b095865359774ee81872af0e35f48611d785295b	2020-11-08 03:27:24.377417+00	2020-11-08 03:27:24.377417+00	\N	c00e76db68a780053d5c2490ec6d6a9d120c3eaf	cipaku	\N	nagayawa	isekai	40323	ciendog
\.


--
-- Data for Name: warehouses; Type: TABLE DATA; Schema: public; Owner: ucok
--

COPY public.warehouses (id, warehouse_name, created_at, updated_at, deleted_at) FROM stdin;
WR002	WAREHOUSE UCOK 002	2020-10-15 07:12:51.642822+00	2020-10-15 07:13:11.750653+00	\N
bcafcd467555fdf8c5e5e9f15da176ffd945ab88		2020-11-01 07:35:56.113073+00	2020-11-01 07:35:56.113073+00	\N
c00e76db68a780053d5c2490ec6d6a9d120c3eaf	WAREHOUSE UCOK 001	2020-11-08 03:27:07.875643+00	2020-11-08 03:27:07.875643+00	\N
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ucok
--

SELECT pg_catalog.setval('public.categories_id_seq', 2, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ucok
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: warehouse_addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ucok
--

SELECT pg_catalog.setval('public.warehouse_addresses_id_seq', 1, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: racks racks_pkey; Type: CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.racks
    ADD CONSTRAINT racks_pkey PRIMARY KEY (id);


--
-- Name: rooms rooms_pkey; Type: CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.rooms
    ADD CONSTRAINT rooms_pkey PRIMARY KEY (id);


--
-- Name: user_addresses user_addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.user_addresses
    ADD CONSTRAINT user_addresses_pkey PRIMARY KEY (address_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: warehouse_addresses warehouse_address_id_key; Type: CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.warehouse_addresses
    ADD CONSTRAINT warehouse_address_id_key UNIQUE (warehouse_id);


--
-- Name: warehouse_addresses warehouse_addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.warehouse_addresses
    ADD CONSTRAINT warehouse_addresses_pkey PRIMARY KEY (id);


--
-- Name: warehouses warehouses_pkey; Type: CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.warehouses
    ADD CONSTRAINT warehouses_pkey PRIMARY KEY (id);


--
-- Name: idx_categories_deleted_at; Type: INDEX; Schema: public; Owner: ucok
--

CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);


--
-- Name: idx_items_deleted_at; Type: INDEX; Schema: public; Owner: ucok
--

CREATE INDEX idx_items_deleted_at ON public.items USING btree (deleted_at);


--
-- Name: idx_user_addresses_deleted_at; Type: INDEX; Schema: public; Owner: ucok
--

CREATE INDEX idx_user_addresses_deleted_at ON public.user_addresses USING btree (deleted_at);


--
-- Name: idx_user_addresses_user_id; Type: INDEX; Schema: public; Owner: ucok
--

CREATE INDEX idx_user_addresses_user_id ON public.user_addresses USING btree (user_id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: ucok
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: idx_warehouse_addresses_deleted_at; Type: INDEX; Schema: public; Owner: ucok
--

CREATE INDEX idx_warehouse_addresses_deleted_at ON public.warehouse_addresses USING btree (deleted_at);


--
-- Name: idx_warehouses_deleted_at; Type: INDEX; Schema: public; Owner: ucok
--

CREATE INDEX idx_warehouses_deleted_at ON public.warehouses USING btree (deleted_at);


--
-- Name: items owner_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT owner_id_fkey FOREIGN KEY (owner_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: items racks_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT racks_id_fkey FOREIGN KEY (rack_id) REFERENCES public.racks(id) ON DELETE CASCADE;


--
-- Name: racks room_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.racks
    ADD CONSTRAINT room_id_fkey FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE;


--
-- Name: warehouse_addresses warehouse_address_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.warehouse_addresses
    ADD CONSTRAINT warehouse_address_fkey FOREIGN KEY (warehouse_id) REFERENCES public.warehouses(id) ON DELETE CASCADE;


--
-- Name: rooms warehouse_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ucok
--

ALTER TABLE ONLY public.rooms
    ADD CONSTRAINT warehouse_fkey FOREIGN KEY (warehouse_id) REFERENCES public.warehouses(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

