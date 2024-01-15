package credential

type Credential string

func New(at string) (Credential, error) {
	return Credential(at), nil
}
