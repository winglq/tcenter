package registration

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func NewService() http.Handler {
	service := new(restful.WebService)
	service.
		Path("/registration").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	service.Route(service.GET("/").To(list))
	container := restful.NewContainer()
	container.Add(service)
	return container
}

func list(request *restful.Request, response *restful.Response) {
	hello := make([]string, 1)
	hello[0] = "hi"
	response.WriteEntity(hello)
}
