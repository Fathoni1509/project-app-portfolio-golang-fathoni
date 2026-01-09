CREATE TABLE activity (
	activity_id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	year SMALLINT NOT NULL,
	category_id INT NOT NULL,
	CONSTRAINT fk_category
		FOREIGN KEY (category_id)
		REFERENCES activity_category(activity_category_id),
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ
)

CREATE TABLE work (
	work_id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	description VARCHAR(512),
	year SMALLINT,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ
)

CREATE TABLE project (
	project_id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	description VARCHAR(512),
	year SMALLINT,
	link TEXT,
	image_data BYTEA,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ
)

CREATE TABLE contact (
	contact_id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	type VARCHAR(50),
	link TEXT,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ
)