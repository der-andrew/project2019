package service

import "repo.nefrosovet.ru/maximus-platform/thesaurus/storage"

var Instance *Service

type Service struct {
	storage.Storage
}

func NewInstance(s storage.Storage) *Service {
	Instance = New(s)

	return Instance
}

func New(s storage.Storage) *Service {
	return &Service{
		Storage: s,
	}
}
