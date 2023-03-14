package timing

import (
	"testing"
	"time"
)

func TestSaveUser(t *testing.T) {
	now := time.Now()
	timeNow = func() time.Time {
		return now
	}

	user := User{}
	SaveUser(&user)
	if user.UpdatedAt != now {
		t.Errorf("user.Updated at = %v, want ~%v", user.UpdatedAt, now)
	}
}
