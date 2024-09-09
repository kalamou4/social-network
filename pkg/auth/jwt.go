package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const secretKey = "your-secret-key" // Utiliser une clé secrète sécurisée

func GenerateToken(userID string) (string, error) {
	// Générer un token aléatoire
	tokenBytes := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encoder le token en chaîne de caractères
	token := base64.StdEncoding.EncodeToString(tokenBytes)

	// Calculer le hachage
	expiration := time.Now().Add(time.Hour * 72).Unix()
	message := fmt.Sprintf("%s|%s|%d", token, userID, expiration)
	hash := computeHMAC(message, secretKey)

	// Ajouter le hachage au token
	tokenWithHash := fmt.Sprintf("%s|%s", message, hash)

	return tokenWithHash, nil
}

func computeHMAC(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func ParseToken(tokenWithHash string) (string, string, error) {
	parts := strings.Split(tokenWithHash, "|")
	if len(parts) != 4 {
		return "", "", fmt.Errorf("format du token invalide")
	}

	message := strings.Join(parts[:3], "|")
	providedHash := parts[3]
	calculatedHash := computeHMAC(message, secretKey)

	if providedHash != calculatedHash {
		return "", "", fmt.Errorf("token modifié")
	}

	token := parts[0]
	userID := parts[1]
	expiration, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return "", "", err
	}

	if time.Now().Unix() > expiration {
		return "", "", fmt.Errorf("token expiré")
	}

	return token, userID, nil
}
