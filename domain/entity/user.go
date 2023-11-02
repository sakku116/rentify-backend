package entity

type User struct {
	ID        string `json:"id" bson:"id"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
	DeletedAt int64  `json:"deleted_at" bson:"deleted_at"`
	CreatedBy string `json:"created_by" bson:"created_by"`
	UpdatedBy string `json:"updated_by" bson:"updated_by"`
	DeletedBy string `json:"deleted_by" bson:"deleted_by"`

	Username    string `json:"username" bson:"username"`
	Fullname    string `json:"fullname" bson:"fullname"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Address     string `json:"address" bson:"address"`
	IDNumber    string `json:"id_number" bson:"id_number"`
	BirthPlace  string `json:"birth_place" bson:"birth_place"`
	BirthDate   string `json:"birth_date" bson:"birth_date"`
	Password    string `json:"password" bson:"password"`
	IsActive    bool   `json:"is_active,omitempty" bson:"is_active"`
	Role        string `json:"role" bson:"role"`
	Image       string `json:"image" bson:"image"`
	SessionID   string `json:"session_id" bson:"session_id"`
}
