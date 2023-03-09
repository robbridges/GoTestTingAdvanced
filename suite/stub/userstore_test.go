package stub_test

import (
	"GoTestingAdvanced/suite"
	"GoTestingAdvanced/suite/stub"
	"GoTestingAdvanced/suite/suitetest"
	"testing"
)

// create variable of the userstore interface, and assign it to sub.userStore. This will fail if stub.UserStore{}
//does not implement suite.userStore

var _ suite.UserStore = &stub.UserStore{}

func TestUserStore(t *testing.T) {
	us := &stub.UserStore{}
	suitetest.UserStore(t, us, nil, nil)
}

func TestUserStore_withStruct(t *testing.T) {
	us := &stub.UserStore{}
	tests := suitetest.UserStoreSuite{
		UserStore: us,
	}
	tests.All(t)
}
