package openrtb_test

import (
	"reflect"
	"testing"

	"gopkg.in/Vungle/openrtb"
	"gopkg.in/Vungle/openrtb/openrtbtest"
)

var UserModelType = reflect.TypeOf(openrtb.User{})

func TestUserMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "user.json", UserModelType)
}
