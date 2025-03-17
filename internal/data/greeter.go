package data

import (
	"context"

	"github.com/go-kratos/kratos-layout/internal/biz"
	"github.com/go-kratos/kratos-layout/model"
	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	greeter := &model.Greeter{
		Hello: g.Hello,
	}
	result := r.data.db.Create(greeter)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Greeter{
		Hello: greeter.Hello,
	}, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	greeter := &model.Greeter{
		Hello: g.Hello,
	}
	result := r.data.db.Save(greeter)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

func (r *greeterRepo) FindByID(ctx context.Context, id int64) (*biz.Greeter, error) {
	greeter := &model.Greeter{}
	result := r.data.db.First(greeter, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Greeter{
		Hello: greeter.Hello,
	}, nil
}

func (r *greeterRepo) ListByHello(ctx context.Context, hello string) ([]*biz.Greeter, error) {
	var greeterList []model.Greeter
	result := r.data.db.Where("hello LIKE ?", "%"+hello+"%").Find(&greeterList)
	if result.Error != nil {
		return nil, result.Error
	}

	var bizGreeterList []*biz.Greeter
	for _, g := range greeterList {
		bizGreeterList = append(bizGreeterList, &biz.Greeter{
			Hello: g.Hello,
		})
	}
	return bizGreeterList, nil
}

func (r *greeterRepo) ListAll(ctx context.Context) ([]*biz.Greeter, error) {
	var greeterList []model.Greeter
	result := r.data.db.Find(&greeterList)
	if result.Error != nil {
		return nil, result.Error
	}

	var bizGreeterList []*biz.Greeter
	for _, g := range greeterList {
		bizGreeterList = append(bizGreeterList, &biz.Greeter{
			Hello: g.Hello,
		})
	}
	return bizGreeterList, nil
}
