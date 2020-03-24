package user

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nathluu/greennit/internal/app/types"
	"github.com/nathluu/greennit/internal/pkg/log"
)

type (
	service interface {
		Register(ctx echo.Context, req *types.RegisterRequest) (*types.User, error)
		FindAll(ctx echo.Context) ([]*types.UserInfo, error)
		GenerateResetPasswordToken(ctx echo.Context, mail string) (string, error)
		ResetPassword(ctx echo.Context, r ResetPasswordRequest) error
	}

	Handler struct {
		srv service
	}
)

func NewHandler(srv service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) Register(ctx echo.Context) error {
	var req types.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		log.Errorf("register request binding failed, error: %v\n", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	user, err := h.srv.Register(ctx, &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, user)
}

func (h *Handler) FindAll(ctx echo.Context) error {
	users, err := h.srv.FindAll(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, users)
}

func (h *Handler) GenerateResetPasswordToken(ctx echo.Context) error {
	var req GenerateResetPasswordTokenRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	_, err := h.srv.GenerateResetPasswordToken(ctx, req.Email)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) ResetPassword(ctx echo.Context) error {
	var req ResetPasswordRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	err := h.srv.ResetPassword(ctx, req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusOK)
}
