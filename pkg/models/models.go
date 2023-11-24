package models

type Offer struct {
	Id           int      `json:"id"`
	Title        string   `json:"title"`
	Role         string   `json:"role"`
	Company      string   `json:"company"`
	Description  string   `json:"description"`
	Requirements []string `json:"requirements"`
	Contract     string   `json:"contract"`
	Location     string   `json:"location"`
	Salary       string   `json:"salary"`
	Benefits     []string `json:"benefits"`
	ContactInfo  string   `json:"contact_info"`
	PublishedAt  string   `json:"published_at"`
	LastUpdate   string   `json:"last_update,omitempty"`
	IsPublished  bool     `json:"is_published,omitempty"`
}

type CompactOffer struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Contract    string `json:"contract"`
	Salary      string `json:"salary"`
	PublishedAt string `json:"published_at"`
}

type Resource struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Summary     string `json:"summary"`
	Content     string `json:"content"`
	PublishedAt string `json:"published_at"`
	LastUpdate  string `json:"last_update,omitempty"`
	IsPublished bool   `json:"is_published,omitempty"`
}

type CompactResource struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Autor       string `json:"author"`
	PublishedAt string `json:"published_at"`
}

type User struct {
	Id        int
	Username  string
	Password  string
	FirstName string
	LastName  string
	Role      string
	IsActive  bool
}
