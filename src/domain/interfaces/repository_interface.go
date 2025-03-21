package interfaces

type FilterRepostiroyInterface interface {
	FindBy(key string, value string) map[string]interface{}
	GetAll(limit int, offset int) map[string]interface{}
}

type RepositoryInterface interface {
	FilterRepostiroyInterface
	Create(data map[string]interface{}) error
	Update(id string, data map[string]interface{}) error
	Delete(id string) error
}
