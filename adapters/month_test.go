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
	"testing"
	"time"

	mpb "google.golang.org/genproto/googleapis/type/month"
)

func TestMonth(t *testing.T) {
	for m := range mpb.Month_name {
		t.Run("Native", func(t *testing.T) {
			native := ToNativeMonth(mpb.Month(m))
			assertEqual(t, native.String(), int(native), int(m))
		})
		t.Run("PB", func(t *testing.T) {
			pb := ToProtoMonth(time.Month(m))
			assertEqual(t, pb.String(), int(pb), int(m))
		})
	}
}
