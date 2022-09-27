package internal_const

import (
	"iot-service/utils/errors"
	"net/http"
)

var (
	//ErrRecordNotFound is
	ErrRecordNotFound = func() error {
		return errors.ErrBase.New("record not found").WithProperty(errors.ErrCodeProperty, http.StatusNotFound).WithProperty(errors.ErrHttpCodeProperty, http.StatusNotFound)
	}

	// ErrBadRequest
	ErrBadRequest = func(err error) error {
		return errors.ErrBase.New(err.Error()).WithProperty(errors.ErrCodeProperty, http.StatusBadRequest).WithProperty(errors.ErrHttpCodeProperty, http.StatusBadRequest)
	}

	// ErrBadRequest
	ErrBadGateway = func(err error) error {
		return errors.ErrBase.New(err.Error()).WithProperty(errors.ErrCodeProperty, http.StatusBadGateway).WithProperty(errors.ErrHttpCodeProperty, http.StatusBadGateway)
	}
)
