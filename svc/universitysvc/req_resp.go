package universitysvc

type CreateRequest struct {
	Title                    string  `json:"title"`
	Address                  string  `json:"address"`
	Country                  string  `json:"country"`
	Region                   string  `json:"region"`
	Scholarships             string  `json:"scholarships"`
	MaleToFemale             string  `json:"maleToFemale"`
	NumberOfStudents         string  `json:"numberOfStudents"`
	TuitionFee               string  `json:"tuitFee"`
	PrcInternationalStudents float64 `json:"percentageOfInternationalStudents"`
	Description              string  `json:"description"`
	StudentsPerStaff         float64 `json:"studentsPerStaff"`
	AcceptanceRate           float64 `json:"acceptanceRate"`
	AvgACTComposite          int     `json:"averageACTComposite"`
	AvgACTEnglish            int     `json:"averageACTEnglish"`
	Rank                     string  `json:"rank"`
	AvgACTMath               int     `json:"averageACTMath"`
	AvgSATReadingWriting     int     `json:"averageSATEvidenceBasedReadingWriting"`
	AvgSATMath               int     `json:"averageSATMath"`
}

type CreateResponse struct {
}
