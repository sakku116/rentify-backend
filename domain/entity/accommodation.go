package entity

type Accommodation struct {
	ID        string `json:"id" bson:"id"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
	DeletedAt int64  `json:"deleted_at" bson:"deleted_at"`
	CreatedBy string `json:"created_by" bson:"created_by"`
	UpdatedBy string `json:"updated_by" bson:"updated_by"`
	DeletedBy string `json:"deleted_by" bson:"deleted_by"`

	Type        string `json:"type" bson:"type"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	Address       string `json:"address" bson:"address"`
	Status        string `json:"status" bson:"status"`
	Price         int64  `json:"price" bson:"price"`
	Currency      string `json:"currency" bson:"currency"`
	BadroomTotal  int32  `json:"badroom_total" bson:"badroom_total"`
	BathroomTotal int32  `json:"bathroom_total" bson:"bathroom_total"`

	RoomTotal  int32 `json:"room_total" bson:"room_total"`
	RoomBooked int32 `json:"room_booked" bson:"room_booked"`

	Images     []string `json:"images" bson:"images"`
	Facilities []string `json:"facilities" bson:"facilities"`
}
