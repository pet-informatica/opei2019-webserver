CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE students (
    id         uuid PRIMARY KEY default uuid_generate_v4(),
    name       text,
    cpf        text,
    birth_date date,
    phone      text,
	school     text,
	period     text,
	modality   text,
	level      text,
	census     text
)