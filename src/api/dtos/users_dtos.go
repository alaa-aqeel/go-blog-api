package dtos

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (r *CreateUserRequest) ToMap() map[string]any {
	return map[string]any{
		"name":     r.Name,
		"username": r.Username,
		"password": r.Password,
	}
}
