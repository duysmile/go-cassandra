package main

import (
	"encoding/json"
	"github.com/duysmile/go-cassandra/models"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/scylladb/gocqlx/v2"
	"log"
	"net/http"
	"time"
)

type heartBeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type Server struct {
	cassClient gocqlx.Session
}

func NewServer(session gocqlx.Session) *Server {
	return &Server{
		cassClient: session,
	}
}

func main() {
	cassandraSess := mustConnectCassandra()
	defer cassandraSess.Close()

	server := NewServer(cassandraSess)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartBeat)

	router.HandleFunc("/users", server.HandleUserAPI)

	log.Println("🌱Server start at port 8080🌱")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("cannot start server", err)
	}
}

func (s *Server) HandleUserAPI(w http.ResponseWriter, req *http.Request) {
	user := models.User{
		ID:        gocql.UUIDFromTime(time.Now()),
		Firstname: "Duy",
		Lastname:  "Nguyen",
		Age:       16,
		Email:     "duy210697@gmail.com",
		City:      "Danang",
	}
	q := s.cassClient.Query(models.UserTable.Insert()).BindStruct(user)
	if err := q.ExecRelease(); err != nil {
		log.Println("failed to insert user", err)
	}

	_, _ = w.Write([]byte("OK"))
}

func heartBeat(w http.ResponseWriter, req *http.Request) {
	_ = json.NewEncoder(w).Encode(heartBeatResponse{
		Status: "OK",
		Code:   200,
	})
}

func mustConnectCassandra() gocqlx.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "chatapi"
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal("cannot connect cassandra", err)
	}

	return session
}
