package course

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Create รับ DTO ที่มี MajorIDs เป็น []string · service ไม่แปลงเอง
// (handler แปลงให้แล้ว ส่งมาเป็น []bson.ObjectID ในรูปของ Course struct)
func (s *Service) Create(ctx context.Context, c *Course) (*Course, error) {
	return s.repo.Insert(ctx, c)
}

// List มี filter เผื่อใช้ ?major=<id>
func (s *Service) List(ctx context.Context, filterMajorID *bson.ObjectID) ([]Course, error) {
	filter := bson.M{}
	if filterMajorID != nil {
		filter["majorIds"] = *filterMajorID
	}
	return s.repo.FindAll(ctx, filter)
}

func (s *Service) Get(ctx context.Context, id bson.ObjectID) (*Course, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id bson.ObjectID, set bson.M) error {
	if len(set) == 0 {
		return nil
	}
	return s.repo.Update(ctx, id, set)
}

func (s *Service) Delete(ctx context.Context, id bson.ObjectID) error {
	return s.repo.Delete(ctx, id)
}
