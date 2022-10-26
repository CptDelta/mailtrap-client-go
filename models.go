package mailtrap

type Account struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	AccessLevels []int  `json:"access_levels"`
}

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Permissions struct {
		CanRead    bool `json:"permission.can_read"`
		CanUpdate  bool `json:"permission.can_update"`
		CanDestroy bool `json:"permission.can_destroy"`
		CanLeave   bool `json:"permission.can_leave"`
	}
}

type ProjectNew struct {
	Name string `json:"name"`
}
