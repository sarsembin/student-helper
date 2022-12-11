create table "studentHelper_university"
(
    id                                      serial           not null
        constraint "studentHelper_university_pkey"
            primary key,
    title                                   varchar(200)     not null,
    address                                 varchar(200),
    country                                 varchar(200),
    region                                  varchar(200),
    scholarships                            varchar(20),
    "maleToFemale"                          varchar(20),
    "numberOfStudents"                      integer,
    "tuitFee"                               varchar(100),
    "percentageOfInternationalStudents"     double precision,
    description                             text,
    "studentsPerStaff"                      double precision,
    "acceptanceRate"                        double precision not null,
    "averageACTComposite"                   integer,
    "averageACTEnglish"                     integer,
    rank                                    varchar(20)      not null,
    "averageACTMath"                        integer,
    "averageSATEvidenceBasedReadingWriting" integer,
    "averageSATMath"                        integer
);

alter table "studentHelper_university"
    owner to postgres;

