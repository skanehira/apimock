package server

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/skanehira/mockapi/app/common"
	"github.com/skanehira/mockapi/app/config"
	"github.com/skanehira/mockapi/app/db"
)

type Server struct {
	db     *db.DB
	config *config.Config
	net.Listener
	http.Server
}

type ErrResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func New(db *db.DB, c *config.Config) *Server {
	return &Server{
		config: c,
		db:     db,
	}
}

func (s *Server) Run() {
	switch s.config.Protocol {
	case "http":
		s.start()
	case "https":
		s.startTLS()
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	end, err := s.GetEndpoint(r.URL.Path, r.Method)
	if err != nil {
		cause := fmt.Sprintf("request url:%s, method:%s", r.URL.Path, r.Method)
		body := common.NewErrNotFoundEndpoint(cause).Error()
		s.NewErrorResponse(w, http.StatusNotFound, body)
		return
	}

	s.NewResponse(w, end.ResponseHeaders, end.ResponseStatus, end.ResponseBody)
}

func (s *Server) NewResponse(w http.ResponseWriter, headers string, status int, body string) {
	for _, header := range strings.Split(headers, "\n") {
		kv := strings.Split(header, ":")
		w.Header().Set(kv[0], kv[1])
	}

	w.WriteHeader(status)
	w.Write([]byte(body))
}

func (s *Server) NewErrorResponse(w http.ResponseWriter, status int, body string) {
	resp := &ErrResponse{
		Status:  status,
		Message: body,
	}

	json, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
		return
	}

	w.WriteHeader(status)
	w.Write(json)
}

func (s *Server) AddEndpoint(endpoint *db.Endpoint) error {
	if err := s.db.RegistEndpoint(endpoint); err != nil {
		return err
	}

	return nil
}

func (s *Server) GetEndpoint(url, method string) (*db.Endpoint, error) {
	endpoint, err := s.db.FindEndpoint(url, method)
	if err != nil {
		return nil, err
	}

	return endpoint, nil
}

func (s *Server) GetEndpointList() ([]*db.Endpoint, error) {
	endpointList, err := s.db.FindEndpointList()
	if err != nil {
		return nil, err
	}

	return endpointList, nil
}

func (s *Server) newTLSlistener() {
	tlsConfig := new(tls.Config)
	tlsConfig.Certificates = make([]tls.Certificate, 1)

	var err error
	tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(s.config.CertFile, s.config.CertKeyFile)
	if err != nil {
		panic(common.NewErrLoadTLSFiles(err))
	}

	tlsConfig.BuildNameToCertificate()

	listener, err := tls.Listen("tcp", s.config.Port, tlsConfig)
	if err != nil {
		panic(common.NewErrListenServer(err))
	}

	s.Listener = listener
}

func (s *Server) start() {
	log.Printf("start http server in %s\n", s.config.Port)
	log.Fatal(http.ListenAndServe(s.config.Port, s))
}

func (s *Server) startTLS() {
	s.newTLSlistener()
	log.Printf("start https server in %s\n", s.config.Port)
	log.Fatal(s.Serve(s.Listener))
}
