package mp

import (
	"github.com/mercadopago/sdk-go/pkg/internal"
)

// SetAccessToken is the SDK init point.
func SetAccessToken(accessToken string) {
	internal.SetAccessToken(accessToken)
}
