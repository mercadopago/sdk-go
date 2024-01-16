package credential

// Credential is a custom type to save and propagate
// the access token during calls.
type Credential string

// New returns a new Credential.
func New(at string) (Credential, error) {
	return Credential(at), nil
}
