package models

import (
	"errors"
	"fmt"
)

type StatusType struct {
	New        string
	Developing string
	Review     string
	Complete   string
}

var (
	Status = StatusType{
		New:        "New",
		Developing: "Developing",
		Review:     "Review",
		Complete:   "Complete",
	}
	Statuses = []string{
		Status.New,
		Status.Developing,
		Status.Review,
		Status.Complete,
	}
)

var (
	nxtSts string
)

func IsStatus(candidate string) bool {
	for _, st := range Statuses {
		if st == candidate {
			return true
		}
	}
	return false
}

func NextStatus(curStatus string) (string, error) {
	finalSts := Statuses[len(Statuses)-1]
	if curStatus == finalSts {
		return "", errors.New(fmt.Sprintf("Cannot increment status, task is already %s", finalSts))
	}
	for i, st := range Statuses {
		if curStatus == st {
			return Statuses[i+1], nil
		}
	}
	return "", errors.New(fmt.Sprintf("The given status, %s, is not in %s", curStatus, Statuses))
}
