package cache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	defaultBasePath = "/cache/"
)

type HTTPPool struct {
	self        string
	basePath    string
	mu          sync.Mutex             // guards peers and httpGetters
	httpGetters map[string]*httpGetter // keyed by e.g. "http://10.0.0.2:8008"
}

type httpGetter struct {
	baseURL string
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (httpPool *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", httpPool.self, fmt.Sprintf(format, v...))
}

func (httpPool *HTTPPool) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if !strings.HasPrefix(request.URL.Path, httpPool.basePath) {
		panic("HTTPPool serving unexpected path: " + request.URL.Path)
	}
	httpPool.Log("%s %s", request.Method, request.URL.Path)

	parts := strings.SplitN(request.URL.Path[len(httpPool.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(responseWriter, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(responseWriter, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/octet-stream")
	responseWriter.Write(view.ByteSlice())
}
