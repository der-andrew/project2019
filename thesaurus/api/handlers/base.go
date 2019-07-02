package handlers

import (
	"fmt"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/service"
)

// PayloadSuccessMessage used in 200 answers
var PayloadSuccessMessage = "SUCCESS"

// PayloadSuccessMessage used in 500 answers
var InternalServerErrorMessage = "Internal server error"

// PayloadValidationErrorMessage - used on composite errors
var PayloadValidationErrorMessage = "Validation error"

// MethodNotAllowedMessage - used in 405 answers
func MethodNotAllowedMessage(methodName string) string {
	return fmt.Sprintf("Method %s not allowed", methodName)
}

// NotFoundMessage used in 404 answers
var NotFoundMessage = "Entity not found"

// AccessDeniedMessage used in 401 answers
var AccessDeniedMessage = "Access denied"

// Version of service
var Version string

// Service instance
var Service *service.Service
