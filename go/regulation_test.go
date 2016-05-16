package openrtb_test

import (
	"reflect"
	"testing"

	"gopkg.in/Vungle/openrtb"
	"gopkg.in/Vungle/openrtb/openrtbtest"
)

var RegulationModelType = reflect.TypeOf(openrtb.Regulation{})

func TestRegulationMarshalUnmarshal(t *testing.T) {
	openrtbtest.VerifyModelAgainstFile(t, "regulation.json", RegulationModelType)
}
