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
    id  SERIAL PRIMARY KEY,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255),
    role character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: -
--
-- password: secret
insert into public.user (first_name, last_name, email, password, role, created_at, updated_at) values ('OWNER', 'Boss', 'owner@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy', 'OWNER', '2021-09-23 00:00:00', '2021-09-23 00:00:00');
insert into public.user (first_name, last_name, email, password, role, created_at, updated_at) values ('DRIVER', 'Employee', 'driver@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy', 'DRIVER', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
insert into public.user (first_name, last_name, email, password, role, created_at, updated_at) values ('USER', 'Customer', 'user@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy', 'USER', '2023-09-23 00:00:00', '2023-09-23 00:00:00');


--
-- Name: restaurant; Type: TABLE; Schema: public; Owner: -
--
CREATE TABLE public.restaurant (
    id  SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES public.user(id) ON DELETE CASCADE,
    name character varying(255),
    address character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
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
    id  SERIAL PRIMARY KEY,
    restaurant_id INTEGER REFERENCES public.restaurant(id) ON DELETE CASCADE,
    name character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Data for Name: menu; Type: TABLE DATA; Schema: public; Owner: -
--
insert into public.menu (restaurant_id, name, created_at, updated_at) values (1, 'Menu 1', '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.menu (restaurant_id, name, created_at, updated_at) values (1, 'Menu 2', '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.menu (restaurant_id, name, created_at, updated_at) values (1, 'Menu 3', '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.menu (restaurant_id, name, created_at, updated_at) values (2, 'Menu 4', '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.menu (restaurant_id, name, created_at, updated_at) values (2, 'Menu 5', '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.menu (restaurant_id, name, created_at, updated_at) values (2, 'Menu 6', '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.menu (restaurant_id, name, created_at, updated_at) values (3, 'Menu 7', '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.menu (restaurant_id, name, created_at, updated_at) values (3, 'Menu 8', '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.menu (restaurant_id, name, created_at, updated_at) values (3, 'Menu 9', '2021-09-23 00:00:00', '2021-09-23 00:00:00');

--
-- Name: dish; Type: TABLE; Schema: public; Owner: -
--
CREATE TABLE public.dish (
    id  SERIAL PRIMARY KEY,
    menu_id INTEGER REFERENCES public.menu(id) ON DELETE CASCADE,
    name character varying(255),
    price NUMERIC(10, 2) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Data for Name: dish; Type: TABLE DATA; Schema: public; Owner: -
--
insert into public.dish (menu_id, name, price, created_at, updated_at) values (1, 'Dish 1', 3.00, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (1, 'Dish 2', 5.00, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (1, 'Dish 3', 7.00, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.dish (menu_id, name, price, created_at, updated_at) values (2, 'Dish 4', 3.00, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (2, 'Dish 5', 5.00, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (2, 'Dish 6', 7.00, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.dish (menu_id, name, price, created_at, updated_at) values (3, 'Dish 7', 3.00, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (3, 'Dish 8', 5.00, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (3, 'Dish 9', 7.00, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.dish (menu_id, name, price, created_at, updated_at) values (4, 'Dish 10', 3.00, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (4, 'Dish 11', 5.00, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (4, 'Dish 12', 7.00, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.dish (menu_id, name, price, created_at, updated_at) values (5, 'Dish 13', 3.00, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (5, 'Dish 14', 5.00, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (5, 'Dish 15', 7.00, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.dish (menu_id, name, price, created_at, updated_at) values (6, 'Dish 16', 3.00, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (6, 'Dish 17', 5.00, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (6, 'Dish 18', 7.00, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.dish (menu_id, name, price, created_at, updated_at) values (7, 'Dish 19', 3.00, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (7, 'Dish 20', 5.00, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (7, 'Dish 21', 7.00, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.dish (menu_id, name, price, created_at, updated_at) values (8, 'Dish 22', 3.00, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (8, 'Dish 23', 5.00, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (8, 'Dish 24', 7.00, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

insert into public.dish (menu_id, name, price, created_at, updated_at) values (9, 'Dish 25', 3.00, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (9, 'Dish 26', 5.00, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.dish (menu_id, name, price, created_at, updated_at) values (9, 'Dish 27', 7.00, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

--
-- Name: order; Type: TABLE; Schema: public; Owner: -
--
CREATE TABLE public.order (
    id  SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES public.user(id) ON DELETE CASCADE,
    driver_id INTEGER REFERENCES public.user(id) ON DELETE CASCADE,
    amount NUMERIC(10, 2) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Data for Name: order; Type: TABLE DATA; Schema: public; Owner: -
--
insert into public.order (user_id, driver_id, amount, created_at, updated_at) values (1, 2, 9.99, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.order (user_id, driver_id, amount, created_at, updated_at) values (1, 2, 11.99, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.order (user_id, driver_id, amount, created_at, updated_at) values (1, 2, 13.99, '2021-09-23 00:00:00', '2021-09-23 00:00:00');

CREATE TABLE public.order_dish (
    order_id INTEGER REFERENCES public.order(id) ON DELETE CASCADE,
    dish_id INTEGER REFERENCES public.dish(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Data for Name: order_dish; Type: TABLE DATA; Schema: public; Owner: -
--
insert into public.order_dish (order_id, dish_id, quantity, created_at, updated_at) values (1, 1, 3, '2019-09-23 00:00:00', '2019-09-23 00:00:00');
insert into public.order_dish (order_id, dish_id, quantity, created_at, updated_at) values (1, 2, 5, '2020-09-23 00:00:00', '2020-09-23 00:00:00');
insert into public.order_dish (order_id, dish_id, quantity, created_at, updated_at) values (1, 3, 7, '2021-09-23 00:00:00', '2021-09-23 00:00:00');