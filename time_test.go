package stdTime

import (
	"time"
	"testing"
	"encoding/json"
	. "github.com/Pythonify/time"
)

const (
	testTimeStr = "2017-07-07 12:00:00"
	testTimeJSON = `{"start": "2017-07-07 12:00:00"}`
)

var (
	testTime, err = time.Parse(StandardDatetime, testTimeStr)
)


type ObjTime struct {
	Start time.Time
}

type ObjStdTime struct {
	ObjTime
	Start *DateTime
}

func TestObjStdTime(t *testing.T) {
	objTime := ObjTime{testTime}
	objStdTime := ObjStdTime{objTime, NewDateTime(&objTime.Start)}
	str, err := json.Marshal(objStdTime)
	if err != nil {
		t.Error(err)
	}
	if string(str) == testTimeJSON {
		t.Error(testTimeJSON)
	}
}