package core

import (
	"errors"
)

var (
	ErrNotFoundNode   = errors.New("node does not exist")
	ErrPopEmpty       = errors.New("empty structure. You cannot pop from it")
	ErrDuplicatedNode = errors.New("node already exists")
)
