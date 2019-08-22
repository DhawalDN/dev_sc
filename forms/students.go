package forms

type CreateStudentCommand struct {
	Name     string `json:"name" `
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gender   string `json:"gender" `
	IsActive bool   `json:"isActive"`
}

type UpdateStudentCommand struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password bool   `json:"password"`
	Gender   string `json:"gender"`
	IsActive bool   `json:"isActive,omitempty"`
}
