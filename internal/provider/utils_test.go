package provider

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func TestUtils_HttpRespToMap(t *testing.T) {
	cases := []struct {
		respRaw     string
		expectOk    bool
		expectedMap map[string]interface{}
	}{
		{
			respRaw:  `{"id":"abc-123","createdAt":"2022-12-19T10:19:52.578Z","updatedAt":"2022-12-19T10:19:52.578Z","displayName":"Foo"}`,
			expectOk: true,
			expectedMap: map[string]interface{}{
				"id":          "abc-123",
				"displayName": "Foo",
				"createdAt":   "2022-12-19T10:19:52.578Z",
				"updatedAt":   "2022-12-19T10:19:52.578Z",
			},
		},
		{
			respRaw:  `{"id":1,"createdAt":"2022-12-19T10:19:52.578Z","updatedAt":"2022-12-19T10:19:52.578Z","displayName":"Bar"}`,
			expectOk: true,
			expectedMap: map[string]interface{}{
				"id":          1.00,
				"displayName": "Bar",
				"createdAt":   "2022-12-19T10:19:52.578Z",
				"updatedAt":   "2022-12-19T10:19:52.578Z",
			},
		},
		{
			respRaw:     `{}`,
			expectOk:    true,
			expectedMap: map[string]interface{}{},
		},
		{
			respRaw:  `InvalidJSON`,
			expectOk: false,
		},
	}

	for _, c := range cases {
		resp := http.Response{}
		resp.Body = io.NopCloser(strings.NewReader(c.respRaw))

		out, ok := httpRespToMap(&resp)
		if ok != c.expectOk {
			testOkErrorMessage(c.respRaw, c.expectOk, ok)
		}

		if ok && !reflect.DeepEqual(out, c.expectedMap) {
			testErrorMessage(c.respRaw, c.expectedMap, out)
		}
	}
}

func TestUtils_UtcTimeToString(t *testing.T) {
	cases := []struct {
		t      time.Time
		expect string
	}{
		{
			t:      time.Date(2000, 01, 01, 12, 00, 00, 00, time.UTC),
			expect: "2000-01-01T12:00:00Z",
		},
		{
			t:      time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			expect: "2009-11-10T23:00:00Z",
		},
	}

	for _, c := range cases {
		out := utcTimeToString(c.t)

		if out != c.expect {
			testErrorMessage(c.t, c.expect, out)
		}
	}
}

func TestUtils_ErrRespToDiag(t *testing.T) {
	cases := []struct {
		respRaw    string
		err        error
		expectDiag diag.Diagnostics
	}{
		{
			respRaw:    `InvalidJSON`,
			err:        errors.New("Cannot process request"),
			expectDiag: diag.FromErr(errors.New("Cannot process request")),
		},
		{
			respRaw:    `{}`,
			err:        errors.New("Cannot process request"),
			expectDiag: diag.FromErr(errors.New("Cannot process request")),
		},
		{
			respRaw: `{"title": "something went wrong"}`,
			err:     errors.New("Cannot process request"),
			expectDiag: diag.Diagnostics{
				{
					Severity: diag.Error,
					Summary:  "something went wrong",
				},
			},
		},
		{
			respRaw: `{"title": "something went wrong", "detail": "but we don't know what"}`,
			err:     errors.New("Cannot process request"),
			expectDiag: diag.Diagnostics{
				{
					Severity: diag.Error,
					Summary:  "something went wrong - but we don't know what",
				},
			},
		},
	}

	for _, c := range cases {
		resp := http.Response{}
		resp.Body = io.NopCloser(strings.NewReader(c.respRaw))

		d := errRespToDiag(c.err, &resp)
		if !reflect.DeepEqual(d, c.expectDiag) {
			testErrorMessage(c.respRaw, c.expectDiag, d)
		}
	}
}

func TestAccessClientApp_SingleListToMap(t *testing.T) {
	cases := []struct {
		list      []interface{}
		expectOk  bool
		expectVal map[string]interface{}
	}{
		{
			list:     []interface{}{},
			expectOk: false,
		},
		{
			list: []interface{}{
				map[string]interface{}{
					"key":       "value",
					"attribute": "data",
					"foo":       "bar",
				},
			},
			expectOk: true,
			expectVal: map[string]interface{}{
				"key":       "value",
				"attribute": "data",
				"foo":       "bar",
			},
		},
	}

	for _, c := range cases {
		m, ok := singleListToMap(c.list)

		if ok != c.expectOk {
			testOkErrorMessage(c.list, c.expectOk, ok)
		} else if ok && !reflect.DeepEqual(m, c.expectVal) {
			testErrorMessage(c.list, c.expectVal, m)
		}
	}
}

func TestUtils_ToSliceConv(t *testing.T) {
	cases := []struct {
		sliceString    []string
		sliceInterface []interface{}
	}{
		{
			sliceString:    []string{},
			sliceInterface: []interface{}{},
		},
		{
			sliceString:    []string{"Single Value"},
			sliceInterface: []interface{}{"Single Value"},
		},
		{
			sliceString:    []string{"Value 1", "Value 2", "Value 3"},
			sliceInterface: []interface{}{"Value 1", "Value 2", "Value 3"},
		},
	}

	for _, c := range cases {
		out := toSliceInterface(c.sliceString)

		if !reflect.DeepEqual(out, c.sliceInterface) {
			testErrorMessage(c.sliceString, c.sliceInterface, out)
		}
	}

	for _, c := range cases {
		out := toSliceString(c.sliceInterface)

		if !reflect.DeepEqual(out, c.sliceString) {
			testErrorMessage(c.sliceInterface, c.sliceString, out)
		}
	}
}

func _testErrorMessage(testValue, expectedValue, actualValue interface{}, isOkValue bool) string {
	if isOkValue {
		return fmt.Sprintf("Expected %#v to return ok %#v, but got %#v instead.", testValue, expectedValue, actualValue)
	} else {
		return fmt.Sprintf("Expected %#v to return %#v, but got %#v instead.", testValue, expectedValue, actualValue)
	}
}

func testOkErrorMessage(testValue, expectedValue, actualValue interface{}) string {
	return _testErrorMessage(testValue, expectedValue, actualValue, true)
}

func testErrorMessage(testValue, expectedValue, actualValue interface{}) string {
	return _testErrorMessage(testValue, expectedValue, actualValue, false)
}
