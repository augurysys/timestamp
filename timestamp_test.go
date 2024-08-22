package timestamp

import (
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"testing"
	"time"
)

func TestMarshalJSON(t *testing.T) {
	tm := time.UnixMilli(3000)
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

// Test the NEW MashalJSON method for dates before 1970
func TestMarshalJSONOld(t *testing.T) {
	tm := time.UnixMilli(-3000)
	ts := Timestamp(tm)

	b, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	temp, err := strconv.Atoi(string(b))
	if err != nil {
		t.Error(err)
	}

	if temp != -3000 {
		t.Fail()
	}
}

// Check the NEW MarshalJSON with the NEW UnmarshalJSON methods for recent date time values
func TestUnmarshalJSON(t *testing.T) {
	tm := time.Now().UTC()
	ts := Timestamp(tm)

	b, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	var temp Timestamp

	if err := temp.UnmarshalJSON(b); err != nil {
		t.Error(err)
	}

	// we expect to lose the milliseconds part of the timestamp
	if temp != Timestamp(ts.Time().Truncate(time.Millisecond)) {
		t.Fail()
	}
	t.Log("temp", temp)
	t.Log("ts", ts)
}

// Check the NEW MarshalJSON with the NEW UnmarshalJSON methods for old dates (slightly after 1970)
func TestUnmarshalJSON1970(t *testing.T) {
	tm := time.Unix(3000, 0).UTC()
	ts := Timestamp(tm)

	b, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	t.Log(len(b))

	var temp Timestamp

	if err := temp.UnmarshalJSON(b); err != nil {
		t.Error(err)
	}

	if temp != ts {
		t.Fail()
	}
	t.Log("temp", temp)
	t.Log("ts", ts)
}

// Check the NEW MarshalJSON with the NEW UnmarshalJSON methods for date before 1970
func TestUnmarshalJSONBefore1970(t *testing.T) {
	tm := time.Unix(-3000, 666000000).UTC()
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
	t.Log("temp", temp)
	t.Log("ts", ts)
}

// Check the NEW MarshalJSON with the OLD UnmarshalJSON methods for recent date time values with milliseconds
func TestUnmarshalJSONWithMilliSec(t *testing.T) {
	tm := time.Now().UTC()
	ts := Timestamp(tm)

	b, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	var temp Timestamp

	if err := temp.UnmarshalJSONOld(b); err != nil {
		t.Error(err)
	}
	if temp != Timestamp(ts.Time().Truncate(time.Second)) {
		t.Fail()
	}

	t.Log("temp", temp)
	t.Log("ts", ts)
	t.Log("done")
}

// Check the NEW MarshalJSON with the OLD UnmarshalJSON methods for dates close to 1970
func TestUnmarshalJSONOld1970(t *testing.T) {
	tm := time.Unix(3, 666000000).UTC()
	ts := Timestamp(tm)

	b, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	var temp Timestamp

	if err := temp.UnmarshalJSONOld(b); err != nil {
		t.Error(err)
	}

	if temp != Timestamp(ts.Time().Truncate(time.Second)) {
		t.Fail()
	}
	t.Log("temp", temp)
	t.Log("ts", ts)
	t.Log("done")
}

// Check the OLD MarshalJSON with the NEW UnmarshalJSON methods for dates close to 1970
func TestOldMarshalWithNewUnmarshalJSON(t *testing.T) {
	tm := time.Unix(3, 666000000).UTC()
	ts := Timestamp(tm)

	b, err := ts.MarshalJSONOld()
	if err != nil {
		t.Error(err)
	}

	var temp Timestamp

	if err := temp.UnmarshalJSON(b); err != nil {
		t.Error(err)
	}

	// we expect to lose the milliseconds part of the timestamp
	if temp != Timestamp(ts.Time().Truncate(time.Second)) {
		t.Fail()
	}
	t.Log("temp", temp)
	t.Log("ts", ts)
	t.Log("done")
}

// Check the OLD MarshalJSON with the NEW UnmarshalJSON methods for recent date time values
func TestOldMarshalWithNewUnmarshalJSONRecentTime(t *testing.T) {
	tm := time.Now().UTC()
	ts := Timestamp(tm)

	b, err := ts.MarshalJSONOld()
	if err != nil {
		t.Error(err)
	}

	var temp Timestamp

	if err := temp.UnmarshalJSON(b); err != nil {
		t.Error(err)
	}

	if temp != Timestamp(ts.Time().Truncate(time.Second)) {
		t.Fail()
	}
	t.Log("temp", temp)
	t.Log("ts", ts)
	t.Log("done")
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

func TestIsEmpty(t *testing.T) {
	var dateTimeNil *Timestamp
	result := dateTimeNil.IsEmpty()
	if !result {
		t.Fail()
	}

	date := time.Time{}
	dateTime := Timestamp(date)
	result = dateTime.IsEmpty()
	if !result {
		t.Fail()
	}

	date = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	dateTime = Timestamp(date)
	result = dateTime.IsEmpty()
	if !result {
		t.Fail()
	}

	date = date.AddDate(1, 0, 0)
	dateTime = Timestamp(date)
	result = dateTime.IsEmpty()
	if result {
		t.Fail()
	}
}
