package models

type TagDistribution struct {
	Tag   string `json:"tag" bson:"_id"`
	Count int    `json:"count" bson:"count"`
}

type MinimalDeveloper struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}

type MinimalTeam struct {
	Name       string   `bson:"name" json:"name"`
	Developers []string `bson:"developers" json:"developers"`
	Requests   []string `json:"requests,omitempty" bson:"requests"`
	TaskCount  uint64   `bson:"task_count" json:"task_count"`
}

type RespondJoinRequestReq struct {
	RequestDev  string
	ApproveDeny ApproveDeny
}
