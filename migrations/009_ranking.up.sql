create table "studentHelper_ranking"
(
	id serial not null
		constraint "studentHelper_ranking_pkey"
			primary key,
	university_id_id integer not null
		constraint "studentHelper_ranking_university_id_id_26e967ca_uniq"
			unique
		constraint "studentHelper_rankin_university_id_id_26e967ca_fk_studentHe"
			references "studentHelper_university"
				deferrable initially deferred,
	rank integer
);

alter table "studentHelper_ranking" owner to postgres;

