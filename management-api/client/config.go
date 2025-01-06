package client

type Config struct {
	// Endpoint is the Logto endpoint
	// caution:For Logto Cloud users: when you’re interacting with Logto Management API, you can not use custom domain,
	// use the default Logto endpoint https://{your_tenant_id}.logto.app
	Endpoint string

	// ManagementApiResourceURL is the URL of the resource you want to access, default is https://default.logto.app/api
	// caution:For Logto Cloud users: when you’re interacting with Logto Management API, you can not use custom domain,
	// use the default Logto endpoint https://{your_tenant_id}.logto.app/api
	// make sure your m2m app has the management api role to access this resource
	ManagementApiResourceURL string

	// AppId is the client id
	AppId string

	// AppSecret is the client secret
	AppSecret string

	// Debug is the debug mode
	Debug bool
}
