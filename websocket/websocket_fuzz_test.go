package websocket

import (
	"testing"
	"time"
)

// FuzzPayload test multiple kind of payloads
func FuzzPayload(f *testing.F) {
	testCases := []string{"Hello World", "123235", "@@&ˆ$(@*#&$)", "\x90\v\aHello W0rd\x90\v\a", "\"'\""}

	for _, testCase := range testCases {
		f.Add(testCase)
	}

	r, err := Listen("3000")
	if err != nil {
		panic(err)
	}

	f.Fuzz(func(t *testing.T, a string) {
		t.Logf("Testing payload: %s", a)
		Send("3000", a)

		time.Sleep(time.Second / 10)
		if r.GetMessage() != a {
			t.Errorf("Expected %s, got %s", a, r.GetMessage())
		}
	})
}
