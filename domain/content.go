package domain

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type Inscription struct {
	ContentType string
	Data        []byte
}

func (i Inscription) Prepare() string {
	return fmt.Sprintf("data:%s;base64,%s", i.ContentType, base64.RawStdEncoding.EncodeToString(i.Data))
}

func (i Inscription) Hash() string {
	hashRaw := crypto.Keccak256(i.Data)
	return hex.EncodeToString(hashRaw)
}

func isAscii(characters []byte) bool {
	for i := 0; i < len(characters); i++ {
		if characters[i] >= 32 && characters[i] <= 126 {
			continue
		} else {
			return false
		}
	}
	return true
}

func DecodeContentURI(c string) (*Inscription, error) {
	c = strings.TrimSpace(c)
	if !strings.HasPrefix(c, "data:") {
		return nil, fmt.Errorf("prefix not found")
	}

	c = strings.TrimPrefix(c, "data:")

	parts := strings.SplitN(c, ",", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("data separator not found")
	}

	meta := parts[0]
	data := parts[1]

	parts = strings.Split(meta, ";")
	mime := parts[0]
	if mime == "" {
		mime = "text/plain"
	}

	parts = parts[1:]
	isBase64 := false
	for _, p := range parts {
		if p == "base64" {
			isBase64 = true
			break
		}
	}

	if isBase64 {
		data = strings.TrimRight(data, "=")
		decoded, err := base64.RawStdEncoding.DecodeString(data)
		if err != nil {
			return nil, fmt.Errorf("failed to decode base64: %w", err)

		}
		return &Inscription{
			ContentType: mime,
			Data:        decoded,
		}, nil
	}

	return &Inscription{
		ContentType: mime,
		Data:        []byte(data),
	}, nil
}
