package messages_test

import (
	"net/http/httptest"
	"testing"

	"github.com/samze/pratter/messages"
	"github.com/samze/pratter/messages/fakes"
)

func TestAddsNewUser(t *testing.T) {
	fakeStore := new(fakes.FakeMessageStore)

	req := httptest.NewRequest("POST", "http://example.com/users/sam/messages", nil)
	w := httptest.NewRecorder()
	handler := messages.AddHandler{fakeStore}

	handler.ServeHTTP(w, req)

	if fakeStore.AddMessageCallCount() != 1 {
		t.Error("adding message wasn't called")
	}
}
