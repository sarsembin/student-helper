create table "studentHelper_universitescore"
(
    id               serial       not null
        constraint "studentHelper_universitescore_pkey"
            primary key,
    title            varchar(100) not null,
    eval             varchar(100),
    university_id_id integer      not null
        constraint "studentHelper_univer_university_id_id_0d2e074a_fk_studentHe"
            references "studentHelper_university"
            deferrable initially deferred
);

alter table "studentHelper_universitescore"
    owner to postgres;

create index "studentHelper_universitescore_university_id_id_0d2e074a"
    on "studentHelper_universitescore" (university_id_id);

