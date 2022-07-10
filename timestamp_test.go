package timestamp

import (
	"go.mongodb.org/mongo-driver/bson"
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
	tm := time.Unix(3000, 0).UTC()
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

func TestBSON(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	typ, result, err := ts.MarshalBSONValue()
	if err != nil {
		t.Error(err)
	}

	var tm2 time.Time
	rv := bson.RawValue{Type: typ, Value: result}
	if err := rv.Unmarshal(&tm2); err != nil {
		t.Error(err)
	}

	if tm2.UTC() != tm.UTC() {
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
	tm := time.Unix(3000, 0).UTC()
	ts := Unix(3000, 0)

	if ts.Time() != tm {
		t.Fail()
	}
}

func TestToMili(t *testing.T) {
	numSeconds := int64(3000)
	tm := time.Unix(numSeconds, 0)
	ts := Timestamp(tm)

	result := ts.ToMili()
	if result != numSeconds*1000 {
		t.Fail()
	}
}
