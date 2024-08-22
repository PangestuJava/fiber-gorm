package category

type Service interface {
	GetAllCategories() ([]Category, error)
	GetCategoryByID(id uint) (*Category, error)
	CreateCategory(name string) error
	UpdateCategory(id uint, name string) error
	DeleteCategory(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAllCategories() ([]Category, error) {
	return s.repo.FindAll()
}

func (s *service) GetCategoryByID(id uint) (*Category, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateCategory(name string) error {
	category := Category{Name: name}
	return s.repo.Create(&category)
}

func (s *service) UpdateCategory(id uint, name string) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	category.Name = name
	return s.repo.Update(category)
}

func (s *service) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}
