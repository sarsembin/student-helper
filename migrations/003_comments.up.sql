create table "studentHelper_comments"
(
	id serial not null
		constraint "studentHelper_comments_pkey"
			primary key,
	content text,
	university_id_id integer not null
		constraint "studentHelper_commen_university_id_id_105e5947_fk_studentHe"
			references "studentHelper_university"
				deferrable initially deferred,
	user_id_id integer not null
		constraint "studentHelper_commen_user_id_id_e19bf407_fk_studentHe"
			references "studentHelper_user"
				deferrable initially deferred
);

alter table "studentHelper_comments" owner to postgres;

create index "studentHelper_comments_university_id_id_105e5947"
	on "studentHelper_comments" (university_id_id);

create index "studentHelper_comments_user_id_id_e19bf407"
	on "studentHelper_comments" (user_id_id);

