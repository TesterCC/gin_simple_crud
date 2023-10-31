package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Command string             `bson:"command" json:"command"`
	Status  uint8              `bson:"status" json:"status"`
	Type    uint8              `bson:"type" json:"type"`

	CreatedBy string `bson:"created_by" json:"created_by"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`
	StartTime int64  `bson:"start_time" json:"start_time"`
	EndTime   int64  `bson:"end_time" json:"end_time"`

	IsDel uint8 `bson:"is_del" json:"is_del"` // 1 deleted, 0 available
}

// root_node mongodb test_vm.task
type CmdTask struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TaskID     uint32             `bson:"task_id" json:"task_id"`
	NodeID     uint32             `bson:"node_id" json:"node_id"`
	Creator    string             `bson:"creator" json:"creator"`
	CreateTime int64              `bson:"create_time" json:"create_time"`
	TaskResult []string           `bson:"task_result" json:"task_result"`
	TaskStatus uint8              `bson:"task_status" json:"task_status"`
	CType      uint8              `bson:"ctype" json:"ctype"`
	Remark     string             `bson:"remark" json:"remark"`
}
