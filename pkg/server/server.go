package server

import (
	"encoding/json"
	"flag"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jeotic/nmapper/pkg"
	"github.com/jeotic/nmapper/pkg/nmap"
	"github.com/jeotic/nmapper/pkg/parser"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
)

const StaticDir = "/web/public/"

type Server struct {
	Env    pkg.ENV
	Router *mux.Router
}

func (s *Server) Init() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "Javascript files")
	flag.Parse()

	db, err := gorm.Open("sqlite3", "./dev-database.db")

	if err != nil {
		log.Fatal(err)
	}
	db = db.Set("gorm:auto_preload", true)

	dialect := goqu.Dialect("sqlite3")

	env := pkg.ENV{
		DB:      db,
		Builder: dialect,
	}

	s.Env = env

	s.Router = mux.NewRouter().StrictSlash(true)

	s.initRoutes()
}

func (s *Server) initRoutes() {
	s.Router.HandleFunc("/runs", s.GetRuns).Methods("GET")
	s.Router.HandleFunc("/runs/{id}", s.GetRun).Methods("GET")
	s.Router.HandleFunc("/runs/{run_id}/tasks", s.GetTasks).Methods("GET")
	s.Router.HandleFunc("/runs/{run_id}/hosts", s.GetHosts).Methods("GET")
	s.Router.HandleFunc("/runs/{run_id}/hosts/{host_id}/ports", s.GetPorts).Methods("GET")
	s.Router.HandleFunc("/runs/{run_id}/hosts/{host_id}/names", s.GetHostNames).Methods("GET")
	s.Router.HandleFunc("/upload", s.UploadXML).Methods("POST")
}

func (s *Server) Run(addr string) {
	handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(addr, handlers.CORS()(s.Router)))
}

func (s *Server) GetRuns(w http.ResponseWriter, r *http.Request) {
	addr := r.URL.Query().Get("addr")

	var runs []nmap.Run

	// TODO: Clean up this section
	if addr != "" {
		var run nmap.Run = nmap.Run{}
		var hostAddrs nmap.HostAddress = nmap.HostAddress{}
		var runTblName string = run.TableName()
		var hostTblName string = hostAddrs.TableName()

		s.Env.DB.LogMode(true).
			Table(runTblName).
			Select("DISTINCT "+runTblName+".*").
			Joins("left join "+hostTblName+" on "+hostTblName+".run_id = "+runTblName+".id").
			Where(hostTblName+".addr = ?", addr).
			Scan(&runs)
	} else {
		s.Env.DB.Find(&runs)
	}

	err := json.NewEncoder(w).Encode(runs)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Server) GetRun(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var run nmap.Run

	s.Env.DB.First(&run, id)

	err = json.NewEncoder(w).Encode(run)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Server) GetTasks(w http.ResponseWriter, r *http.Request) {
	run_id, err := strconv.ParseInt(mux.Vars(r)["run_id"], 10, 64)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tasks []nmap.Task

	s.Env.DB.Where("run_id = ?", run_id).Find(&tasks)

	err = json.NewEncoder(w).Encode(tasks)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Server) GetHosts(w http.ResponseWriter, r *http.Request) {
	run_id, err := strconv.ParseInt(mux.Vars(r)["run_id"], 10, 64)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var hosts []nmap.Host

	s.Env.DB.Where("run_id = ?", run_id).Find(&hosts)

	err = json.NewEncoder(w).Encode(hosts)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Server) GetPorts(w http.ResponseWriter, r *http.Request) {
	run_id, err := strconv.ParseInt(mux.Vars(r)["run_id"], 10, 64)
	host_id, err := strconv.ParseInt(mux.Vars(r)["host_id"], 10, 64)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var ports []nmap.HostPort

	s.Env.DB.Where("run_id = ? AND host_id = ?", run_id, host_id).Find(&ports)

	err = json.NewEncoder(w).Encode(ports)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Server) GetHostNames(w http.ResponseWriter, r *http.Request) {
	run_id, err := strconv.ParseInt(mux.Vars(r)["run_id"], 10, 64)
	host_id, err := strconv.ParseInt(mux.Vars(r)["host_id"], 10, 64)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var hostNames []nmap.HostName

	s.Env.DB.Where("run_id = ? AND host_id = ?", run_id, host_id).Find(&hostNames)

	err = json.NewEncoder(w).Encode(hostNames)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Server) UploadXML(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("xml")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	nmap_id, err := parser.ParseReader(s.Env, file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]int64{"Id": nmap_id}

	err = json.NewEncoder(w).Encode(data)
}
