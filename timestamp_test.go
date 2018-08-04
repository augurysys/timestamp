package timestamp

import (
	"bytes"
	"encoding/gob"
	"strconv"
	"testing"
	"time"
)

func TestMarshalJSON(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	b, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	temp, err := strconv.Atoi(string(b))
	if err != nil {
		t.Error(err)
	}

	if temp != 3000 {
		t.Fail()
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	b, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	var temp Timestamp

	if err := temp.UnmarshalJSON(b); err != nil {
		t.Error(err)
	}

	if temp != ts {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	if tm.String() != ts.String() {
		t.Fail()
	}
}

func TestGetBSON(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	result, err := ts.GetBSON()
	if err != nil {
		t.Error(err)
	}

	if result != tm {
		t.Fail()
	}
}

func TestTime(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	if ts.Time() != tm {
		t.Fail()
	}
}

func TestNow(t *testing.T) {
	tm := time.Now()
	ts := Now()

	if tm.Unix() != ts.Time().Unix() {
		t.Fail()
	}
}

func TestUnix(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	if ts.Unix() != tm.Unix() {
		t.Fail()
	}
}

func TestFromTime(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Time(tm)

	if ts.Unix() != tm.Unix() {
		t.Fail()
	}
}

func TestFromUnix(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Unix(3000, 0)

	if ts.Time() != tm {
		t.Fail()
	}
}

func TestGobEncodeDecode(t *testing.T) {
	ts := Now()

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(ts); err != nil {
		t.Fatal(err)
	}

	var n Timestamp
	r := bytes.NewBuffer(b.Bytes())

	dec := gob.NewDecoder(r)
	if err := dec.Decode(&n); err != nil {
		t.Fatal(err)
	}

	if ts.Time() != n.Time() {
		t.Fail()
	}
}

func TestFromMili(t *testing.T) {
	ts := int64(1529065200999)
	result := FromMili(ts).Time().Month().String()
	if result != "June" {
		t.Fail()
	}
}

func TestToMili(t *testing.T) {
	numSeconds := int64(3000)
	tm := time.Unix(numSeconds, 0)
	ts := Timestamp(tm)

	result := ts.ToMili()
	if result != numSeconds * 1000 {
		t.Fail()
	}
}