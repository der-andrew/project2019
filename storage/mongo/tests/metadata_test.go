package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"gitlab.com/project2019-02/thesaurus/models"
	"gitlab.com/project2019-02/thesaurus/storage"
)

type MetadataSuite struct {
	suite.Suite
}

func (s *MetadataSuite) SetupTest() {
	_, err := Storage.DeleteMetadata("patient.genders_storageMetadataTest")
	if err != storage.ErrMetadataNotFound {
		s.Require().NoError(err)
	}
}

func (s *MetadataSuite) TestStoreOrUpdateMetadata() {
	metadata, err := Storage.StoreOrUpdateMetadata("patient.genders_storageMetadataTest", nil)
	s.NoError(err)

	s.Equal(models.DefaultMetadataType, metadata.Type)

	_, err = Storage.GetMetadata("patient.genders_storageMetadataTest")
	s.NoError(err)
}

func (s *MetadataSuite) TestStoreMetadata() {
	metadata, err := Storage.StoreMetadata("patient.genders_storageMetadataTest", models.TypeStatic)
	s.NoError(err)

	s.Equal(models.TypeStatic, metadata.Type)

	_, err = Storage.GetMetadata("patient.genders_storageMetadataTest")
	s.NoError(err)
}

func (s *MetadataSuite) TestUpdateMetadata() {
	_, err := Storage.UpdateMetadata("patient.genders_storageMetadataTest", models.UpdMetadata{
		DateOfChange: models.PtrT(time.Now()),
	})
	s.Error(err)
	s.Equal(storage.ErrMetadataNotFound, err)

	_, err = Storage.StoreOrUpdateMetadata("patient.genders_storageMetadataTest", nil)
	s.NoError(err)

	static := models.TypeStatic
	_, err = Storage.UpdateMetadata("patient.genders_storageMetadataTest", models.UpdMetadata{
		Type: &static,
	})
	s.NoError(err)

	foundedMetadata, err := Storage.GetMetadata("patient.genders_storageMetadataTest")
	s.NoError(err)

	s.Equal(models.TypeStatic, foundedMetadata.Type)
}

func (s *MetadataSuite) TestGetMetadata() {
	_, err := Storage.GetMetadata("patient.genders_storageMetadataTest")
	s.Error(err)
	s.Equal(storage.ErrMetadataNotFound, err)

	_, err = Storage.StoreOrUpdateMetadata("patient.genders_storageMetadataTest", nil)
	s.NoError(err)

	metadata, err := Storage.GetMetadata("patient.genders_storageMetadataTest")
	s.NoError(err)

	s.Equal(models.DefaultMetadataType, metadata.Type)
}

func (s *MetadataSuite) TestDeleteMetadata() {
	_, err := Storage.StoreOrUpdateMetadata("patient.genders_storageMetadataTest", nil)
	s.NoError(err)

	_, err = Storage.DeleteMetadata("patient.genders_storageMetadataTest")
	s.NoError(err)

	_, err = Storage.GetMetadata("patient.genders_storageMetadataTest")
	s.Error(err)
	s.Equal(storage.ErrMetadataNotFound, err)
}

func TestMetadata(t *testing.T) {
	suite.Run(t, new(MetadataSuite))
}
