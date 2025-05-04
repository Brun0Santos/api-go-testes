package request

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"  binding:"required,min=4,max=41"`
	Name  string `json:"name" binding:"required,min=4,max=20"`
	Age   int8   `json:"age"`
}
