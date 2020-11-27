package docente

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
	addDocenteHandler := kithttp.NewServer(
		makeAddDocenteEndpoint(s),
		addDocenteRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addDocenteHandler)
	return r
}

func addDocenteRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addDocenteRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
