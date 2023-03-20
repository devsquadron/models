package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name,omitempty"`
	Developers []string           `bson:"developers" json:"developers,omitempty"`
	Tasks      []Task             `bson:"tasks" json:"tasks,omitempty"`
	TaskCount  uint64             `bson:"task_count" json:"task_count,omitempty"`
	Requests   []string           `json:"requests" bson:"requests"`
}

type ApproveDeny uint

const (
	APPROVE ApproveDeny = 0
	DENY    ApproveDeny = 1
)
