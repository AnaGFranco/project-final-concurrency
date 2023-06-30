package main

import (
	"fmt"
	"github.com/bwmarrin/go-alone"
	"strings"
)

const secret = "abc123abc123abc123"

var secretKey []byte

// NewURLSigner criar um novo URLSigner
func NewURLSigner() {
	secretKey = []byte(secret)
}

// GenerateTokenFromString gerar token
func GenerateTokenFromString(data string) string {
	var urlToSign string

	s := goalone.New(secretKey, goalone.Timestamp)
	if strings.Contains(data, "?") {
		urlToSign = fmt.Sprintf("%s&hash=", data)
	} else {
		urlToSign = fmt.Sprintf("%s?hash=", data)
	}

	tokenBytes := s.Sign([]byte(urlToSign))
	token := string(tokenBytes)

	return token
}

// VerifyToken verifica um token
func VerifyToken(token string) bool {
	s := goalone.New(secretKey, goalone.Timestamp)
	_, err := s.Unsign([]byte(token))

	if err != nil {
		// a assinatura não é válida. O token foi adulterado, forjado ou talvez seja
		// nem mesmo um token! De qualquer forma, não é seguro usá-lo.
		return false
	}
	// validar hash
	return true

}
