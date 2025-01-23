package pdf

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrFailedToSeekForwardInPDFFile = errors.New("failed to see forward in PDF file")
	ErrInvalid                      = &struct {
		Classification string
		FileHeader     error
		MissingEOF     error
	}{
		Classification: "Invalid PDF file",
		FileHeader:     errors.New("invalid header"),
		MissingEOF:     errors.New("missing %%%%EOF"),
	}
	ErrMalformedPDF = &struct {
		Classification                     string
		SeeOtherError                      error
		MissingFinalStartXRef              error
		MissingFirstStartXRef              error
		NoIntAfterStartXRef                error
		CrossRefTableNotFound              error
		NoTrailerDictAfterXRefTable        error
		XRefPrevNotAnInteger               error
		TrailerMissingSizeEntry            error
		XRefPrevDoesNotPointToXRef         error
		XRefPrevStreamDoesNotHaveTypeXRef  error
		XRefStreamMissingSize              error
		XRefPrevStreamNotFound             error
		StreamHeaderHasNoPrevElement       error
		XRefPrevStreamNotAStream           error
		XRefPrevStreamLargerThanLastStream error
		ReadingXrefPrevStream              error
	}{
		Classification:                     "malformed PDF file",
		SeeOtherError:                      errors.New("see other error"),
		MissingFinalStartXRef:              errors.New("missing final startxref"),
		MissingFirstStartXRef:              errors.New("missing first startxref"),
		NoIntAfterStartXRef:                errors.New("startxref not followed by integer"),
		CrossRefTableNotFound:              errors.New("cross-reference table not found"),
		NoTrailerDictAfterXRefTable:        errors.New("xref table not followed by trailer dictionary"),
		TrailerMissingSizeEntry:            errors.New("trailer missing size entry"),
		XRefPrevNotAnInteger:               errors.New("xref Prev is not an integer"),
		XRefPrevDoesNotPointToXRef:         errors.New("xref Prev does not point to xref"),
		XRefPrevStreamDoesNotHaveTypeXRef:  errors.New("xref Prev stream does not have type XRef"),
		XRefStreamMissingSize:              errors.New("xref stream missing size"),
		XRefPrevStreamNotFound:             errors.New("xref Prev stream not found"),
		StreamHeaderHasNoPrevElement:       errors.New("stream header has no 'Prev' element"),
		XRefPrevStreamNotAStream:           errors.New("xref prev stream not a stream"),
		XRefPrevStreamLargerThanLastStream: errors.New("xref prev stream larger than last stream"),
		ReadingXrefPrevStream:              errors.New("reading prev stream"),
	}
)

func init() {
	var class string

	err := errors.New("")
	for _, errs := range []any{ErrInvalid, ErrMalformedPDF} {
		v := reflect.ValueOf(errs).Elem()
		for i := range v.NumField() {
			f := v.Field(i)
			if f.Kind() == reflect.String {
				class = f.String()
				continue
			}
			msg := fmt.Sprintf("%s: %s", class, f.Interface().(error).Error())
			err = errors.New(msg)
			f.Set(reflect.ValueOf(err))
		}
	}
}

type errorArg struct {
	Key   string
	Value any
}

func (e errorArg) Error() string {
	return fmt.Sprintf("%s=%v", e.Key, e.Value)
}

func ErrorArg(key string, value any) error {
	return &errorArg{Key: key, Value: value}
}
