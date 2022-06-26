// Package demo 直接与数据库打交道 可以理解成dao层
package demo

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

// GetUserIds 获取用户id
func (r *Repository) GetUserIds() []int {
	return []int{1, 2}
}

// GetUserByIds 根据用户id获取具体的用户信息
func (r *Repository) GetUserByIds(id int) []UserModel {
	return []UserModel{
		{
			UserID: "123",
			Name:   "lili",
			Age:    20,
		},
	}
}
