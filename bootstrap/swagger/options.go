package swagger

import (
	"github.com/realHoangHai/kratos/bootstrap/swagger/handler"
)

type HandlerOption func(opt *handler.Config)

// WithTitle Title of index file.
func WithTitle(title string) HandlerOption {
	return func(opt *handler.Config) {
		opt.Title = title
	}
}

// WithBasePath Base URL to docs.
func WithBasePath(path string) HandlerOption {
	return func(opt *handler.Config) {
		opt.BasePath = path
	}
}

// WithShowTopBar Show navigation top bar, hidden by default.
func WithShowTopBar(show bool) HandlerOption {
	return func(opt *handler.Config) {
		opt.ShowTopBar = show
	}
}

// WithHideCurl Hide curl code snippet
func WithHideCurl(hide bool) HandlerOption {
	return func(opt *handler.Config) {
		opt.HideCurl = hide
	}
}

// WithJsonEditor Enable visual json editor support (experimental, can fail with complex schemas).
func WithJsonEditor(enable bool) HandlerOption {
	return func(opt *handler.Config) {
		opt.JsonEditor = enable
	}
}

// WithPreAuthorizeApiKey Map of security name to key value
func WithPreAuthorizeApiKey(keys map[string]string) HandlerOption {
	return func(opt *handler.Config) {
		opt.PreAuthorizeApiKey = keys
	}
}

// WithSettingsUI contains keys and plain javascript values of SwaggerUIBundle configuration.
// Overrides default values.
// See https://swagger.io/docs/open-source-tools/swagger-ui/usage/configuration/ for available options.
func WithSettingsUI(settings map[string]string) HandlerOption {
	return func(opt *handler.Config) {
		opt.SettingsUI = settings
	}
}

func WithLocalFile(filePath string) HandlerOption {
	return func(opt *handler.Config) {
		opt.LocalOpenApiFile = filePath
	}
}

func WithMemoryData(data []byte, ext string) HandlerOption {
	return func(opt *handler.Config) {
		opt.OpenApiData = data
		opt.OpenApiDataType = ext
	}
}

// WithRemoteFile URL to openapi.json/swagger.json document specification.
func WithRemoteFile(filePath string) HandlerOption {
	return func(opt *handler.Config) {
		opt.SwaggerJSON = filePath
	}
}
