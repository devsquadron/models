package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Developers []string           `bson:"developers" json:"developers"`
	Tasks      []Task             `bson:"tasks" json:"tasks"`
	TaskCount  uint64             `bson:"task_count" json:"task_count"`
}
