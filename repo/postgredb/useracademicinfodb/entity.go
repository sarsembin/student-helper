package useracademicinfodb

//nolint:unused //def
type Entity struct {
	tableName       struct{} `pg:"studentHelper_universitescore"`
	ID              int      `pg:"id,pk"`
	AvgACTComposite string   `pg:"averageACTComposite"`
	AvgACTEnglish   string   `pg:"averageACTEnglish"`
	AvgACTMath      string   `pg:"averageACTMath"`
	AvgSATRW        string   `pg:"averageSATEvidenceBasedReadingWriting"`
	AvgSATMath      string   `pg:"averageSATMath"`
	UserID          int      `pg:"user_id_id,fk:studentHelper_user_id"`
}
