create table "studentHelper_educationalprogramme"
(
	id serial not null
		constraint "studentHelper_educationalprogramme_pkey"
			primary key,
	title varchar(200) not null,
	university_id_id integer not null
		constraint "studentHelper_educat_university_id_id_1098951a_fk_studentHe"
			references "studentHelper_university"
				deferrable initially deferred,
	faculty varchar(200) not null
);

alter table "studentHelper_educationalprogramme" owner to postgres;

create index "studentHelper_educationalprogramme_university_id_id_1098951a"
	on "studentHelper_educationalprogramme" (university_id_id);
