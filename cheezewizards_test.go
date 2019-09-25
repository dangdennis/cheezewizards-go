package cheezewizards

import (
	"os"
	"testing"
)

func TestGetWizardByID(t *testing.T) {
	key := os.Getenv("CHEEZEWIZARDS_KEY")
	email := os.Getenv("CHEEZEWIZARDS_EMAIL")

	if key == "" {
		t.Fatalf("require valid cheezewizard api key")
	}

	if email == "" {
		t.Fatalf("require valid cheezewizard email")
	}

	cw := NewCheezeWizards(key, email)
	res, err := cw.GetWizardByID(5)
	if err != nil {
		t.Fatalf("failed to get wizard by id. err=%+v", err)
	}

	if res == nil {
		t.Fatalf("failed to get response for wizard. res=%+v", res)
	}
}

func TestGetWizardsByAttributes(t *testing.T) {
	key := os.Getenv("CHEEZEWIZARDS_KEY")
	email := os.Getenv("CHEEZEWIZARDS_EMAIL")

	if key == "" {
		t.Fatalf("require valid cheezewizard api key")
	}

	if email == "" {
		t.Fatalf("require valid cheezewizard email")
	}

	cw := NewCheezeWizards(key, email)

	res, err := cw.GetWizardsByAttributes("0xF0128825b0c518858971d8521498769148137936", "4", "100000", "200000")
	if err != nil {
		t.Fatalf("failed to get wizard by id. err=%+v", err)
	}

	if res == nil {
		t.Fatalf("failed to get response for wizard. res=%+v", res)
	}
}

func TestGetDuelByID(t *testing.T) {
	key := os.Getenv("CHEEZEWIZARDS_KEY")
	email := os.Getenv("CHEEZEWIZARDS_EMAIL")

	if key == "" {
		t.Fatalf("require valid cheezewizard api key")
	}

	if email == "" {
		t.Fatalf("require valid cheezewizard email")
	}

	cw := NewCheezeWizards(key, email)

	res, err := cw.GetDuelByID(1)
	if err != nil {
		t.Fatalf("failed to get duel by id. err=%+v", err)
	}

	if res == nil {
		t.Fatalf("failed to get response for duel. res=%+v", res)
	}

}

func TestGetDuelsByAttributes(t *testing.T) {
	key := os.Getenv("CHEEZEWIZARDS_KEY")
	email := os.Getenv("CHEEZEWIZARDS_EMAIL")

	if key == "" {
		t.Fatalf("require valid cheezewizard api key")
	}

	if email == "" {
		t.Fatalf("require valid cheezewizard email")
	}

	cw := NewCheezeWizards(key, email)

	res, err := cw.GetDuelsByAttributes("5991,5993", "false", "true", "4943216", "", "", "")
	if err != nil {
		t.Fatalf("failed to get duels. err=%+v", err)
	}

	if res == nil {
		t.Fatalf("failed to get response for duels. res=%+v", res)
	}

}
