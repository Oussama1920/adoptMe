-- Version 1.0

-- Create Table adoptMe.users;
SET search_path TO adoptMe;

CREATE TABLE adoptMe.users
(
  id  SERIAL,
  name character varying(20),
  firstName character varying(20),
  address character varying(100),
  dateOfBirth character varying(20),
  phoneNumber character varying(20),
  CONSTRAINT users_pkey PRIMARY KEY (id)
);
ALTER TABLE adoptMe.users OWNER TO g_adoptMe;
