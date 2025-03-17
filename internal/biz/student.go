package biz

import (
    "context"


)

type StudentRepo interface {
    CreateStudent(ctx context.Context, in *student.CreateStudentRequest) (*student.CreateStudentReply, error)
    UpdateStudent(ctx context.Context, in *student.UpdateStudentRequest) (*student.UpdateStudentReply, error)
    DeleteStudent(ctx context.Context, in *student.DeleteStudentRequest) (*student.DeleteStudentReply, error)
    GetStudent(ctx context.Context, in *student.GetStudentRequest) (*student.GetStudentReply, error)
}

type StudentUsecase struct {
    repo StudentRepo
    data *ent.Client
}

func NewStudentUsecase(repo StudentRepo, data *ent.Client) *StudentUsecase {
    return &StudentUsecase{
        repo: repo,
        data: data,
    }
}

func (uc *StudentUsecase) CreateStudent(ctx context.Context, in *student.CreateStudentRequest) (*student.CreateStudentReply, error) {
    // 创建学生
    s, err := uc.repo.CreateStudent(ctx, in)
    if err != nil {
        return nil, err
    }

    // 创建班级
    _, err = uc.data.StudentClass.Create().
        SetStudentID(s.Id).
        SetClassID(in.ClassId).
}




