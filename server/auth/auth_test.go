package auth

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
)

func TestResp(t *testing.T) {
	var resp Resp
	err := json.Unmarshal([]byte(`{"id":"853c80ef3c3749fdaa49938b674adae6","name":"jeb_","properties":[{"name":"textures","value":"eyJ0aW1lc3RhbXAiOjE1NTk1NDM5MzMwMjUsInByb2ZpbGVJZCI6Ijg1M2M4MGVmM2MzNzQ5ZmRhYTQ5OTM4YjY3NGFkYWU2IiwicHJvZmlsZU5hbWUiOiJqZWJfIiwidGV4dHVyZXMiOnsiU0tJTiI6eyJ1cmwiOiJodHRwOi8vdGV4dHVyZXMubWluZWNyYWZ0Lm5ldC90ZXh0dXJlLzdmZDliYTQyYTdjODFlZWVhMjJmMTUyNDI3MWFlODVhOGUwNDVjZTBhZjVhNmFlMTZjNjQwNmFlOTE3ZTY4YjUifSwiQ0FQRSI6eyJ1cmwiOiJodHRwOi8vdGV4dHVyZXMubWluZWNyYWZ0Lm5ldC90ZXh0dXJlLzU3ODZmZTk5YmUzNzdkZmI2ODU4ODU5ZjkyNmM0ZGJjOTk1NzUxZTkxY2VlMzczNDY4YzVmYmY0ODY1ZTcxNTEifX19"}]}`), &resp)
	if err != nil {
		panic(err)
	}
	wantID := uuid.Must(uuid.Parse("853c80ef3c3749fdaa49938b674adae6"))

	// check UUID
	if resp.ID != wantID {
		t.Errorf("uuid doesn't match: %v, want %s", resp.ID, wantID)
	}

	// check name
	if resp.Name != "jeb_" {
		t.Errorf("name doesn't match: %s, want %s", resp.Name, "jeb_")
	}

	// check texture
	texture, err := resp.Texture()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(texture.TimeStamp)

	if texture.ID != wantID {
		t.Errorf("uuid doesn't match: %v, want %s", texture.ID, wantID)
	}

	if texture.Name != "jeb_" {
		t.Errorf("name doesn't match: %s, want %s", texture.Name, "jeb_")
	}

	const (
		wantSKIN = "http://textures.minecraft.net/texture/7fd9ba42a7c81eeea22f1524271ae85a8e045ce0af5a6ae16c6406ae917e68b5"
		wantCAPE = "http://textures.minecraft.net/texture/5786fe99be377dfb6858859f926c4dbc995751e91cee373468c5fbf4865e7151"
	)
	if texture.Textures.SKIN.URL != wantSKIN {
		t.Errorf("skin url not match: %s, want %s",
			texture.Textures.SKIN.URL,
			wantSKIN)
	}
	if texture.Textures.CAPE.URL != wantCAPE {
		t.Errorf("cape url not match: %s, want %s",
			texture.Textures.CAPE.URL,
			wantCAPE)
	}
}
