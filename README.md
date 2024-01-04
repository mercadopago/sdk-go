# sdk-go
Mercado Pago's Official Go Lang

### Http client
By default SDK uses an retryable http client, but it is possible:

#### Change http client configurations
- Disable retries (using httpclient.New()).
- Change timeout.
- Change retry strategy and options (such as max retries).

#### Switch used http client
- General http client for all entity clients and requests.
- Http client by each entity client and its requests.
- Http client by each request.
- Custom headers by each request.

If you want, can create your own http client, following httpclient.Requester interface (there is a simple Do method, *http.Client struct implements it also).
