package entity

type Review struct {
	ID        string `json:"id" bson:"id"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
	DeletedAt int64  `json:"deleted_at" bson:"deleted_at"`
	CreatedBy string `json:"created_by" bson:"created_by"`
	UpdatedBy string `json:"updated_by" bson:"updated_by"`
	DeletedBy string `json:"deleted_by" bson:"deleted_by"`

	AccommodationID string `json:"accommodation_id" bson:"accommodation_id"`
	Rating          int8   `json:"rating" bson:"rating"`
	Comment         string `json:"comment" bson:"comment"`
}
