package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/service"
)

//Check whether the directory can be used
func dirCorrect(path string) bool {
	if path == "" {
		log.Error("Path should not be empty")
		return false
	}

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.WithField("path", path).Error("Directory does'nt exist")
		return false
	}

	if err != nil {
		log.WithField("path", path).WithError(err)
		return false
	}

	if !info.IsDir() {
		log.Error("Path is'nt a directory")
		return false
	}

	return true
}

//Results of transformToBSON
const (
	success   int = iota //All document correctly transform
	failure              //Document cannot be transform
	partially            //Only part of document transform
)

//Error wrapper for case when dir is incorrect
func logIncorrectCSV(path string, lineNum int, err string) {
	log.WithFields(log.Fields{
		"file":    path,
		"lineNum": lineNum,
		"error":   "Incorrect fields number"},
	).Error("Incorrect csv structure")
}

//Check when field "code" exist in doc
func fieldCodeExist(header []string) bool {
	codeExist := false
	for _, el := range header {
		if el == "code" {
			codeExist = true
		}
	}
	return codeExist
}

const minPathLen = 2

//Transform document to Bson prepaired for import to DB
func transformToBSON(path string, input [][]string) (*[]bson.M, int) {
	pathParts := strings.Split(path, string(os.PathSeparator))
	if len(pathParts) < minPathLen {
		log.WithFields(log.Fields{"path": path}).Error("Incorrect path size, required: /{type}/{locale}.csv")
		return nil, failure
	}
	filename := strings.Split(pathParts[len(pathParts)-1], ".")[0]

	header := input[0]
	if !fieldCodeExist(header) {
		logIncorrectCSV(path, 1, "Column type \"code\" doesn't exist")
		return nil, failure
	}

	payload := []bson.M{}

	var line *[]string
	var currentBson bson.M
	for lineNum := 1; lineNum < len(input); lineNum++ {
		line = &input[lineNum]
		if len(header) != len(*line) {
			logIncorrectCSV(path, lineNum, "Incorrect fields number")
			continue
		}

		currentBson = bson.M{}
		for fieldNum := 0; fieldNum < len(*line); fieldNum++ {
			currentBson[header[fieldNum]] = (*line)[fieldNum]
		}

		if currentBson["code"] == "" {
			logIncorrectCSV(path, lineNum, "Empty field \"code\" doesn't allowed")
			continue
		}

		currentBson["locale"] = filename
		currentBson["type"] = pathParts[len(pathParts)-2]
		payload = append(payload, currentBson)
	}

	if len(payload) != len(input)-1 {
		return &payload, partially
	}

	return &payload, success
}

const updateStatusUPDATING = "UPDATING"
const updateStatusUPDATED = "UPDATED"
const updateStatusERROR = "ERROR"

func wrapDocLog(docType string, updateStatus string) {
	log.WithFields(log.Fields{
		"context":      "DOCUMENT",
		"documentType": docType,
		"updateStatus": updateStatus,
	}).Info("Thesaurus update status")
}

func importToDB(payload *[]bson.M) {
	s := service.Instance

	for _, data := range *payload {
		wrapDocLog(data["type"].(string), updateStatusUPDATING)

		filter := storage.ExtractDocumentBsonID(data)

		_, err := s.GetDocument(filter)
		if err != nil {
			switch err {
			case storage.ErrDocumentNotFound:
				_, err := s.StoreDocument(data)
				if err != nil {
					log.WithError(err).Error("Store document error")
					wrapDocLog(data["type"].(string), updateStatusERROR)
				}

				wrapDocLog(data["type"].(string), updateStatusUPDATED)
			default:
				log.WithError(err).Error("Get document error")
				wrapDocLog(data["type"].(string), updateStatusERROR)
			}
		}

		_, err = s.StoreDocument(data)
		if err == storage.ErrDocumentAlreadyExists {
			_, err = s.UpdateDocument(filter, data)
		}

		if err != nil {
			switch err {
			case storage.ErrDocumentNotModified:
				log.WithFields((map[string]interface{})(data)).Info("Document doesnt modified")
			default:
				log.WithError(err).Error("Update into db error")
				wrapDocLog(data["type"].(string), updateStatusERROR)
			}
		}

		wrapDocLog(data["type"].(string), updateStatusUPDATED)
	}
}

func processFile(path string, info os.FileInfo, colSep rune) error {
	if filepath.Ext(path) == ".csv" {
		log.WithFields(log.Fields{"path": path}).Info("Start processing file")

		file, err := os.Open(path)
		if err != nil {
			log.WithError(err).Error("Open file error")
			return nil
		}

		defer func() {
			if err := file.Close(); err != nil {
				log.WithError(err).Error("Error while file closing")
			}
		}()

		reader := csv.NewReader(file)
		reader.Comma = colSep

		input, err := reader.ReadAll()
		if err != nil {
			log.WithError(err).Error("CSV read error")
			return nil
		}

		if payload, result := transformToBSON(path, input); result != failure {
			if result == partially {
				log.WithFields(log.Fields{"path": path}).Info("CSV partially transform to BSON")
			}

			log.WithFields(log.Fields{"path": path}).Info("CSV successfully transform to BSON")

			importToDB(payload)
		}
	}
	return nil
}

// Recursively walk at dir and import all CVS document into database.
func ImportCSV(database *mongo.Database, path string, colSep string) {
	log.Info("Start importing CSV")
	if !dirCorrect(path) {
		return
	}

	conColSep, err1 := strToRune(colSep)
	if err1 != nil {
		log.Error("Column separator length should be one")
		return
	}

	err3 := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.WithError(err).Error("Iteration path error")
			return nil
		}

		processFile(path, info, *conColSep)

		return nil
	})
	if err3 != nil {
		log.WithError(err3).Error("Iteration path error")
	}

	log.Info("Finish importing CSV")
}

func strToRune(target string) (*rune, error) {
	if len(target) != 1 {
		return nil, errors.New("string's length should equal 1")
	}

	convertedVal := rune(target[0])
	return &convertedVal, nil
}
