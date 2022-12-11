create table "studentHelper_userprofile"
(
	id serial not null
		constraint "studentHelper_userprofile_pkey"
			primary key,
	"dateOfBirth" date,
	country varchar(50),
	city varchar(50),
	"phoneNumber" varchar(50),
	gender boolean,
	user_id integer not null
		constraint "studentHelper_userprofile_user_id_key"
			unique
		constraint "studentHelper_userpr_user_id_f54bbda9_fk_studentHe"
			references "studentHelper_user"
				deferrable initially deferred
);

alter table "studentHelper_userprofile" owner to postgres;
