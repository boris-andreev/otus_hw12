package service

import "hw12/internal/model"

type operationType int

type operationItem struct {
	item          model.Identifier
	operationType operationType
}

const (
	add operationType = iota
	update
)
