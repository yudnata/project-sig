package points

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePoint(ctx context.Context, ownerID string, input CreatePointReq) (*GeoPoint, error) {
	point := &GeoPoint{
		TypeID:            &input.TypeID,
		Name:              input.Name,
		Latitude:          input.Latitude,
		Longitude:         input.Longitude,
		Address:           input.Address,
		OwnerID:           ownerID,
		TahunBerdiri:      &input.TahunBerdiri,
		StatusKepemilikan: input.StatusKepemilikan,
		Description:       input.Description,
		IsActive:          true,
	}

	err := s.repo.Create(ctx, point)
	if err != nil {
		return nil, err
	}

	return point, nil
}

func (s *Service) GetAllPoints(ctx context.Context) ([]GeoPoint, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) UpdatePoint(ctx context.Context, id int, req CreatePointReq) (*GeoPoint, error) {
	typeID := req.TypeID
	tahunBerdiri := req.TahunBerdiri

	point := &GeoPoint{
		TypeID:            &typeID,
		Name:              req.Name,
		Latitude:          req.Latitude,
		Longitude:         req.Longitude,
		Address:           req.Address,
		TahunBerdiri:      &tahunBerdiri,
		StatusKepemilikan: req.StatusKepemilikan,
		Description:       req.Description,
	}

	if err := s.repo.Update(ctx, id, point); err != nil {
		return nil, err
	}

	return point, nil
}

func (s *Service) DeletePoint(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
