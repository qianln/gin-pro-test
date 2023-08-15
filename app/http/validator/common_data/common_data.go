package common_data

type Page struct {
	Page  float64 `form:"page" json:"page" binding:"min=1"`   // 必填，页面值>=1
	Limit float64 `form:"limit" json:"limit" binding:"min=1"` // 必填，每页条数值>=1
}

type IntId struct {
	Id float64 `form:"id" json:"id" binding:"required,min=1"` // 必填，页面值>=1
}
