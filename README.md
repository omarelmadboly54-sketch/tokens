🏗️ Architecture & Component Overview
The project is structured to completely decouple token generation and verification logic from external libraries (like JWT packages), allowing you to easily swap implementations if needed.


token/
├── payload.go               # Core business models (e.g., Payload containing user/admin roles)
├── Maker.go                 # Token Maker interface 
└── Jwt_Maker.go             # Where the func are declared 


Key Components:
1.Maker Interface: A clean abstraction layer defining methods for creating tokens and verifying them.

2.JWTMaker: The concrete implementation of the Maker interface using JSON Web Tokens (JWT), complete with a constructor for initialization in main.go.

3.Payload: Contains the token's claims (such as username, permissions, and admin roles) used across your system to secure endpoints.

✨ Features
Interface-Driven Design: The Maker interface allows your application to depend on abstractions rather than concrete libraries.

Role-Based Access Control: Built-in Payload structure handling custom claims like admin roles for securing endpoints.

Easy Dependency Injection: JWTMaker comes with a clean constructor function, making it simple to inject into your HTTP handlers or middleware via main.go.

Testable & Modular: Easy to mock the Maker interface for unit testing your API endpoints and middleware.


🚀 Getting Started
Prerequisites
Go 1.18 or higher installed

Installation
Bash
# Clone the repository
git clone https://github.com/omarelmadboly54-sketch/tokens

# Navigate to the project directory
cd token

# Download dependencies
go mod tidy


💡 Code Structure Example
1. The Maker Interface
Go
type Maker interface {
    CreateToken(username string, duration time.Duration) (string, error)
    VerifyToken(token string) (*Payload, error)
}
2. Payload with Admin Role
Go
type Payload struct {
    Email       string    `json:"email"`
    Role        String      `json:"role"` // Used for securing endpoints
    IssuedAt    time.Time `json:"issued_at"`
    ExpiredAt   time.Time `json:"expired_at"`
}
