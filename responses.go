package models

type TagDistribution struct {
	Tag   string `json:"tag" bson:"_id"`
	Count int    `json:"count" bson:"count"`
}
