module github.com/crayboi420/chirpy

go 1.22.2

replace github.com/crayboi420/chirpy/internal/database => ./internal/database

require github.com/crayboi420/chirpy/internal/database v0.0.0

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.22.0 // indirect
)
