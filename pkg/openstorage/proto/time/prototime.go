/*
Copyright © 2019 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package prototime

import (
	"time"

	google_duration "github.com/golang/protobuf/ptypes/duration"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

// TimeToTimestamp converts a go Time to a protobuf Timestamp.
func TimeToTimestamp(t time.Time) *google_protobuf.Timestamp {
	return &google_protobuf.Timestamp{
		Seconds: t.UnixNano() / int64(time.Second),
		Nanos:   int32(t.UnixNano() % int64(time.Second)),
	}
}

// TimestampToTime converts a protobuf Timestamp to a go Time.
func TimestampToTime(timestamp *google_protobuf.Timestamp) time.Time {
	if timestamp == nil {
		return time.Unix(0, 0).UTC()
	}
	return time.Unix(
		timestamp.Seconds,
		int64(timestamp.Nanos),
	).UTC()
}

// TimestampLess returns true if i is before j.
func TimestampLess(i *google_protobuf.Timestamp, j *google_protobuf.Timestamp) bool {
	if j == nil {
		return false
	}
	if i == nil {
		return true
	}
	if i.Seconds < j.Seconds {
		return true
	}
	if i.Seconds > j.Seconds {
		return false
	}
	return i.Nanos < j.Nanos
}

// Now returns the current time as a protobuf Timestamp.
func Now() *google_protobuf.Timestamp {
	return TimeToTimestamp(time.Now().UTC())
}

// DurationToProto converts a go Duration to a protobuf Duration.
func DurationToProto(d time.Duration) *google_duration.Duration {
	return &google_duration.Duration{
		Seconds: int64(d) / int64(time.Second),
		Nanos:   int32(int64(d) % int64(time.Second)),
	}
}

// DurationFromProto converts a protobuf Duration to a go Duration.
func DurationFromProto(duration *google_duration.Duration) time.Duration {
	if duration == nil {
		return 0
	}
	return time.Duration((duration.Seconds * int64(time.Second)) + int64(duration.Nanos))
}
