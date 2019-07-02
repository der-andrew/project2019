package storage

import "errors"

// Document
var ErrDocumentAlreadyExists = errors.New("document with this id already exists")
var ErrDocumentNotFound = errors.New("document not found")
var ErrDocumentNotModified = errors.New("document not modified")

// Metadata
var ErrMetadataNotFound = errors.New("metadata not found")
