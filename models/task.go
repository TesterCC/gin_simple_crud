package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Command   string             `bson:"command" json:"command"`
	Status    uint8              `bson:"status" json:"status"`
	Type      uint8              `bson:"type" json:"type"`

	CreatedBy string             `bson:"created_by" json:"created_by"`
	CreatedAt int64 `bson:"created_at" json:"created_at"`
	UpdatedAt int64 `bson:"updated_at" json:"updated_at"`
	StartTime int64 `bson:"start_time" json:"start_time"`
	EndTime   int64 `bson:"end_time" json:"end_time"`

	IsDel     uint8              `bson:"is_del" json:"is_del"` // 1 deleted, 0 available
}
