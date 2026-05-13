package course

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	coll *mongo.Collection
}

func NewRepository(coll *mongo.Collection) *Repository {
	return &Repository{coll: coll}
}

func (r *Repository) Insert(ctx context.Context, c *Course) (*Course, error) {
	res, err := r.coll.InsertOne(ctx, c)
	if err != nil {
		return nil, err
	}
	c.ID = res.InsertedID.(bson.ObjectID)
	return c, nil
}

// FindAll รองรับ filter เผื่อใช้ ?major=<id>
func (r *Repository) FindAll(ctx context.Context, filter bson.M) ([]Course, error) {
	if filter == nil {
		filter = bson.M{}
	}
	cur, err := r.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var items []Course
	if err := cur.All(ctx, &items); err != nil {
		return nil, err
	}
	if items == nil {
		items = []Course{}
	}
	return items, nil
}

func (r *Repository) FindByID(ctx context.Context, id bson.ObjectID) (*Course, error) {
	var c Course
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *Repository) Update(ctx context.Context, id bson.ObjectID, set bson.M) error {
	_, err := r.coll.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": set})
	return err
}

func (r *Repository) Delete(ctx context.Context, id bson.ObjectID) error {
	_, err := r.coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
