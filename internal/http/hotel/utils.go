package hotel

import (
	"errors"
	"fmt"
	"net/http"

	app_err "github.com/marshevms/two_gis/internal/errors"
	"github.com/marshevms/two_gis/internal/logger"
	usecase_err "github.com/marshevms/two_gis/internal/usecase/errors"
)

func logAndReturnErr(cause string, err error, w http.ResponseWriter) {
	code := toHTTPCode(err)
	if code == http.StatusOK {
		return
	}

	logAndReturnCode(fmt.Sprintf("%s: %s", cause, err), code, w)
}

func logAndReturnCode(cause string, code int, w http.ResponseWriter) {
	http.Error(w, http.StatusText(code), code)
	logger.Errorf("%s: %s", cause, http.StatusText(code))
}

func toHTTPCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	var usecaseErr usecase_err.Code
	var appErr app_err.Code

	if errors.As(err, &usecaseErr) {
		return usecaseCodeToHTTPCode(usecaseErr)
	}

	if errors.As(err, &appErr) {
		return appCodeToHTTPCode(appErr)
	}

	return http.StatusInternalServerError
}

func usecaseCodeToHTTPCode(code usecase_err.Code) int {
	switch code {
	case usecase_err.DontHaveAvailableRooms,
		usecase_err.InvalidEmail:
		return http.StatusBadRequest
	case usecase_err.OrderForThatTimeAlreadyExist:
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}

func appCodeToHTTPCode(code app_err.Code) int {
	switch code {
	case app_err.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case app_err.Canceled:
		return 499
	case app_err.NotFound:
		return http.StatusNotFound
	case app_err.AlreadyExists:
		return http.StatusConflict
	case app_err.BadRequest:
		return http.StatusBadRequest
	case app_err.PermissionDenied:
		return http.StatusForbidden
	case app_err.Unauthenticated:
		return http.StatusUnauthorized
	case app_err.Unimplemented:
		return http.StatusNotImplemented
	case app_err.Unavailable:
		return http.StatusServiceUnavailable
	}

	return http.StatusInternalServerError
}
