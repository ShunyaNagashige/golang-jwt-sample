package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

type errorRes struct {
	Error         string   `json:"error"`
	InvalidFields []string `json:"invalidField,omitempty"`
}

func apiError(w http.ResponseWriter, err error, code int, invalidFields []string) {
	// レスポンスのMIMEタイプを指定
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	log.Printf("%+v\n", err)

	eres := &errorRes{
		Error:         err.Error(),
		InvalidFields: invalidFields,
	}

	if err := json.NewEncoder(w).Encode(&eres); err != nil {
		log.Printf("Failed to write the JSON encodig to response : %+v", err)
		return
	}
}

func dbError(w http.ResponseWriter, err error) {
	var mysqlErr *mysql.MySQLError
	var code int
	if xerrors.As(err, &mysqlErr) {
		// 1062の場合，
		if mysqlErr.Number == 1062 {
			code = http.StatusBadRequest
		} else {
			code = http.StatusInternalServerError
		}
	}

	apiError(
		w,
		xerrors.Errorf("Failed to validate a UserCreateRequest : %w", err),
		code,
		nil,
	)
	return
}
