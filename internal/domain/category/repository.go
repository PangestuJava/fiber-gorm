package category

type Repository interface {
	FindAll() ([]Category, error)
	FindByID(id uint) (*Category, error)
	Create(category *Category) error
	Update(category *Category) error
	Delete(id uint) error
}
