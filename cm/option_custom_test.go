package cm

import (
	"encoding/json"
	"testing"
)

func TestOptionMarshalJSON(t *testing.T) {
	type TestStruct struct {
		Field Option[string] `json:"field"`
	}

	// Test marshaling None
	ts1 := TestStruct{Field: None[string]()}
	data1, err := json.Marshal(ts1)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	expected1 := `{"field":null}`
	if string(data1) != expected1 {
		t.Errorf("json.Marshal: got %s, expected %s", data1, expected1)
	}

	// Test marshaling Some
	ts2 := TestStruct{Field: Some("hello")}
	data2, err := json.Marshal(ts2)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	expected2 := `{"field":"hello"}`
	if string(data2) != expected2 {
		t.Errorf("json.Marshal: got %s, expected %s", data2, expected2)
	}

	// Test marshaling custom option type
	type OptionalI32 Option[int32]
	ts3 := OptionalI32(Some(int32(42)))
	data3, err := json.Marshal(ts3)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	expected3 := `42`
	if string(data3) != expected3 {
		t.Errorf("json.Marshal: got %s, expected %s", data3, expected3)
	}

	// Test marshaling nested option type
	type NestedStruct struct {
		Field Option[Option[int32]] `json:"field"`
	}
	ts4 := NestedStruct{Field: Some(Some(int32(42)))}
	data4, err := json.Marshal(ts4)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	expected4 := `{"field":42}`
	if string(data4) != expected4 {
		t.Errorf("json.Marshal: got %s, expected %s", data4, expected4)
	}
}

func TestOptionUnmarshalJSON(t *testing.T) {
	type TestStruct struct {
		Field Option[string] `json:"field"`
	}

	// Test unmarshaling None
	data1 := []byte(`{"field":null}`)
	var ts1 TestStruct
	if err := json.Unmarshal(data1, &ts1); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if got, want := ts1.Field.None(), true; got != want {
		t.Errorf("ts1.Field.None: %t, expected %t", got, want)
	}

	// Test unmarshaling None (not present)
	data2 := []byte(`{}`)
	var ts2 TestStruct
	if err := json.Unmarshal(data2, &ts2); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if got, want := ts2.Field.None(), true; got != want {
		t.Errorf("ts1.Field.None: %t, expected %t", got, want)
	}

	// Test unmarshaling Some
	data3 := []byte(`{"field":"hello"}`)
	var ts3 TestStruct
	if err := json.Unmarshal(data3, &ts3); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if got, want := ts3.Field.isSome, true; got != want {
		t.Errorf("ts2.Field.Some: %t, expected %t", got, want)
	}
	if got, want := ts3.Field.Value(), "hello"; got != want {
		t.Errorf("ts2.Field.Value: %v, expected %v", got, want)
	}

	// Test unmarshaling nested option type
	type NestedStruct struct {
		Field Option[Option[int32]] `json:"field"`
	}
	data5 := []byte(`{"field":42}`)
	var ns NestedStruct
	if err := json.Unmarshal(data5, &ns); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if got, want := ns.Field.isSome, true; got != want {
		t.Errorf("ns.Field.Some: %t, expected %t", got, want)
	}
	if got, want := ns.Field.Value(), Some(int32(42)); got != want {
		t.Errorf("ns.Field.Value: %v, expected %v", got, want)
	}
	// Deref the inner option to get the value
	if got, want := ns.Field.Value().Value(), int32(42); got != want {
		t.Errorf("ns.Field.Value.Value: %v, expected %v", got, want)
	}
}
