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

	affinity := 4
	args := AttrArgs{
		Owner:    "0xF0128825b0c518858971d8521498769148137936",
		Affinity: &affinity,
		MinPower: "100000",
		MaxPower: "900000000000000",
	}

	res, err := cw.GetWizardsByAttributes(args)
	if err != nil {
		t.Fatalf("failed to get wizard by id. err=%+v", err)
	}

	if res == nil {
		t.Fatalf("failed to get response for wizard. res=%+v", res)
	}

}
