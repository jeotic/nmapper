package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jeotic/nmapper/pkg"
	"github.com/jeotic/nmapper/pkg/nmap"
	"net/http"
	"path/filepath"
	"strconv"
)

type Route struct {
	Name           string
	Method         []string
	Pattern        string
	HandlerWrapper *HandlerWrapper
}

type HandlerWrapper struct {
	pkg.ENV
	HandlerWrapperFunc func(pkg.ENV, http.ResponseWriter, *http.Request)
}

func (handler HandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.HandlerWrapperFunc(handler.ENV, w, r)
}

type Routes []Route

func GetRoutes(env pkg.ENV) Routes {
	return Routes{
		Route{
			"Home",
			[]string{"GET"},
			"/",
			&HandlerWrapper{env, Home},
		},
		Route{
			"GetRuns",
			[]string{"GET"},
			"/runs",
			&HandlerWrapper{env, GetRuns},
		},
		Route{
			"GetRun",
			[]string{"GET"},
			"/runs/{id}",
			&HandlerWrapper{env, GetRun},
		},
		Route{
			"GetTasks",
			[]string{"GET"},
			"/runs/{run_id}/tasks",
			&HandlerWrapper{env, GetTasks},
		},
		Route{
			"GetHosts",
			[]string{"GET"},
			"/runs/{run_id}/hosts",
			&HandlerWrapper{env, GetHosts},
		},
		Route{
			"GetHost",
			[]string{"GET"},
			"/runs/{run_id}/hosts/{id}",
			&HandlerWrapper{env, GetHost},
		},
	}
}

func AddRoutes(env pkg.ENV, router *mux.Router) {
	for _, route := range GetRoutes(env) {
		//Check all routes to make sure the users are properly authenticated
		router.
			Methods(route.Method...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerWrapper)
	}
}

func Home(env pkg.ENV, w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs("web/index.html")

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.ServeFile(w, r, path)
}

func GetRuns(env pkg.ENV, w http.ResponseWriter, r *http.Request) {
	runs, err := nmap.GetRuns(env)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(runs)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetRun(env pkg.ENV, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	run, err := nmap.GetRun(env, id)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(run)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetTasks(env pkg.ENV, w http.ResponseWriter, r *http.Request) {
	run_id, err := strconv.ParseInt(mux.Vars(r)["run_id"], 10, 64)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := nmap.GetTasks(env, run_id)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetHosts(env pkg.ENV, w http.ResponseWriter, r *http.Request) {

}

func GetHost(env pkg.ENV, w http.ResponseWriter, r *http.Request) {

}
