package service

type Task interface {
	AnswerTask1(data []int) (int, error)
	AnswerTask2(input1, input2 string) (bool, error)
	AnswerTask3(data string, key int) (string, error)
}

func NewService() *Service { return &Service{} }

type Service struct{}
