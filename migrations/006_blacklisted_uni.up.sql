create table "studentHelper_blacklisteduniversity"
(
	id serial not null
		constraint "studentHelper_blacklisteduniversity_pkey"
			primary key,
	university_id_id integer not null
		constraint "studentHelper_blackl_university_id_id_595b19c5_fk_studentHe"
			references "studentHelper_university"
				deferrable initially deferred,
	user_id_id integer not null
		constraint "studentHelper_blackl_user_id_id_dfb42bf3_fk_studentHe"
			references "studentHelper_user"
				deferrable initially deferred
);

alter table "studentHelper_blacklisteduniversity" owner to postgres;

create index "studentHelper_blacklisteduniversity_university_id_id_595b19c5"
	on "studentHelper_blacklisteduniversity" (university_id_id);

create index "studentHelper_blacklisteduniversity_user_id_id_dfb42bf3"
	on "studentHelper_blacklisteduniversity" (user_id_id);

