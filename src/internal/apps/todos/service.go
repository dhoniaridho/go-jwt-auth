package todos

type Service struct{}

func (s *Service) GetAll() []Todo {
	todos := []Todo{
		{
			Id:   1,
			Name: "Do the Job",
			Done: false,
		},
		{
			Id:   2,
			Name: "Code",
			Done: false,
		},
		{
			Id:   3,
			Name: "Sleep",
			Done: false,
		},
	}
	return todos
}

func (s *Service) GetOne(id string) Todo {

	todo := Todo{
		Id:   1,
		Name: "Do the Job",
		Done: false,
	}

	return todo
}
