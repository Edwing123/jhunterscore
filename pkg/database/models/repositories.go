package models

type AuthRepository interface {
	Login(username string, password string) (int, error)
}

type UsersRepository interface {
	GetById(id int) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	UpdateStatus(id int, status bool) error
	UpdateRole(id int, role string) error
}

type FilesRepository interface {
	GetById(id int) (File, error)
	GetAll() ([]File, error)
	Create(file File) (File, error)
	Update(file File) (File, error)
	Delete(id int) error
}

type CompaniesRepository interface {
	GetById(id int) (Company, error)
	GetAll() ([]Company, error)
	Create(company Company) (Company, error)
	Update(company Company) (Company, error)
	Delete(id int) error
}
