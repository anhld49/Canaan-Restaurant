--
-- Name: user; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user (
    id integer NOT NULL PRIMARY KEY,
    name character varying(255),
    email character varying(255) unique,
    friends TEXT [],
    subscribe TEXT [],
    blocks TEXT [],
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

insert into public.user (name, email, friends, subscribe, created_at, updated_at) values ('Tom Nguyen', 'tom@test.com', ARRAY [ 'andrew@test.com','peter@test.com' ], ARRAY [ 'donald@test.com','peter@test.com' ], '2022-08-23 00:00:00', '2022-09-23 00:00:00');
insert into public.user (name, email, friends, blocks, created_at, updated_at) values ('Andrew Do', 'andrew@test.com', ARRAY [ 'andrew@test.com' ], ARRAY [ 'tom@test.com' ], '2022-08-23 00:00:00', '2022-09-23 00:00:00');
insert into public.user (name, email, created_at, updated_at) values ('Peter Do', 'peter@test.com', '2022-11-21 00:00:00', '2022-12-21 00:00:00');
insert into public.user (name, email, created_at, updated_at) values ('Donald Tran', 'donald@test.com', '2021-07-21 00:00:00', '2021-07-21 00:00:00');

