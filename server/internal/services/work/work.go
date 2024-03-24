package payments

type RepositoryWork interface {
	Create()
	Update()
	GetAll()
	GetAllByEmail()
	Delete()
}
