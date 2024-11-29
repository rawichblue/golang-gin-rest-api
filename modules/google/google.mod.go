package google

type GoogleModule struct {
	Svc *GoogleOAuthService
}

func New() *GoogleModule {
	svc := newService()
	return &GoogleModule{
		Svc: svc,
	}
}
