create table "studentHelper_academicinformation"
(
	id serial not null
		constraint "studentHelper_academicinformation_pkey"
			primary key,
	user_id_id integer not null
		constraint "studentHelper_academicinformation_user_id_id_8343ff34_uniq"
			unique
		constraint "studentHelper_academ_user_id_id_8343ff34_fk_studentHe"
			references "studentHelper_user"
				deferrable initially deferred,
	"averageACTComposite" integer,
	"averageACTEnglish" integer,
	"averageACTMath" integer,
	"averageSATEvidenceBasedReadingWriting" integer,
	"averageSATMath" integer
);

alter table "studentHelper_academicinformation" owner to postgres;

