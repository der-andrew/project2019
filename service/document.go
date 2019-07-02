package service

import (
	"errors"
	"time"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"

	"github.com/mongodb/mongo-go-driver/bson"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/models"
)

var ErrAccessDenied = errors.New("access denied")

func (s *Service) StoreDocument(in bson.M) (models.Document, error) {
	m, err := s.GetMetadata(in["type"].(string))
	if err != storage.ErrMetadataNotFound {
		if err != nil {
			return models.Document{}, err
		}

		if m.Type == models.TypeStatic {
			return models.Document{}, ErrAccessDenied
		}
	}

	doc, err := s.Storage.StoreDocument(in)
	if err != nil {
		return models.Document{}, err
	}

	_, err = s.StoreOrUpdateMetadata(*doc.Type, nil)
	if err != nil {
		return doc, err
	}

	return doc, nil
}

func (s *Service) UpdateDocument(in bson.M, inChanges bson.M) (models.Document, error) {
	m, err := s.GetMetadata(in["type"].(string))
	if err != storage.ErrMetadataNotFound {
		if err != nil {
			return models.Document{}, err
		}

		if m.Type == models.TypeStatic {
			return models.Document{}, ErrAccessDenied
		}
	}

	doc, err := s.Storage.UpdateDocument(in, inChanges)
	if err != nil {
		return models.Document{}, err
	}

	_, err = s.UpdateMetadata(in["type"].(string), models.UpdMetadata{
		DateOfChange: models.PtrT(time.Now()),
	})
	if err != nil {
		return doc, err
	}

	return doc, err
}

func (s *Service) DeleteDocuments(in bson.M) (bool, error) {
	m, err := s.GetMetadata(in["type"].(string))
	if err != storage.ErrMetadataNotFound {
		if err != nil {
			return false, err
		}

		if m.Type == models.TypeStatic {
			return false, ErrAccessDenied
		}
	}

	ok, err := s.Storage.DeleteDocuments(in)
	if err != nil {
		return false, err
	}

	count, err := s.Storage.CountDocuments(in["type"].(string))
	if err != nil {
		return ok, err
	}

	if count == 0 {
		_, err := s.Storage.DeleteMetadata(in["type"].(string))
		if err != nil {
			return ok, err
		}
	}

	return ok, nil
}

func (s *Service) DeleteDocument(in bson.M) (bool, error) {
	m, err := s.GetMetadata(in["type"].(string))
	if err != storage.ErrMetadataNotFound {
		if err != nil {
			return false, err
		}

		if m.Type == models.TypeStatic {
			return false, ErrAccessDenied
		}
	}

	ok, err := s.Storage.DeleteDocument(in)
	if err != nil {
		return false, err
	}

	count, err := s.Storage.CountDocuments(in["type"].(string))
	if err != nil {
		return ok, err
	}

	if count == 0 {
		_, err := s.Storage.DeleteMetadata(in["type"].(string))
		if err != nil {
			return ok, err
		}
	}

	return ok, nil
}
