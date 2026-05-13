package course

import "go.mongodb.org/mongo-driver/v2/bson"

// Course = entity ที่ map กับ collection "courses"
// 3 challenge ใหม่จากที่เคยทำใน Major:
//   1. Credits = int (ใช้ validate:"min=1,max=10")
//   2. MajorIDs = []bson.ObjectID (รับเข้ามาเป็น string[] · ต้องแปลง)
//   3. Instructor = nested struct (ต้องมี bson tag ที่ทั้ง outer + inner)
type Course struct {
	ID         bson.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name       string          `bson:"name"          json:"name"`
	Code       string          `bson:"code"          json:"code"`
	Credits    int             `bson:"credits"       json:"credits"`
	MajorIDs   []bson.ObjectID `bson:"majorIds"      json:"majorIds"`
	Instructor Instructor      `bson:"instructor"    json:"instructor"`
}

type Instructor struct {
	Name  string `bson:"name"  json:"name"`
	Email string `bson:"email" json:"email"`
}
