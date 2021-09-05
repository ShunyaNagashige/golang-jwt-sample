package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ShunyaNagashige/golang-jwt-sample/auth"
	"github.com/ShunyaNagashige/golang-jwt-sample/domain/model"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/xerrors"
)

// userにおけるHandlerのインタフェース
type UserHandler interface {
	Create(http.ResponseWriter, *http.Request, httprouter.Params)
	Auth(http.ResponseWriter, *http.Request, httprouter.Params)
}

type userHandler struct {
}

// Userデータに関するHandlerを生成
func NewUserHandler() UserHandler {
	return &userHandler{}
}

// model.Userにvalidateタグはつけず，
// それぞれでstructを用意してvalidateタグをつける
// 各メソッド+URIの組み合わせ毎に，必須フィールドが異なるから．

// validatorやencoding/jsonは
// 他パッケージ空間にあるので，
// 自分のソースコードのモデルになんやかんやするなら
// exportする必要がある．
// したがって，
// validatorは，exportされているstructureの，
// exportされているメンバに対して適用可能．
// encoding/jsonも同様

type UserCreateReq struct {
	UserName string `validate:"required" json:"userName"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserCreateRes struct {
	Token string `json:"token"`
}

type AuthRes struct {
	UserId   uint
	UserName string
	Email    string
}

var validate *validator.Validate

func (uh *userHandler) Auth(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			apiError(
				w,
				xerrors.Errorf("Failed to r.Body.Close() : %w", err),
				http.StatusInternalServerError,
				nil,
			)
		}
		return
	}()

	authHeader := r.Header.Get("Authorization")

	// authHeaderは「Bearer <JWT>」の形式になっているので，
	// JWTのみを取り出す
	token := strings.Split(authHeader, " ")[1]
	u, err := auth.ValidateToken(token)
	if err != nil {
		apiError(
			w,
			xerrors.Errorf("Error in JwtMiddleWare : %w", err),
			http.StatusBadRequest,
			nil,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	ares := AuthRes{UserId: u.UserId, UserName: u.UserName, Email: u.Email}

	if err := json.NewEncoder(w).Encode(&ares); err != nil {
		apiError(
			w,
			xerrors.Errorf("Failed to write the JSON encodig to ResponseWriter : %w", err),
			http.StatusInternalServerError,
			nil,
		)
		return
	}
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			apiError(
				w,
				xerrors.Errorf("Failed to r.Body.Close() : %w", err),
				http.StatusInternalServerError,
				nil,
			)
		}
		return
	}()

	// リクエストボディの取得
	var ucreq UserCreateReq
	if err := json.NewDecoder(r.Body).Decode(&ucreq); err != nil {
		apiError(
			w,
			xerrors.Errorf("Failed to decode a JSON object : %w", err),
			http.StatusInternalServerError,
			nil,
		)
		return
	}

	// バリデーション
	validate = validator.New()
	if err := validate.Struct(&ucreq); err != nil {
		invalidFields := make([]string, 0)

		for _, err := range err.(validator.ValidationErrors) {
			invalidFields = append(invalidFields, err.Field())
		}

		apiError(
			w,
			xerrors.Errorf("Failed to validate a UserCreateRequest : %w", err),
			http.StatusBadRequest,
			invalidFields,
		)
		return
	}

	uCreated := &model.User{UserId: 1, UserName: "shige", Email: "hohogege@example.com", Password: "aaaa"}
	// uCreated, err := uh.userUseCase.Create(ucreq.UserName, ucreq.Email, ucreq.Password)
	// if err != nil {
	// 	dbError(w, err)
	// 	return
	// }

	token, err := auth.CreateToken(uCreated.UserId, uCreated.UserName, uCreated.Email)
	if err != nil {
		apiError(
			w,
			xerrors.Errorf("Failed to validate a UserCreateRequest : %w", err),
			http.StatusBadRequest,
			nil,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	ucres := UserCreateRes{Token: token}

	if err := json.NewEncoder(w).Encode(&ucres); err != nil {
		apiError(
			w,
			xerrors.Errorf("Failed to write the JSON encodig to ResponseWriter : %w", err),
			http.StatusInternalServerError,
			nil,
		)
		return
	}
}

// func (uh *userHandler) Create(c *gin.Context) {
// var u model.User
// if err := c.BindJSON(&u); err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	return
// }

// uCreated, err := uh.userUseCase.Create(u.UserName, u.Email, u.Password)
// if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	return
// }

// token, err := auth.CreateToken(uCreated.UserId, uCreated.UserName, uCreated.Email)
// if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	return
// }

// c.JSON(http.StatusCreated, gin.H{
// 	"token": token,
// })
// }

// // UserIndex : GET /users -> 検索結果を返す
// func (uh userHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
// 	// GETパラメータ
// 	userName := r.FormValue("userName")

// 	user, err := uh.userUseCase.Search(userName)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	// クライアントにレスポンスを返却
// 	if err = json.NewEncoder(w).Encode(user); err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// }
