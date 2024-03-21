package aggregates

import (
	"go_design_pattern/src/gof/behavioral/iterator/example/entities"
	"go_design_pattern/src/gof/behavioral/iterator/example/iterators"
)

func NewResumeList(resumeList ...*entities.Resume) *ResumeList {
	return &ResumeList{
		list: resumeList,
	}
}

type ResumeList struct {
	list []*entities.Resume
}

func (r *ResumeList) CreateIterator() *iterators.ResourceIterator[*entities.Resume] {
	return iterators.NewResourceIterator(r.list)
}

type ProcessList struct {
	list []*entities.Process
}

func (r *ProcessList) CreateIterator() *iterators.ResourceIterator[*entities.Process] {
	return iterators.NewResourceIterator(r.list)
}
