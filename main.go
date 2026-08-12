package function
import "net/http"

var r = newRouter()
func Main(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
