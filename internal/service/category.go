package service

import (
	"context"

	"github.com/paulomalandrim/fullcycle-gRPC/internal/database"
	"github.com/paulomalandrim/fullcycle-gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CateryDB database.Category
}

func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{
		CateryDB: db,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := s.CateryDB.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}
