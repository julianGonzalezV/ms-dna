package main

// Main Es el ENTRY POINT de la aplicación
import (
	"flag"
	"fmt"
	"log"

	"ms-dna/pkg/dna/application"
	"ms-dna/pkg/dna/domain/repository"
	"ms-dna/pkg/dna/domain/service"
	"ms-dna/pkg/dna/infrastructure/repositoryimpl"
	"ms-dna/pkg/dna/infrastructure/rest"
	"ms-dna/shared/server"
	"ms-dna/shared/storageconn"

	"net/http"
	"os"
	"strconv"

	"github.com/apex/gateway"
)

/// initializeRepo returns a repository based on database type name
func initializeRepo(database *string) repository.DnaRepository {
	switch *database {
	case "mongo":
		return newMongoRepository()
	default:
		return nil
	}
}

/// newMongoRepository returns the mongoDB implementation
func newMongoRepository() repository.DnaRepository {
	mongoAddr := os.Getenv("DATABASE_CONN")
	mongoClient := storageconn.Connect(mongoAddr)
	return repositoryimpl.New(mongoClient)
}

// ClientHandler set up all dependencies
func ClientHandler() {
	var (
		defaultHost    = os.Getenv("CLIENTAPI_SERVER_HOST")
		defaultPort, _ = strconv.Atoi(os.Getenv("CLIENTAPI_SERVER_PORT"))
		dbDriver       = os.Getenv("DATABASE_DRIVER")
	)
	host := flag.String("host", defaultHost, "define host of the server")
	port := flag.Int("port", defaultPort, "define port of the server")
	database := flag.String("database", dbDriver, "initialize the api using the given db engine")

	dnaRepository := initializeRepo(database)
	// Injecting services and repos to Application Layer
	dnaUseCase := application.New(service.New(dnaRepository))

	httpAddr := fmt.Sprintf("%s:%d", *host, *port)

	// Injecting server configuration
	dnaRoute := rest.New(dnaUseCase)
	server := server.New(dnaRoute)

	// Next two lines are for AWS Conf

	http.Handle("/", server.Router())
	log.Fatal(gateway.ListenAndServe(httpAddr, nil))

	// Next line is for Local conf
	//log.Fatal(http.ListenAndServe(httpAddr, server.Router()))
	fmt.Println("The server is running", httpAddr)

}

func main() {
	fmt.Println("V1.0.0")
	ClientHandler()
}
