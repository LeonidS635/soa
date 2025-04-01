package user

type GateWayUserHandlers struct {
	userServiceURL string
}

func NewGateWayUserHandlers(userServiceURL string) *GateWayUserHandlers {
	return &GateWayUserHandlers{
		userServiceURL: userServiceURL,
	}
}
