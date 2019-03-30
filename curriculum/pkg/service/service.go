package service

import (
	"context"
	"book/utils/db"
	"fmt"
	"book/model"
)

type CurriculumInfo struct {
	Id     int64
	UserId int64
	Title  string
	Desc   string
}

// CurriculumService describes the service.
type CurriculumService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	//根据用户ID获取所有的课程ID
	CurriculumList(ctx context.Context, userId int64) ([]CurriculumInfo, error)
	//获取某个课程ID的信息
	Curriculum(ctx context.Context, id int64) (*CurriculumInfo, error)
}

type basicCurriculumService struct{}

func (b *basicCurriculumService) CurriculumList(ctx context.Context, userId int64) (c0 []CurriculumInfo, e1 error) {
	// TODO implement the business logic of CurriculumList
	currInfo := make([]model.Curriculum, 0)
	fmt.Println(userId)
	err := db.G_db.Where("user_id = ?",userId).Find(&currInfo)
	var curr = []CurriculumInfo{}
	for _,v := range currInfo{
		var exd = CurriculumInfo{
			Id:v.Id,
			UserId:v.UserId,
			Title:v.Title,
			Desc:v.Desc,
		}
		curr = append(curr,exd)
	}
	return curr,err
}
func (b *basicCurriculumService) Curriculum(ctx context.Context, id int64) (c0 *CurriculumInfo, e1 error) {
	currInfo := make([]model.Curriculum, 0)
	err := db.G_db.Where("id = ?",id).Find(&currInfo)
	var userInfo = &CurriculumInfo{
		Id:currInfo[0].Id,
		UserId:currInfo[0].UserId,
		Title:currInfo[0].Title,
		Desc:currInfo[0].Desc,
	}
	return userInfo, err
}

// NewBasicCurriculumService returns a naive, stateless implementation of CurriculumService.
func NewBasicCurriculumService() CurriculumService {
	return &basicCurriculumService{}
}

// New returns a CurriculumService with all of the expected middleware wired in.
func New(middleware []Middleware) CurriculumService {
	var svc CurriculumService = NewBasicCurriculumService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
