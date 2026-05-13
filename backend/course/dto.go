package course

// CreateCourseDTO = payload ที่ client ส่งมา
// สังเกต MajorIDs รับเป็น []string เพราะ JSON ไม่มี ObjectID type
// dive,len=24 = ตรวจทุก element ต้องเป็น hex 24 ตัว
type CreateCourseDTO struct {
	Name       string        `json:"name"       validate:"required,min=3,max=200"`
	Code       string        `json:"code"       validate:"required,min=2,max=20"`
	Credits    int           `json:"credits"    validate:"required,min=1,max=10"`
	MajorIDs   []string      `json:"majorIds"   validate:"dive,len=24,hexadecimal"`
	Instructor InstructorDTO `json:"instructor" validate:"required"`
}

// nested struct ก็ใส่ validate tag ภายในได้
type InstructorDTO struct {
	Name  string `json:"name"  validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateCourseDTO struct {
	Name       *string        `json:"name"       validate:"omitempty,min=3,max=200"`
	Code       *string        `json:"code"       validate:"omitempty,min=2,max=20"`
	Credits    *int           `json:"credits"    validate:"omitempty,min=1,max=10"`
	MajorIDs   *[]string      `json:"majorIds"   validate:"omitempty,dive,len=24,hexadecimal"`
	Instructor *InstructorDTO `json:"instructor"`
}
