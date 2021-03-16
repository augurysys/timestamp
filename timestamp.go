// Package timestamp is used for marshaling/unmarshaling UNIX timestamps
// to/from JSON, GOB and BSON by implementing the appropriate interfaces for
// encoding/json, encoding/gob and labix.org/v2/mgo respectively.
package timestamp

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Timestamp is a named alias for time.Time,
// it represents a UNIX timestamp
// and provides functions for marshaling and unmarshaling both to/from JSON and
// to/from BSON
type Timestamp time.Time

// MarshalJSON defines how encoding/json marshals the object to JSON,
// the result is a string of the UNIX timestamp
func (t Timestamp) MarshalJSON() ([]byte, error) {
	ts := t.Time().Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

// UnmarshalJSON defines how encoding/json unmarshals the object from JSON,
// a UNIX timestamp string is converted to int which is used for the Timestamp
// object value
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	int64ts := int64(ts)
	if len(b) > 10 {
		//support for milisecond timestamps
		int64ts = int64(ts / 1000)
	}
	*t = Timestamp(time.Unix(int64ts, 0).UTC())

	return nil
}

// GetBSON defines how labix.org/v2/mgo marshals the object to BSON,
// the result is a time.Time object which is then handled by mgo
func (t Timestamp) GetBSON() (interface{}, error) {
	if t.Time().IsZero() {
		return nil, nil
	}

	return t.Time(), nil
}

// SetBSON defines how labix.org/v2/mgo unmarshals the object from BSON,
// the raw BSON data is unmarshaled to a time.Time object which is used for the
// Timestamp object value
func (t *Timestamp) SetBSON(raw bson.Raw) error {
	var tm time.Time

	if err := raw.Unmarshal(&tm); err != nil {
		return err
	}

	*t = Timestamp(tm)

	return nil
}

// String returns the string representation of the Timestamp object,
// it is equal to the time.Time string representation of the Timestamp object
// value
func (t Timestamp) String() string {
	return t.Time().String()
}

// Time returns a time.Time object with the same time value as the Timestamp
// object
func (t Timestamp) Time() time.Time {
	if time.Time(t).IsZero() {
		return time.Unix(0, 0)
	}

	return time.Time(t)
}

// Now returns a pointer to a Timestamp object with the current time,
// it is equal to creating a Timestamp object from time.Now()
func Now() *Timestamp {
	t := Timestamp(time.Now())
	return &t
}

// Unix calls the Unix() method of a time.Time object with the same time values
// as the timestamp object
func (t Timestamp) Unix() int64 {
	return t.Time().Unix()
}

// Time returns a pointer to a Timestamp object which is created
// from a time.Time object
func Time(t time.Time) *Timestamp {
	ts := Timestamp(t)
	return &ts
}

// Unix returns a pointer to a Timestamp object which is created from
// a UNIX timestamp
func Unix(sec, nsec int64) *Timestamp {
	t := time.Unix(sec, nsec).UTC()
	return Time(t)
}

// GobEncode returns a byte slice representing the encoding of the Timestamp
// object, it implements the GobEncoder interface
func (t Timestamp) GobEncode() ([]byte, error) {
	return t.Time().MarshalBinary()
}

// GobDecode decodes a Timestamp object from a byte slice
// and overwrites the receiver,
// it implements the GobDecoder interface
// GobDecode implements the gob.GobDecoder interface.
func (t *Timestamp) GobDecode(data []byte) error {
	var tm time.Time

	if err := tm.UnmarshalBinary(data); err != nil {
		return err
	}

	*t = Timestamp(tm)

	return nil
}

// MarshalXML defines how encoding/xml marshals the object to XML,
// the result is a string of the UNIX timestamp
func (t Timestamp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	ts := t.Time().Unix()
	stamp := fmt.Sprint(ts)

	return e.EncodeElement(stamp, start)
}

// UnmarshalXML defines how encoding/xml unmarshals the object from XML,
// a UNIX timestamp string is converted to int which is used for the Timestamp
// object value
func (t *Timestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	ts, err := strconv.Atoi(content)
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(int64(ts), 0))

	return nil
}

func (t Timestamp) ToMili() int64 {
	return t.Time().UnixNano() / int64(time.Millisecond)
}
