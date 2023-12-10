package models

type Database struct {
	AuthRepository      AuthRepository
	UsersRepository     UsersRepository
	FilesRepository     FilesRepository
	CompaniesRepository CompaniesRepository
	OffersRepository    OffersRepository
	ResourcesRepository ResourcesRepository
	LocationsRepository LocationsRepository
}

type AuthRepository interface {
	Login(username string, password string) (int, error)
}

type UsersRepository interface {
	GetById(id int) (User, error)
	GetAll() ([]User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
}

type FilesRepository interface {
	GetById(id int) (File, error)
	GetFileUserIdByFileId(id int) (int, error)
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

type OffersRepository interface {
	GetById(id int) (Offer, error)
	GetUserIdByOfferId(id int) (int, error)
	GetAll() ([]Offer, error)
	Create(offer Offer) (Offer, error)
	Update(offer Offer) (Offer, error)
	Delete(id int) error
}

type ResourcesRepository interface {
	GetById(id int) (Resource, error)
	GetUserIdByResourceId(id int) (int, error)
	GetAll() ([]Resource, error)
	Create(resource Resource) (Resource, error)
	Update(resource Resource) (Resource, error)
	Delete(id int) error
}

type LocationsRepository interface {
	GetAll() ([]Location, error)
}
