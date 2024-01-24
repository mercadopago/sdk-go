package credential

// Credential is a custom type to save and propagate
// the access token during calls.
type Credential string

// New returns a new Credential.
func New(at string) (*Credential, error) {
	cdt := Credential(at)
	return &cdt, nil
}
