package performance

import (
	"encoding/json"
	"testing"
)

func BenchmarkCodegen(b *testing.B) {
	s := []byte(`{"Login": "user", "Password": "123"}`)
	type user struct {
		Login    string
		Password string
	}
	u := &user{}
	for i := 0; i < b.N; i++ {
		u = &user{}
		json.Unmarshal(s, u)
	}
}
