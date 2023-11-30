package models

type Offer struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Role        string  `json:"role"`
	Company     string  `json:"company"`
	Content     string  `json:"content"`
	Contract    string  `json:"contract"`
	Location    string  `json:"location"`
	Salary      float64 `json:"salary"`
	ContactInfo string  `json:"contact_info"`
	CreatedAt   string  `json:"created_at"`
	Author      string  `json:"author,omitempty"`
	IsPublished bool    `json:"is_published,omitempty"`
}

type CompactOffer struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Role        string  `json:"role"`
	Company     string  `json:"company"`
	Contract    string  `json:"contract"`
	Salary      float64 `json:"salary"`
	CreatedAt   string  `json:"created_at"`
	IsPublished bool    `json:"is_published,omitempty"`
}

type Resource struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Summary     string `json:"summary"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	IsPublished bool   `json:"is_published,omitempty"`
}

type CompactResource struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	Id        int
	Username  string
	Password  string
	FirstName string
	LastName  string
	Role      string
	IsActive  bool
	CreatedAt string
}

type Company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
