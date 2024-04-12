package structures

import (
	"container/list"
)

var (
	List = list.New()
)

type Meta struct {
	BaseForm       string
	PastSimple     string
	PastParticiple string
}
