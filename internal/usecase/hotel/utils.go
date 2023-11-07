package hotel

import (
	app_err "github.com/marshevms/two_gis/internal/errors"
	rep_err "github.com/marshevms/two_gis/internal/repository/errors"
)

func toUsescaseError(err error) error {
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case rep_err.Code:
		return repErrToUsecaseErr(e)
	case app_err.Code:
		return err
	default:
		return err
	}
}

func repErrToUsecaseErr(code rep_err.Code) error {
	return code
}
