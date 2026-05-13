package major

import "go.mongodb.org/mongo-driver/v2/bson"

// Major = entity ที่ map กับ collection "majors" ใน Mongo
type Major struct {
	ID   bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"name"          json:"name"`
	Code string        `bson:"code"          json:"code"`
}
