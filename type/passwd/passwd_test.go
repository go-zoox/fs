package hosts

import (
	"testing"

	"github.com/go-zoox/fs"
)

func TestHosts(t *testing.T) {
	users := &Passwd{
		FilePath: fs.JoinPath(fs.CurrentDir(), "tests/passwd"),
		Mapping:  make(map[string]*PasswdItem),
	}

	if err := users.Load(); err != nil {
		t.Fatal(err)
	}

	// for key, user := range users.Mapping {
	// 	fmt.Printf("%s => %s | %d | %d | %s | %s | %s\n", key, user.Pass, user.UID, user.GID, user.Gecos, user.Home, user.Shell)
	// }

	if users.Length() != 28 {
		t.Fatal("invalid hosts length, expected 8, got", users.Length())
	}

	// root
	if v, err := users.GetHome("root"); err != nil || v != "/root" {
		t.Fatal("expected localhost => /root, but got", v, err)
	}

	if v, err := users.GetShell("root"); err != nil || v != "/bin/bash" {
		t.Fatal("expected localhost => /bin/bash, but got", v, err)
	}

	if v, err := users.GetUID("root"); err != nil || v != 0 {
		t.Fatal("expected localhost => 0, but got", v, err)
	}

	if v, err := users.GetGID("root"); err != nil || v != 0 {
		t.Fatal("expected localhost => 0, but got", v, err)
	}

	// sshd
	if v, err := users.GetHome("sshd"); err != nil || v != "/run/sshd" {
		t.Fatal("expected localhost => /run/sshd, but got", v, err)
	}

	if v, err := users.GetShell("sshd"); err != nil || v != "/usr/sbin/nologin" {
		t.Fatal("expected localhost => /usr/sbin/nologin, but got", v, err)
	}

	if v, err := users.GetUID("sshd"); err != nil || v != 106 {
		t.Fatal("expected localhost => 106, but got", v, err)
	}

	if v, err := users.GetGID("sshd"); err != nil || v != 65534 {
		t.Fatal("expected localhost => 65534, but got", v, err)
	}
}
