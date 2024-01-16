package maxma

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

var ErrNotFound = errors.New("Maxma: not found")

// ErrorWrap - обертка для ошибок HTTP
func ErrorWrap(code int, message string) error {
	if code == http.StatusOK {
		return nil
	}

	if code == NotFoundError {
		return ErrNotFound
	}

	message = strings.TrimSpace(message)
	if message == "" {
		switch code {
		case BadRequestError:
			message = "BadRequestError"
		case UnauthorizedError:
			message = "UnauthorizedError"
		case NotFoundError:
			message = "NotFoundError"
		case InternalServerError:
			message = "InternalServerError"
		case MethodNotImplementedErr:
			message = "MethodNotImplementedErr"
		default:
			message = "NotErrorFound"
		}
	}
	return errors.New(fmt.Sprintf("%d: %s", code, message))
}
