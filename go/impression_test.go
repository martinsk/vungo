package openrtb_test

import (
	"reflect"
	"testing"

	"gopkg.in/Vungle/openrtb"
	"gopkg.in/Vungle/openrtb/openrtbtest"
)

var ImpressionModelType = reflect.TypeOf(openrtb.Impression{})

func TestImpressionMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "impression.json", ImpressionModelType)
}
