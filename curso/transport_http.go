package curso

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler hacer llamados de los metdos*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()
	addCursoHandler := kithttp.NewServer(
		makeAddCursoEndpoint(s),
		addCursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addCursoHandler)
	return r
}

func addCursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addCursoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
