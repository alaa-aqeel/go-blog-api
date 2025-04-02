package helpers

type DtoCollectionInterface interface {
	ToDTO() any
}

func DtoCollection[T DtoCollectionInterface](models []T) []any {
	dtos := make([]any, len(models))
	for i, model := range models {
		dtos[i] = model.ToDTO()
	}
	return dtos
}
