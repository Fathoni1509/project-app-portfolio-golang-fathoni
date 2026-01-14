--
-- PostgreSQL database dump
--

\restrict R8ZvxOyXD8G640lXU3P0kEnjuvvr6bggBPx6CvNNFzvMzYoCt68ajD1dWi99KWr

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: activity; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.activity (
    activity_id integer NOT NULL,
    name character varying(50) NOT NULL,
    year smallint NOT NULL,
    category_id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.activity OWNER TO postgres;

--
-- Name: activity_activity_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.activity_activity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.activity_activity_id_seq OWNER TO postgres;

--
-- Name: activity_activity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.activity_activity_id_seq OWNED BY public.activity.activity_id;


--
-- Name: activity_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.activity_category (
    activity_category_id integer NOT NULL,
    name character varying(50) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.activity_category OWNER TO postgres;

--
-- Name: activity_category_activity_category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.activity_category_activity_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.activity_category_activity_category_id_seq OWNER TO postgres;

--
-- Name: activity_category_activity_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.activity_category_activity_category_id_seq OWNED BY public.activity_category.activity_category_id;


--
-- Name: contact; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contact (
    contact_id integer NOT NULL,
    name character varying(50) NOT NULL,
    type character varying(50),
    link text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.contact OWNER TO postgres;

--
-- Name: contact_contact_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contact_contact_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.contact_contact_id_seq OWNER TO postgres;

--
-- Name: contact_contact_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contact_contact_id_seq OWNED BY public.contact.contact_id;


--
-- Name: personal; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.personal (
    personal_id integer NOT NULL,
    name character varying(50) NOT NULL,
    age integer NOT NULL,
    description character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.personal OWNER TO postgres;

--
-- Name: personal_personal_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.personal_personal_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.personal_personal_id_seq OWNER TO postgres;

--
-- Name: personal_personal_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.personal_personal_id_seq OWNED BY public.personal.personal_id;


--
-- Name: project; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.project (
    project_id integer NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(512),
    year smallint,
    link text,
    image_data character varying(256),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.project OWNER TO postgres;

--
-- Name: project_project_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.project_project_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.project_project_id_seq OWNER TO postgres;

--
-- Name: project_project_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.project_project_id_seq OWNED BY public.project.project_id;


--
-- Name: work; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.work (
    work_id integer NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(512),
    year smallint,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.work OWNER TO postgres;

--
-- Name: work_work_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.work_work_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.work_work_id_seq OWNER TO postgres;

--
-- Name: work_work_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.work_work_id_seq OWNED BY public.work.work_id;


--
-- Name: activity activity_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity ALTER COLUMN activity_id SET DEFAULT nextval('public.activity_activity_id_seq'::regclass);


--
-- Name: activity_category activity_category_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity_category ALTER COLUMN activity_category_id SET DEFAULT nextval('public.activity_category_activity_category_id_seq'::regclass);


--
-- Name: contact contact_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contact ALTER COLUMN contact_id SET DEFAULT nextval('public.contact_contact_id_seq'::regclass);


--
-- Name: personal personal_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.personal ALTER COLUMN personal_id SET DEFAULT nextval('public.personal_personal_id_seq'::regclass);


--
-- Name: project project_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.project ALTER COLUMN project_id SET DEFAULT nextval('public.project_project_id_seq'::regclass);


--
-- Name: work work_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.work ALTER COLUMN work_id SET DEFAULT nextval('public.work_work_id_seq'::regclass);


--
-- Data for Name: activity; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.activity (activity_id, name, year, category_id, created_at, updated_at, deleted_at) FROM stdin;
1	Informatics UNS	2023	1	2026-01-10 00:45:59.553909+07	2026-01-12 23:33:13.479246+07	\N
2	SMA N 3 Surakarta	2023	1	2026-01-10 00:48:28.469437+07	2026-01-12 23:33:38.097331+07	\N
3	Internship Entrepreneurship Himaster UNS	2023	2	2026-01-11 11:26:50.873158+07	2026-01-12 23:34:13.676373+07	\N
4	Administrator SMA N 3 Surakarta	2022	2	2026-01-12 17:35:46.034874+07	2026-01-12 23:38:49.016458+07	\N
5	OSIS SMA N 3 Surakarta	2022	2	2026-01-12 23:39:21.157599+07	2026-01-12 23:39:21.157599+07	\N
\.


--
-- Data for Name: activity_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.activity_category (activity_category_id, name, created_at, updated_at, deleted_at) FROM stdin;
1	education	2026-01-09 23:21:51.769403+07	2026-01-09 23:21:51.769403+07	\N
2	organization	2026-01-09 23:22:02.063714+07	2026-01-09 23:22:02.063714+07	\N
\.


--
-- Data for Name: contact; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.contact (contact_id, name, type, link, created_at, updated_at, deleted_at) FROM stdin;
2	Fathoni Nur Habibi	linkedin	https://www.linkedin.com/in/fathoni-nur-habibi/	2026-01-12 23:57:28.390725+07	2026-01-12 23:57:28.390725+07	\N
3	fathoni1509	Github	https://github.com/Fathoni1509	2026-01-12 23:57:53.502685+07	2026-01-12 23:57:53.502685+07	\N
1	fn._habibi	instagram	https://www.instagram.com/fn._habibi	2026-01-11 15:13:08.741424+07	2026-01-13 00:00:19.301182+07	\N
\.


--
-- Data for Name: personal; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.personal (personal_id, name, age, description, created_at, updated_at, deleted_at) FROM stdin;
1	fathoni nur habibi	21	I am a fifth-semester computer science student. I am interested in web development, UI/UX, and cybersecurity. I aspire to become a full-stack developer.	2026-01-08 17:08:42.445138+07	2026-01-12 06:15:47.687278+07	\N
\.


--
-- Data for Name: project; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.project (project_id, name, description, year, link, image_data, created_at, updated_at, deleted_at) FROM stdin;
1	latar sehat	Latar Sehat is an application that guides users through the process of hydroponic farming from start to harvest.	2025	https://bit.ly/PrototypeLatarSehat	/public/uploads/project-Screenshot 2025-11-01 161531.png	2026-01-11 14:10:16.040518+07	2026-01-12 23:50:19.467155+07	\N
2	JoKer App	The JOKER (Job Seeker) application is an application that makes it easier for users to find jobs.	2025	https://bit.ly/PrototypeMobileAppJoker	/public/uploads/project-Screenshot 2025-11-01 161735.png	2026-01-11 14:35:58.431107+07	2026-01-12 23:52:01.571217+07	\N
3	Food Recipe	A website that displays a list of recipes for food and drinks. Develope using HTML, Tailwind CSS, and JavaScript	2025	https://food-recipes-ideas.vercel.app/	/public/uploads/project-Screenshot 2025-09-21 125502.png	2026-01-12 23:55:58.71972+07	2026-01-12 23:55:58.71972+07	\N
\.


--
-- Data for Name: work; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.work (work_id, name, description, year, created_at, updated_at, deleted_at) FROM stdin;
1	frontend developer enzo group	Creating a modern, minimalist, and responsive web display using Laravel Blade	2024	2026-01-11 11:51:47.198399+07	2026-01-12 23:43:19.082825+07	\N
2	wordpress developer unisri	Creating a modern, minimalist, and responsive web design using WordPress 	2024	2026-01-12 09:09:52.861659+07	2026-01-12 23:44:22.057043+07	\N
3	Computer Systems Organization Lab Assistant	Prepare materials according to the semester learning plan (RPS). Carry out learning activities and actively communicate with students.	2025	2026-01-12 23:46:07.313721+07	2026-01-12 23:46:07.313721+07	\N
4	Digital Systems Lab Assistant	Prepare materials according to the semester learning plan (RPS). Carry out learning activities and actively communicate with students.	2024	2026-01-12 23:46:44.532986+07	2026-01-12 23:46:44.532986+07	\N
\.


--
-- Name: activity_activity_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.activity_activity_id_seq', 5, true);


--
-- Name: activity_category_activity_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.activity_category_activity_category_id_seq', 2, true);


--
-- Name: contact_contact_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.contact_contact_id_seq', 3, true);


--
-- Name: personal_personal_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.personal_personal_id_seq', 1, true);


--
-- Name: project_project_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.project_project_id_seq', 3, true);


--
-- Name: work_work_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.work_work_id_seq', 4, true);


--
-- Name: activity_category activity_category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity_category
    ADD CONSTRAINT activity_category_pkey PRIMARY KEY (activity_category_id);


--
-- Name: activity activity_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity
    ADD CONSTRAINT activity_pkey PRIMARY KEY (activity_id);


--
-- Name: contact contact_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contact
    ADD CONSTRAINT contact_pkey PRIMARY KEY (contact_id);


--
-- Name: personal personal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.personal
    ADD CONSTRAINT personal_pkey PRIMARY KEY (personal_id);


--
-- Name: project project_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.project
    ADD CONSTRAINT project_pkey PRIMARY KEY (project_id);


--
-- Name: work work_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.work
    ADD CONSTRAINT work_pkey PRIMARY KEY (work_id);


--
-- Name: activity fk_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity
    ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES public.activity_category(activity_category_id);


--
-- PostgreSQL database dump complete
--

\unrestrict R8ZvxOyXD8G640lXU3P0kEnjuvvr6bggBPx6CvNNFzvMzYoCt68ajD1dWi99KWr

