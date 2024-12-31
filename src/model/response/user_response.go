package request

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name" binding:"required,min=4,max=50"`
	Age   int8   `json:"age"`
}
