package advantshop

type GetCustomersResult struct {
	Pagination struct {
		Count          int `json:"count"`
		CurrentPage    int `json:"currentPage"`
		TotalCount     int `json:"totalCount"`
		TotalPageCount int `json:"totalPageCount"`
	} `json:"pagination"`
	Customers []Customer `json:"customers"`
}

type Customer struct {
	Id                string      `json:"id"`
	Email             string      `json:"email"`
	Phone             int64       `json:"phone"`
	FirstName         string      `json:"firstName"`
	LastName          string      `json:"lastName"`
	Patronymic        string      `json:"patronymic"`
	Organization      string      `json:"organization"`
	SubscribedForNews bool        `json:"subscribedForNews"`
	Birthday          interface{} `json:"birthday"`
	AdminComment      interface{} `json:"adminComment"`
	ManagerId         interface{} `json:"managerId"`
	GroupId           int         `json:"groupId"`
}

type GetBonusesResult struct {
	CardId    int    `json:"cardId"`
	Amount    int    `json:"amount"`
	Percent   int    `json:"percent"`
	GradeName string `json:"gradeName"`
	GradeId   int    `json:"gradeId"`
	IsBlocked bool   `json:"isBlocked"`
	Status    string `json:"status"`
	Errors    string `json:"errors"`
}
