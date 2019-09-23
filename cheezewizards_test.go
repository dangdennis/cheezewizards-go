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
