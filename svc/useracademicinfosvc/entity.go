package useracademicinfosvc

type JsonEntity struct {
	AvgACTComposite string `json:"avg_act_composite"`
	AvgACTEnglish   string `json:"avg_act_english"`
	AvgACTMath      string `json:"avg_act_Math"`
	AvgSATRW        string `json:"avg_sat_rw"`
	AvgSATMath      string `json:"avg_sat_math"`
	UserID          int    `json:"user_id_id,fk:studentHelper_user_id"`
}
