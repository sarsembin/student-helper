create table "studentHelper_user"
(
	id serial not null
		constraint "studentHelper_user_pkey"
			primary key,
	password varchar(128) not null,
	last_login timestamp with time zone,
	is_superuser boolean not null,
	first_name varchar(150) not null,
	last_name varchar(150) not null,
	is_staff boolean not null,
	is_active boolean not null,
	date_joined timestamp with time zone not null,
	username varchar(50),
	email varchar(254) not null
		constraint "studentHelper_user_email_key"
			unique
);

alter table "studentHelper_user" owner to postgres;

create index "studentHelper_user_email_9b05fa58_like"
	on "studentHelper_user" (email varchar_pattern_ops);

