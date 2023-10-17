package entity

type User struct {
	ID        string `json:"id" bson:"id" bson_patch:"id,omitempty"`
	Username  string `json:"username" bson:"username" bson_patch:"username,omitempty"`
	Password  string `json:"password" bson:"password" bson_patch:"password,omitempty"`
	IsActive  bool   `json:"is_active" bson:"is_active" bson_patch:"is_active,omitempty"`
	SessionID string `json:"session_id" bson:"session_id" bson_patch:"session_id,omitempty"`
	CreatedAt int64  `json:"created_at" bson:"created_at" bson_patch:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at" bson_patch:"updated_at,omitempty"`
	CreatedBy string `json:"created_by" bson:"created_by" bson_patch:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by" bson:"updated_by" bson_patch:"updated_by,omitempty"`
}
