package module

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/rs/cors"
)

var csrfTokenKey = []byte("secret-key-for-csrf-token") // Secret key untuk token CSRF
var s = securecookie.New(csrfTokenKey, nil)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize router
	r := mux.NewRouter()

	// Define API endpoints
	r.HandleFunc("/generate-csrf-token", generateCSRFTokenHandler).Methods("GET")
	r.HandleFunc("/submit-form", submitFormHandler).Methods("POST")

	// Enable CORS for cross-origin requests
	handler := cors.Default().Handler(r)

	// Start the server
	http.ListenAndServe(":8080", handler)
}

// Fungsi untuk mengenerate CSRF token
func generateCSRFTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Membuat token acak
	token := generateCSRFToken()

	// Menyimpan token di cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "csrf_token",
		Value: token,
		Path:  "/",
	})

	// Mengirimkan token ke client agar bisa digunakan pada form atau request berikutnya
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"csrf_token": "%s"}`, token)))
}

// Fungsi untuk membuat token CSRF
func generateCSRFToken() string {
	token := randString(32) // Token acak sepanjang 32 karakter
	encodedToken, err := s.Encode("csrf_token", token)
	if err != nil {
		fmt.Println("Error encoding CSRF token:", err)
	}
	return encodedToken
}

// Fungsi untuk menangani form submission dengan validasi CSRF token
func submitFormHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil token CSRF yang dikirim dari form
	clientToken := r.FormValue("csrf_token")

	// Mengambil token CSRF dari cookie
	cookie, err := r.Cookie("csrf_token")
	if err != nil {
		http.Error(w, "CSRF token not found", http.StatusForbidden)
		return
	}

	// Menyamakan token yang dikirim dengan token yang ada di cookie
	var decodedToken string
	err = s.Decode("csrf_token", cookie.Value, &decodedToken)
	if err != nil || decodedToken != clientToken {
		http.Error(w, "Invalid CSRF token", http.StatusForbidden)
		return
	}

	// Jika valid, lanjutkan dengan pemrosesan form
	w.Write([]byte("Form submitted successfully!"))
}

// Fungsi untuk membuat string acak
func randString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var output = make([]byte, length)
	for i := range output {
		output[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(output)
}
