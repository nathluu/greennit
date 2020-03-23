package auth

type (
	service interface {
		Auth() error
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

func (h *Handler) Auth() {
	h.srv.Auth()
}
