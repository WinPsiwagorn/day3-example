package major

// CreateMajorDTO = payload ที่ client ส่งมาตอน POST /majors
type CreateMajorDTO struct {
	Name string `json:"name" validate:"required,min=3,max=200"`
	Code string `json:"code" validate:"required,min=2,max=10"`
}

// UpdateMajorDTO ใช้ตอน PUT — รับเฉพาะ field ที่จะแก้ (optional)
type UpdateMajorDTO struct {
	Name *string `json:"name" validate:"omitempty,min=3,max=200"`
	Code *string `json:"code" validate:"omitempty,min=2,max=10"`
}
