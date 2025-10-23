package api

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	users    = make(map[int]*User)
	sessions = make(map[string]*Session)
	nextID   = 1
	mu       = sync.RWMutex{}
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Session struct {
	Token     string
	UserID    int
	ExpiresAt time.Time
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UserID  int    `json:"user_id,omitempty"`
}

func (s *Session) IsValid() bool {
	return time.Now().Before(s.ExpiresAt)
}

func lookupSession(token string) *Session {
	mu.RLock()
	defer mu.RUnlock()
	return sessions[token]
}

func generateSessionToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func findUserByEmail(email string) *User {
	mu.RLock()
	defer mu.RUnlock()
	for _, u := range users {
		if u.Email == email {
			return u
		}
	}
	return nil
}

func findUserByID(id int) *User {
	mu.RLock()
	defer mu.RUnlock()
	return users[id]
}

func CurrentUser(r *http.Request) (*User, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}

	session := lookupSession(cookie.Value)
	if session == nil || !session.IsValid() {
		return nil, errors.New("invalid or expired session")
	}

	user := findUserByID(session.UserID)
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func EmailFromRequest(r *http.Request) (string, error) {
	u, err := CurrentUser(r)
	if err != nil {
		return "", err
	}
	return u.Email, nil
}

func storeUser(email, hashedPassword string) int {
	mu.Lock()
	defer mu.Unlock()
	id := nextID
	nextID++
	users[id] = &User{
		ID:       id,
		Email:    email,
		Password: hashedPassword,
	}
	return id
}

func storeSession(token string, userID int, duration time.Duration) {
	mu.Lock()
	defer mu.Unlock()
	sessions[token] = &Session{
		Token:     token,
		UserID:    userID,
		ExpiresAt: time.Now().Add(duration),
	}
}

func deleteSession(token string) {
	mu.Lock()
	defer mu.Unlock()
	delete(sessions, token)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "Email and password required"})
		return
	}

	if findUserByEmail(req.Email) != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "User already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	userID := storeUser(req.Email, string(hashedPassword))

	token, err := generateSessionToken()
	if err != nil {
		http.Error(w, "Error generating session", http.StatusInternalServerError)
		return
	}

	storeSession(token, userID, 24*time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(AuthResponse{Success: true, Message: "Account created", UserID: userID})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user := findUserByEmail(req.Email)
	if user == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "Invalid credentials"})
		return
	}

	token, err := generateSessionToken()
	if err != nil {
		http.Error(w, "Error generating session", http.StatusInternalServerError)
		return
	}

	storeSession(token, user.ID, 24*time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{Success: true, Message: "Logged in successfully", UserID: user.ID})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "Not logged in", http.StatusUnauthorized)
		return
	}

	deleteSession(cookie.Value)

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{Success: true, Message: "Logged out successfully"})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		session := lookupSession(cookie.Value)
		if session == nil || !session.IsValid() {
			http.Error(w, "Session expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", session.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Me(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email, err := EmailFromRequest(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "unauthorized"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Success bool   `json:"success"`
		Email   string `json:"email"`
	}{Success: true, Email: email})
}
