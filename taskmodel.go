package models

// TODO : maybe change size to uint16 to save space (allows 65,000 tasks)
type Task struct {
	ID        uint64   `bson:"task_id" json:"task_id"`
	Title     string   `bson:"title" json:"title"`
	Status    string   `bson:"status" json:"status"`
	Developer string   `bson:"developer" json:"developer"`
	Percent   int      `bson:"percent" json:"percent"`
	Criterion string   `bson:"criterion" json:"criterion"`
	Tags      []string `bson:"tags" json:"tags"`
}
