-- Version 1.0

-- Create Table adoptMe.users;
SET search_path TO adoptMe;

CREATE TABLE users
(
  id  text,
  name character varying(20),
  firstName character varying(20),
  address character varying(100),
  dateOfBirth character varying(20),
  phoneNumber character varying(20),
  password text,
  email  text UNIQUE ,
  verified boolean DEFAULT false,
  createdAt timestamp,
  updatedAt timestamp,
  provider text,
  preferences jsonb,
  photo text,
  verificationCode text,
  role text,
  CONSTRAINT users_pkey PRIMARY KEY (id)
);
ALTER TABLE adoptMe.users OWNER TO g_adoptMe;

/*create pets table */
CREATE TABLE pets
(
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id text,
  name character varying(20),
  type character varying(20),
  age character varying(20),
  photo text,
  createdAt timestamp,
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
ALTER TABLE adoptMe.users OWNER TO g_adoptMe;
