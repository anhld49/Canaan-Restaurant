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
    id integer NOT NULL,
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

