package user_models

type User struct {
	Id                string   `json:"id"`
	UserCode          string   `json:"user_code"`
	PersonnelNumber   string   `json:"personnel_number"`
	LastName          string   `json:"last_name"`
	FirstName         string   `json:"first_name"`
	MiddleName        string   `json:"middle_name"`
	TLastName         string   `json:"t_last_name"`
	TFirstName        string   `json:"t_first_name"`
	TMiddleName       string   `json:"t_middle_name"`
	CompanyCode       string   `json:"company_code"`
	CompanyName       string   `json:"company_name"`
	DepartmentId      string   `json:"department_id"`
	DepartmentName    string   `json:"department_name"`
	PositionCode      string   `json:"position_code"`
	PositionName      string   `json:"position_name"`
	LoginAd           string   `json:"login_ad"`
	BitrixUserId      int      `json:"bitrix_user_id"`
	ChiefidHuman      int      `json:"chiefid_human"`
	HumanId           int      `json:"human_id"`
	Photo             string   `json:"photo"`
	BitrixRights      []string `json:"bitrix_rights"`
	IsMyTeam          bool     `json:"is_my_team"`
	BitrixArea        Area     `json:"bitrix_area"`
	ChiefId           string   `json:"chief_id"`
	City              string   `json:"city"`
	VerificationLevel int      `json:"verification_level"`
	CompanyId         string   `json:"company_id"`
	ChiefHuman        string   `json:"chief_human"`
	Email             string   `json:"email"`
	PhoneInternal     string   `json:"phone_internal"`
	PhoneMobile       string   `json:"phone_mobile"`
	DayBirth          int      `json:"day_birth"`
	MonthBirth        int      `json:"month_birth"`
	Birthday          int      `json:"birthday"`
}

type Area struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
