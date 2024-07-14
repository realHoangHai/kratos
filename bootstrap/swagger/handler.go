package swagger

import (
	"fmt"
	handler2 "github.com/realHoangHai/kratos/bootstrap/swagger/handler"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

// Handler handles swagger UI request.
type Handler = handler2.Handler

// New creates HTTP handler for Swagger UI.
func New(title, swaggerJSONPath string, basePath string) http.Handler {
	return newHandler(title, swaggerJSONPath, basePath)
}

// NewWithOption creates configurable handler constructor.
func NewWithOption(handlerOpts ...HandlerOption) http.Handler {
	opts := handler2.NewConfig()

	for _, o := range handlerOpts {
		o(opts)
	}

	return newHandlerWithConfig(opts)
}

// newHandlerWithConfig creates HTTP handler for Swagger UI.
func newHandlerWithConfig(config *handler2.Config) *Handler {
	return handler2.NewHandlerWithConfig(config, assetsBase, faviconBase, staticServer)
}

// NewHandler creates HTTP handler for Swagger UI.
func newHandler(title, swaggerJSONPath string, basePath string) *Handler {
	return newHandlerWithConfig(&handler2.Config{
		Title:       title,
		SwaggerJSON: swaggerJSONPath,
		BasePath:    basePath,
	})
}

type openJsonFileHandler struct {
	Content []byte
}

func (h *openJsonFileHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	_, _ = writer.Write(h.Content)
}

func (h *openJsonFileHandler) loadOpenApiFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	content, err := io.ReadAll(file)
	return content, err
}

func (h *openJsonFileHandler) LoadFile(filePath string) error {
	content, err := h.loadOpenApiFile(filePath)
	if err != nil {
		return err
	}

	h.Content = content
	return nil
}

type httpServerInterface interface {
	HandlePrefix(prefix string, h http.Handler)
	Handle(path string, h http.Handler)
	HandleFunc(path string, h http.HandlerFunc)
}

func RegisterSwaggerUIServer[T httpServerInterface](srv T, title, swaggerJSONPath string, basePath string) {
	swaggerHandler := newHandler(title, swaggerJSONPath, basePath)
	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

func RegisterSwaggerUIServerWithOption[T httpServerInterface](srv T, handlerOpts ...HandlerOption) {
	opts := handler2.NewConfig()

	for _, o := range handlerOpts {
		o(opts)
	}

	if opts.LocalOpenApiFile != "" {
		registerOpenApiLocalFileRouter(srv, opts)
	} else if len(opts.OpenApiData) != 0 {
		registerOpenApiMemoryDataRouter(srv, opts)
	}

	swaggerHandler := newHandlerWithConfig(opts)

	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

// var _openJsonFileHandler = &openJsonFileHandler{}

func registerOpenApiLocalFileRouter[T httpServerInterface](srv T, cfg *handler2.Config) {
	var _openJsonFileHandler = &openJsonFileHandler{}
	err := _openJsonFileHandler.LoadFile(cfg.LocalOpenApiFile)
	if err == nil {
		pattern := strings.TrimRight(cfg.BasePath, "/") + "/openapi" + path.Ext(cfg.LocalOpenApiFile)
		cfg.SwaggerJSON = pattern
		srv.Handle(pattern, _openJsonFileHandler)
	} else {
		fmt.Println("load openapi file failed: ", err)
	}
}

func registerOpenApiMemoryDataRouter[T httpServerInterface](srv T, cfg *handler2.Config) {
	var _openJsonFileHandler = &openJsonFileHandler{}
	_openJsonFileHandler.Content = cfg.OpenApiData
	pattern := strings.TrimRight(cfg.BasePath, "/") + "/openapi." + cfg.OpenApiDataType
	cfg.SwaggerJSON = pattern
	srv.Handle(pattern, _openJsonFileHandler)
	cfg.OpenApiData = nil
}
