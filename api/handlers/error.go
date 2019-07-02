package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/getsentry/raven-go"
	apiErrors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/models"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/restapi/operations/document"
)

var (
	UnknownPathMessage = "Unknown Path"
)

// ServeError the error handler interface implementation
func ServeError(w http.ResponseWriter, r *http.Request, err error) {
	var data []string
	var dataErrors []string

	w.Header().Set("Content-Type", "application/json")

	switch e := err.(type) {
	case *apiErrors.CompositeError:
		payload := new(document.DocumentCreateBadRequestBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = PayloadValidationErrorMessage

		payload.Data = data
		payload.Errors = parseComposite(e)

		res := document.NewDocumentCreateBadRequest().WithPayload(payload)
		res.WriteResponse(w, runtime.JSONProducer())
	case *apiErrors.MethodNotAllowedError:
		w.Header().Add("Allow", strings.Join(err.(*apiErrors.MethodNotAllowedError).Allowed, ","))

		if r == nil || r.Method != "HEAD" {
			payload := new(document.DocumentCreateMethodNotAllowedBody)
			payload.Version = &Version
			payload.Data = nil

			switch r {
			case nil:
				payload.Message = "Method not allowed"
			default:
				payload.Message = fmt.Sprintf("Method %v not allowed", r.Method)
			}

			payload.Errors = dataErrors
			payload.Data = data

			res := document.NewDocumentCreateMethodNotAllowed().WithPayload(payload)
			res.WriteResponse(w, runtime.JSONProducer())
		}
	case apiErrors.Error:
		switch code := e.Code(); code {
		case 404:
			payload := new(models.Error404Data)
			payload.Version = &Version
			payload.Data = data
			payload.Errors = nil
			payload.Message = &UnknownPathMessage

			w.WriteHeader(404)
			if err := runtime.JSONProducer().Produce(w, payload); err != nil {
				panic(err)
			}
		default:
			answer500(&w, r, err)
		}
	default:
		answer500(&w, r, err)
	}
}

func answer500(rw *http.ResponseWriter, r *http.Request, err error) {
	defer func() {
		if recoverValue := recover(); recoverValue != nil {
			str := fmt.Sprint(recoverValue)

			packet := raven.NewPacket(
				str,
				raven.NewException(errors.New(str), raven.GetOrNewStacktrace(recoverValue.(error), 2, 3, nil)),
				raven.NewHttp(r),
			)

			raven.Capture(packet, nil)
		}
	}()

	payload := new(models.Error500Data)
	payload.Version = &Version
	payload.Data = []string{}
	payload.Message = &InternalServerErrorMessage

	payload.Errors = map[string]interface{}{
		"core": err.Error(),
	}

	(*rw).WriteHeader(500)

	if err := runtime.JSONProducer().Produce(*rw, payload); err != nil {
		panic(err)
	}
}

func flattenComposite(errs *apiErrors.CompositeError) *apiErrors.CompositeError {
	var res []error
	for _, er := range errs.Errors {
		switch e := er.(type) {
		case *apiErrors.CompositeError:
			if len(e.Errors) > 0 {
				flat := flattenComposite(e)
				if len(flat.Errors) > 0 {
					res = append(res, flat.Errors...)
				}
			}
		default:
			if e != nil {
				res = append(res, e)
			}
		}
	}
	return apiErrors.CompositeValidationError(res...)
}

var (
	errorBadTypePattern  = regexp.MustCompile(`cannot unmarshal \w+ into Go struct field \.(\w+) of type (\w+)`)
	errorEOFPattern      = regexp.MustCompile(`parsing .* failed, because unexpected EOF`)
	errorBodyTypePattern = regexp.MustCompile(`parsing .* failed, because parse error: expected (\w+) .*`)
	errorRequiredPattern = regexp.MustCompile(`(\w+) in \w+ is required`)
	errorEnumPattern     = regexp.MustCompile(`(\w+) in \w+ should be one of .*`)
	errorFormatPattern   = regexp.MustCompile(`(\w+) in \w+ must be of type .*`)
)

func parseComposite(err *apiErrors.CompositeError) *models.Document400DataAO1Errors {
	res := new(models.Document400DataAO1Errors)
	res.Validation = new(models.Document400DataAO1ErrorsValidation)

	validation := make(map[string]interface{})
	var core, json string

	for _, subErr := range flattenComposite(err).Errors {
		switch e := subErr.(type) {
		case *apiErrors.ParseError:
			if m := errorBadTypePattern.FindStringSubmatch(e.Error()); m != nil {
				validation[m[1]] = m[2]
			} else if m := errorBodyTypePattern.FindStringSubmatch(e.Error()); m != nil {
				core = "JSON parse error"
			} else if errorEOFPattern.MatchString(e.Error()) {
				json = "EOF"
			} else {
				core = core + e.Error() + "\n"
			}
		case *apiErrors.Validation:
			if m := errorRequiredPattern.FindStringSubmatch(e.Error()); m != nil {
				validation[m[1]] = "required"
			} else if m := errorEnumPattern.FindStringSubmatch(e.Error()); m != nil {
				validation[m[1]] = "enum"
			} else if m := errorFormatPattern.FindStringSubmatch(e.Error()); m != nil {
				validation[m[1]] = "format"
			} else {
				core = core + e.Error() + "\n"
			}
		default:
			core = core + e.Error() + "\n"
		}
	}

	if len(validation) != 0 {
		for k, v := range validation {
			v := v.(string)

			switch k {
			case "type":
				res.Validation.Type = v
			case "locale":
				res.Validation.Locale = v
			case "code":
				res.Validation.Code = v
			case "text":
				res.Validation.Text = v
			}
		}
	}

	res.Core = core
	res.JSON = json

	return res
}

type EntitiesCountError struct {
	CollectionName string
	Filter         map[string]interface{}
	ExpectedCount  int
	ActualCount    int
}

func (e EntitiesCountError) Error() string {
	return fmt.Sprintf(
		"Expected %d entity(-ies) in %v collection with filter: %+v, found %d",
		e.ExpectedCount,
		e.CollectionName,
		e.Filter,
		e.ActualCount,
	)
}
