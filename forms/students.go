package forms

type CreateStudentCommand struct {
	Name     string `json:"name" `
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gender   string `json:"gender" `
	IsActive bool   `json:"isActive"`
}
