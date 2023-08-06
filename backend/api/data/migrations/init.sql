--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Debian 14.5-1.pgdg110+1)
-- Dumped by pg_dump version 14.5 (Homebrew)

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
-- Name: user; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user (
    id integer PRIMARY KEY NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255),
    role character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.user ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: -
--
-- password: secret
insert into public.user (first_name, last_name, email, password, role, created_at, updated_at) values ('Owner', 'Boss', 'owner@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy', 'OWNER', '2021-09-23 00:00:00', '2021-09-23 00:00:00');
insert into public.user (first_name, last_name, email, password, role, created_at, updated_at) values ('Driver', 'Employee', 'driver@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy', 'DRIVER', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
insert into public.user (first_name, last_name, email, password, role, created_at, updated_at) values ('User', 'Customer', 'user@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy', 'USER', '2023-09-23 00:00:00', '2023-09-23 00:00:00');


--
-- Name: restaurant; Type: TABLE; Schema: public; Owner: -
--
CREATE TABLE public.restaurant (
    id integer PRIMARY KEY NOT NULL,
    owner_id INTEGER REFERENCES public.user(id) ON DELETE CASCADE,
    name character varying(255),
    address character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Name: restaurant_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.restaurant ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.restaurant_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

--
-- Data for Name: restaurant; Type: TABLE DATA; Schema: public; Owner: -
--
insert into public.restaurant (owner_id, name, address, created_at, updated_at) values (1, 'Restaurant 1', '1 Tokyo', '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.restaurant (owner_id, name, address, created_at, updated_at) values (1, 'Restaurant 2', '2 Tokyo', '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.restaurant (owner_id, name, address, created_at, updated_at) values (1, 'Restaurant 3', '3 Tokyo', '2021-09-23 00:00:00', '2021-09-23 00:00:00');

--
-- Name: menu; Type: TABLE; Schema: public; Owner: -
--
CREATE TABLE public.menu (
    id integer PRIMARY KEY NOT NULL,
    restaurant_id INTEGER REFERENCES public.restaurant(id) ON DELETE CASCADE,
    name character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Name: menu_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.menu ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.menu_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

--
-- Data for Name: menu; Type: TABLE DATA; Schema: public; Owner: -
--
insert into public.menu (restaurant_id, name, created_at, updated_at) values (1, 'Menu 1', '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.menu (restaurant_id, name, created_at, updated_at) values (1, 'Menu 2', '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.menu (restaurant_id, name, created_at, updated_at) values (1, 'Menu 3', '2021-09-23 00:00:00', '2021-09-23 00:00:00');

--
-- Name: dish; Type: TABLE; Schema: public; Owner: -
--
CREATE TABLE public.dish (
    id integer PRIMARY KEY NOT NULL,
    menu_id INTEGER REFERENCES public.menu(id) ON DELETE CASCADE,
    name character varying(255),
    price NUMERIC(10, 2) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Name: dish_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.dish ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.dish_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

--
-- Name: order; Type: TABLE; Schema: public; Owner: -
--
CREATE TABLE public.order (
    id integer PRIMARY KEY NOT NULL,
    user_id INTEGER REFERENCES public.user(id) ON DELETE CASCADE,
    driver_id INTEGER REFERENCES public.user(id) ON DELETE CASCADE,
    amount NUMERIC(10, 2) NOT NULL,
    status character varying(255) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Name: order_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.order ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.order_dish (
    order_id INTEGER REFERENCES public.order(id) ON DELETE CASCADE,
    dish_id INTEGER REFERENCES public.dish(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);