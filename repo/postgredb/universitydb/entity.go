package universitydb

type University struct {
	tableName                struct{} `pg:"studentHelper_university"`
	ID                       int      `pg:"id,pk"`
	Title                    string   `pg:"title"`
	Address                  string   `pg:"address"`
	Country                  string   `pg:"country"`
	Region                   string   `pg:"region"`
	Scholarships             string   `pg:"scholarships"`
	MaleToFemale             string   `pg:"maleToFemale"`
	NumberOfStudents         string   `pg:"numberOfStudents"`
	TuitionFee               string   `pg:"tuitFee"`
	PrcInternationalStudents float64  `pg:"percentageOfInternationalStudents"`
	Description              string   `pg:"description"`
	StudentsPerStaff         float64  `pg:"studentsPerStaff"`
	AcceptanceRate           float64  `pg:"acceptanceRate"`
	AvgACTComposite          int      `pg:"averageACTComposite"`
	AvgACTEnglish            int      `pg:"averageACTEnglish"`
	Rank                     string   `pg:"rank"`
	AvgACTMath               int      `pg:"averageACTMath"`
	AvgSATReadingWriting     int      `pg:"averageSATEvidenceBasedReadingWriting"`
	AvgSATMath               int      `pg:"averageSATMath"`
}
