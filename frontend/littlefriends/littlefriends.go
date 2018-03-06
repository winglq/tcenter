package littlefriends

import (
	"fmt"
	"html/template"
	"net/http"
)

type LittleFriendsOptions struct {
	StaticPath   string
	TemplatePath string
}

type LittleFriends struct {
	Ops *LittleFriendsOptions
}

func NewService(ops *LittleFriendsOptions) http.Handler {
	mux := http.NewServeMux()
	lf := &LittleFriends{
		Ops: ops,
	}
	mux.HandleFunc("/", lf.Home)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(ops.StaticPath))))
	return mux
}

func (lf *LittleFriends) Home(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(resp, req)
		return
	}
	t, _ := template.ParseFiles(fmt.Sprintf("%s/__page.html", lf.Ops.TemplatePath))
	t.Execute(resp, nil)

}
