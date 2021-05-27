package Models

type Todo struct {
	ID uint64  `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

func (b *Todo) TableName() string {
	return "todo"
}