package management

type Servicer interface {
	Run() error
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Run() error {
	return nil
}
