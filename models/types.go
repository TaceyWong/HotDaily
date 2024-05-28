package models

type ListItem struct {
	Id        string `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Cover     string `json:"cover" `
	Author    string `json:"author"`
	Desc      string `json:"desc"`
	Hot       int    `json:"hot"`
	URL       string `json:"url" validate:"required"`
	MobileURL string `json:"mobile_url" validate:"required"`
}

type RequestParams struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Title string `json:"title" form:"title" query:"title" validate:"required"`
	// Type string  `json:"name" form:"name" query:"name" validate:"required"`
	// Description string `json:"name" form:"name" query:"name" validate:"required"`
	// ParameData map[string]interface{}  `json:"name" form:"name" query:"name" validate:"required"`
	// Total int `json:"name" form:"name" query:"name" validate:"required"`
	// Link  string `json:"name" form:"name" query:"name"`
	// UpdateTime string `json:"name" form:"name" query:"name" validate:"required"`
	// FromCache bool `json:"name" form:"name" query:"name" validate:"required"`
	// Data ListItem[] `json:"name" form:"name" query:"name" validate:"required"`

}
