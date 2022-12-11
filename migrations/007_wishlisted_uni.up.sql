create table "studentHelper_wishlisteduniversity"
(
	id serial not null
		constraint "studentHelper_wishlisteduniversity_pkey"
			primary key,
	priority integer not null,
	university_id_id integer not null
		constraint "studentHelper_wishli_university_id_id_7dbefb93_fk_studentHe"
			references "studentHelper_university"
				deferrable initially deferred,
	user_id_id integer not null
		constraint "studentHelper_wishli_user_id_id_4f9c4c0e_fk_studentHe"
			references "studentHelper_user"
				deferrable initially deferred
);

alter table "studentHelper_wishlisteduniversity" owner to postgres;

create index "studentHelper_wishlisteduniversity_university_id_id_7dbefb93"
	on "studentHelper_wishlisteduniversity" (university_id_id);

create index "studentHelper_wishlisteduniversity_user_id_id_4f9c4c0e"
	on "studentHelper_wishlisteduniversity" (user_id_id);

