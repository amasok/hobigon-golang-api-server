package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yyh-gl/hobigon-golang-api-server/app"
	model "github.com/yyh-gl/hobigon-golang-api-server/app/domain/model/blog"
	"github.com/yyh-gl/hobigon-golang-api-server/app/usecase"
)

// Blog : Blog用REST Handlerのインターフェース
type Blog interface {
	Create(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Like(w http.ResponseWriter, r *http.Request)
}

type blog struct {
	usecase usecase.Blog
}

// NewBlog : Blog用REST Handlerを取得
func NewBlog(u usecase.Blog) Blog {
	return &blog{
		usecase: u,
	}
}

// blogResponse : Blog用共通レスポンス
type blogResponse struct {
	Blog *model.Blog `json:"blog,omitempty"`
}

// Create : ブログ情報を新規作成
func (b blog) Create(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Title string `json:"title" validate:"required"`
	}

	errRes := errorResponse{}

	req := request{}
	if err := bindReqWithValidate(r, &req); err != nil {
		errInfo := fmt.Errorf("bindReqWithValidate()でエラー: %w", err)
		app.Logger.Println(errInfo)

		errRes.Error = errInfo.Error()
		DoResponse(w, errRes, http.StatusBadRequest)
		return
	}

	blog, err := b.usecase.Create(r.Context(), req.Title)
	if err != nil {
		errInfo := fmt.Errorf("BlogUseCase.Create()でエラー: %w", err)
		app.Logger.Println(errInfo)

		errRes.Error = errInfo.Error()
		DoResponse(w, errRes, http.StatusInternalServerError)
		return
	}

	resp := blogResponse{
		Blog: blog,
	}
	DoResponse(w, resp, http.StatusCreated)
}

// Show : ブログ情報を1件取得
func (b blog) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// TODO: httprouterに依存することについて考える
	ps := httprouter.ParamsFromContext(ctx)

	resp := new(blogResponse)
	errRes := new(errorResponse)

	blog, err := b.usecase.Show(ctx, ps.ByName("title"))
	if err != nil {
		errInfo := fmt.Errorf("BlogUseCase.Show()でエラー: %w", err)
		app.Logger.Println(errInfo)

		if errors.Is(err, usecase.ErrBlogNotFound) {
			DoResponse(w, resp, http.StatusNoContent)
			return
		}

		errRes.Error = errInfo.Error()
		DoResponse(w, errRes, http.StatusInternalServerError)
		return
	}

	resp.Blog = blog
	DoResponse(w, resp, http.StatusOK)
}

// Like : 指定ブログにいいねをプラス1
func (b blog) Like(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ps := httprouter.ParamsFromContext(ctx)

	resp := new(blogResponse)
	errRes := new(errorResponse)
	blog, err := b.usecase.Like(ctx, ps.ByName("title"))
	if err != nil {
		errInfo := fmt.Errorf("BlogUseCase.Like()でエラー: %w", err)
		app.Logger.Println(errInfo)

		errRes.Error = errInfo.Error()

		if errors.Is(err, usecase.ErrBlogNotFound) {
			DoResponse(w, resp, http.StatusNoContent)
			return
		}

		DoResponse(w, errRes, http.StatusInternalServerError)
		return
	}

	resp.Blog = blog
	DoResponse(w, resp, http.StatusOK)
}