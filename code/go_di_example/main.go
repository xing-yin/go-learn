package main

import (
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"go-learn/code/go_di_example/config"
	"go-learn/code/go_di_example/db"
	"go-learn/code/go_di_example/infra"
	"go-learn/code/go_di_example/service"
	"go.uber.org/dig"
	"net/http"
)

// 参考文章链接：https://blog.drewolson.org/dependency-injection-in-go

// The manual way: 不使用依赖注入
//func main() {
//	config := config.NewConfig()
//
//	db, err := db.NewDatabase(config)
//	if err != nil {
//		panic(err)
//	}
//
//	personRepository := infra.NewPersonRepository(db)
//
//	personService := service.NewPersonService(config, personRepository)
//
//	sever := NewServer(config, personService)
//
//	sever.Run()
//}

// 采用依赖注入的方式
func main() {
	container := BuildContainer()

	err := container.Invoke(func(server *Server) {
		server.Run()
	})
	if err != nil {
		panic(err)
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(config.NewConfig)
	container.Provide(db.NewDatabase)
	container.Provide(infra.NewPersonRepository)
	container.Provide(service.NewPersonService)
	container.Provide(NewServer)
	return container
}

type Server struct {
	config        *config.Config
	personService *service.PersonService
}

func NewServer(config *config.Config, personService *service.PersonService) *Server {
	return &Server{config: config, personService: personService}
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}
	httpServer.ListenAndServe()
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/people", s.findPeople)
	return mux
}

func (s *Server) findPeople(writer http.ResponseWriter, request *http.Request) {
	people := s.personService.FindAll()
	bytes, err := json.Marshal(people)
	if err != nil {
		return
	}

	writer.Header().Set("Content-Type", "application-json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}
