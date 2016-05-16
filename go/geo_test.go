package openrtb_test

import (
	"reflect"
	"testing"

	"gopkg.in/Vungle/openrtb"
	"gopkg.in/Vungle/openrtb/openrtbtest"
)

var GeoModelType = reflect.TypeOf(openrtb.Geo{})

func TestGeoMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "geo.json", GeoModelType)
}
