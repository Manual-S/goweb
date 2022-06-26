package demo

type Service struct {
	repository *Repository
}

func NewService() *Service {
	return &Service{
		repository: NewRepository(),
	}
}

func (s *Service) GetUsers() []UserModel {
	userIds := s.repository.GetUserIds()
	return s.repository.GetUserByIds(userIds[0])
}
