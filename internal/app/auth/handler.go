package auth

type (
	service interface {
		Auth() error
	}

	Handler struct {
		srv service
	}
)

func (h *Handler) Auth() {
	h.srv.Auth()
}
