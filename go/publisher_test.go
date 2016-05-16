package openrtb_test

import (
	"reflect"
	"testing"

	"gopkg.in/Vungle/openrtb"
	"gopkg.in/Vungle/openrtb/openrtbtest"
)

var PublisherModelType = reflect.TypeOf(openrtb.Publisher{})

func TestPublisherMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "publisher.json", PublisherModelType)
}
