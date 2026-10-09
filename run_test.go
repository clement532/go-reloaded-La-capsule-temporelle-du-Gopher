package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "in.txt")
	out := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(in, []byte("ff (hex) red\n\nhello (up)\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if code := run([]string{in, out}); code != 0 {
		t.Fatalf("code %d, attendu 0", code)
	}
	got, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "255 red\n\nHELLO\n" {
		t.Errorf("sortie inattendue : %q", got)
	}
}

func TestRunEmptyFile(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "in.txt")
	out := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(in, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	if code := run([]string{in, out}); code != 0 {
		t.Fatalf("code %d, attendu 0", code)
	}
	got, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("le fichier de sortie devrait être vide, obtenu %q", got)
	}
}

func TestRunErrors(t *testing.T) {
	dir := t.TempDir()
	out := filepath.Join(dir, "out.txt")
	if code := run([]string{filepath.Join(dir, "absent.txt"), out}); code != 1 {
		t.Errorf("fichier absent : code %d, attendu 1", code)
	}
	if _, err := os.Stat(out); err == nil {
		t.Error("le fichier de sortie ne doit pas être créé")
	}
	if code := run([]string{"seul.txt"}); code != 1 {
		t.Errorf("1 argument : code %d, attendu 1", code)
	}
	if code := run(nil); code != 1 {
		t.Errorf("0 argument : code %d, attendu 1", code)
	}
}
