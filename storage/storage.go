package storage

import (
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/Includeoyasi/tgbot/lib/e"
)

type Storage interface {
	Save(p *Page) error
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
	PickRandom(UserName string) (*Page, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	// todo
	// two simple steps - writeString into h(new sha1 obj) and conver it with Sum()
	// actually I can to send the second string in Sum method as parameter, but ok

	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can not write hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can not write hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
