//go:generate mockgen --source=usecase.go --destination=usecase_mock.go --package=answersheet_usecase
package answersheet_usecase

import (
	"context"
	"picket/src/entities"
	"picket/src/repository"
)

type IRepository interface {
	repository.IBaseRepository
	SendEvent(ctx context.Context, event *entities.AnswerSheetEvent) error
	GetLatestEvent(ctx context.Context,userId int, testId int, number int) ([]entities.AnswerSheetEvent,error)
}

type usecase struct {
	repository IRepository
	testUsecase ITestUsecase
}

type ITestUsecase interface {
	GetById(ctx context.Context, id int) (*entities.Test,error)
	GetContent(ctx context.Context, testId int) (*entities.TestContent,error)
}



func New(repository IRepository,testUsecase ITestUsecase) *usecase {
	return &usecase{repository: repository,testUsecase: testUsecase}
}