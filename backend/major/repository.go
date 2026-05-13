package major

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Repository = layer ที่คุย MongoDB อย่างเดียว · ไม่มี business logic
type Repository struct {
	coll *mongo.Collection
}

func NewRepository(coll *mongo.Collection) *Repository {
	return &Repository{coll: coll}
}

func (r *Repository) Insert(ctx context.Context, m *Major) (*Major, error) {
	res, err := r.coll.InsertOne(ctx, m)
	if err != nil {
		return nil, err
	}
	m.ID = res.InsertedID.(bson.ObjectID)
	return m, nil
}

func (r *Repository) FindAll(ctx context.Context) ([]Major, error) {
	cur, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var items []Major
	if err := cur.All(ctx, &items); err != nil {
		return nil, err
	}
	if items == nil {
		items = []Major{}
	}
	return items, nil
}

func (r *Repository) FindByID(ctx context.Context, id bson.ObjectID) (*Major, error) {
	var m Major
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *Repository) Update(ctx context.Context, id bson.ObjectID, set bson.M) error {
	_, err := r.coll.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": set})
	return err
}

func (r *Repository) Delete(ctx context.Context, id bson.ObjectID) error {
	_, err := r.coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
