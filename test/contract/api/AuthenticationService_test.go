package provider

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"testing"

	response "authentication/model/response"

	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

// The Provider verification
func TestPactProvider(t *testing.T) {
	godotenv.Load("../../../.env.development.local")
	go startInstrumentedProvider()

	pact := createPact()
	// Verify the Provider - Tag-based Published Pacts for any known consumers
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: os.Getenv("API"),
		PactURLs:        []string{os.Getenv("PACT_BROKER_BASE_URL")},
		ProviderVersion: "1.0.0",
		BrokerToken:     os.Getenv("PACT_BROKER_TOKEN"),
	})
	if err != nil {
		// t.Log(err)
	}
}

// Starts the provider API with hooks for provider states.
// This essentially mirrors the main.go file, with extra routes added.
func startInstrumentedProvider() {
	mux := GetHTTPHandler()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Printf("API starting: port %s (%s)", port, ln.Addr())
	log.Printf("API terminating: %v", http.Serve(ln, mux))

}

// Configuration / Test Data
var dir, _ = os.Getwd()
var logDir = fmt.Sprintf("%s/log", dir)
var port = strconv.Itoa(3001)

// Setup the Pact client.
func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "UsersCrudAuthentication",
		LogDir:   logDir,
		LogLevel: "INFO",
	}
}

// Authenticate return user info if credentials are correct
func Authenticate(w http.ResponseWriter, r *http.Request) {
	setResponse(w, http.StatusOK, []response.User{{Username: "admin", ID: 100001, Photo: "", Active: true}})
}

func setResponse(w http.ResponseWriter, statusCode int, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(content)
}

func GetHTTPHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/latest/users/auth", Authenticate)

	return mux
}
