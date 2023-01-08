package queue

import (
	"github.com/tomo25/open-data-structures/utils"
)

type Queue struct {
	n, cap, i int
	buf       []utils.V
}

func New() Queue {
	return Queue{}
}

func (q Queue) Add(x utils.V) string {
	//TODO implement me
	panic("implement me")
}

func (q Queue) Remove(x utils.V) string {
	//TODO implement me
	panic("implement me")
}
