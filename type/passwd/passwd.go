package hosts

import (
	"strings"

	"github.com/go-zoox/core-utils/cast"
	"github.com/go-zoox/errors"
	"github.com/go-zoox/fs"
)

type Passwd struct {
	FilePath string
	Mapping  map[string]*PasswdItem
}

func New(filepath string) *Passwd {
	return &Passwd{
		FilePath: filepath,
		Mapping:  make(map[string]*PasswdItem),
	}
}

func (p *Passwd) Load() error {
	lines, err := fs.ReadFileLines(p.FilePath)
	if err != nil {
		return err
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		item, err := parsePasswdItem(line)
		if err != nil {
			return err
		}

		p.Mapping[item.User] = item
	}

	return nil
}

func (p *Passwd) Get(user string) (*PasswdItem, error) {
	if item, ok := p.Mapping[user]; ok {
		return item, nil
	}

	return nil, errors.Errorf("not found: %s", user)
}

func (p *Passwd) GetUID(user string) (int, error) {
	item, err := p.Get(user)
	if err != nil {
		return 0, err
	}

	return item.UID, nil
}

func (p *Passwd) GetGID(user string) (int, error) {
	item, err := p.Get(user)
	if err != nil {
		return 0, err
	}

	return item.GID, nil
}

func (p *Passwd) GetHome(user string) (string, error) {
	item, err := p.Get(user)
	if err != nil {
		return "", err
	}

	return item.Home, nil
}

func (p *Passwd) GetShell(user string) (string, error) {
	item, err := p.Get(user)
	if err != nil {
		return "", err
	}

	return item.Shell, nil
}

func (p *Passwd) Length() int {
	return len(p.Mapping)
}

type PasswdItem struct {
	User  string
	Pass  string
	UID   int
	GID   int
	Gecos string
	Home  string
	Shell string
}

func parsePasswdItem(text string) (*PasswdItem, error) {
	items := strings.Split(text, ":")
	if len(items) != 7 {
		return nil, errors.Errorf("invalid passwd item: %s", text)
	}

	return &PasswdItem{
		User:  items[0],
		Pass:  items[1],
		UID:   cast.ToInt(items[2]),
		GID:   cast.ToInt(items[3]),
		Gecos: items[4],
		Home:  items[5],
		Shell: items[6],
	}, nil
}
