// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapters

import (
	"fmt"
	"testing"
	"time"

	dtpb "google.golang.org/genproto/googleapis/type/datetime"
	durpb "google.golang.org/protobuf/types/known/durationpb"
)

func TestDateTime(t *testing.T) {
	for _, test := range []struct {
		name               string
		y, mo, d, h, mi, s int
		tz                 *dtpb.TimeZone
		offset             *durpb.Duration
	}{
		{"DateTimeTZ", 2012, 4, 21, 11, 30, 0, &dtpb.TimeZone{Id: "America/New_York"}, nil},
		{"DateTimeTZ", 2012, 4, 21, 11, 30, 0, nil, &durpb.Duration{Seconds: 3600 * 5}},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Get the starting object.
			dt := &dtpb.DateTime{
				Year:    int32(test.y),
				Month:   int32(test.mo),
				Day:     int32(test.d),
				Hours:   int32(test.h),
				Minutes: int32(test.mi),
				Seconds: int32(test.s),
			}
			if test.tz != nil {
				dt.TimeOffset = &dtpb.DateTime_TimeZone{TimeZone: test.tz}
			}
			if test.offset != nil {
				dt.TimeOffset = &dtpb.DateTime_UtcOffset{UtcOffset: test.offset}
			}

			// Convert to a time.Time.
			tm, err := DateTimeToNativeTime(dt)
			if err != nil {
				t.Fatal(err)
			}
			t.Run("ToTime", func(t *testing.T) {
				assertEqual(t, "Year", tm.Year(), test.y)
				assertEqual(t, "Month", tm.Month(), time.Month(test.mo))
				assertEqual(t, "Day", tm.Day(), test.d)
				assertEqual(t, "Hour", tm.Hour(), test.h)
				assertEqual(t, "Minute", tm.Minute(), test.mi)
				assertEqual(t, "Second", tm.Second(), test.s)
				if test.tz != nil {
					assertEqual(t, "TZ", tm.Location().String(), test.tz.GetId())
				}
				if test.offset != nil {
					assertEqual(t, "Offset", tm.Location().String(), fmt.Sprintf("UTC+%d", test.offset.GetSeconds()/3600))
				}
			})

			// Convert back to a duration.
			t.Run("ToDateTime", func(t *testing.T) {
				durPb, err := TimeToProtoDateTime(tm)
				if err != nil {
					t.Fatal(err)
				}
				assertEqual(t, "Year", durPb.GetYear(), int32(test.y))
				assertEqual(t, "Month", durPb.GetMonth(), int32(test.mo))
				assertEqual(t, "Day", durPb.GetDay(), int32(test.d))
				assertEqual(t, "Hour", durPb.GetHours(), int32(test.h))
				assertEqual(t, "Minute", durPb.GetMinutes(), int32(test.mi))
				assertEqual(t, "Second", durPb.GetSeconds(), int32(test.s))
				if test.tz != nil {
					assertEqual(t, "TZ", durPb.GetTimeZone().GetId(), test.tz.GetId())
				}
				if test.offset != nil {
					assertEqual(t, "Offset", durPb.GetUtcOffset().GetSeconds(), test.offset.GetSeconds())
				}
			})
		})
	}
}
