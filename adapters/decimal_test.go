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

	dpb "google.golang.org/genproto/googleapis/type/decimal"
)

func TestDecimal(t *testing.T) {
	// Test convertion to big.Float.
	for v, want := range map[string]string{
		"123":          "123",
		"123.45":       "123.45",
		"-123.45":      "-123.45",
		"+123.45":      "123.45",
		"123.45e1000":  "1.2345e+1002",
		"123.45e+10":   "1.2345e+12",
		"123.45e01":    "1234.5",
		"123.45e-10":   "1.2345e-08",
		"123.45e-1000": "1.2345e-998",
		"123.45E10":    "1.2345e+12",
		"123.45e0":     "123.45",
	} {
		decimalPb := &dpb.Decimal{Value: v}
		float, err := DecimalToFloat(decimalPb)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Run("ToFloat/"+v, func(t *testing.T) {
			assertEqual(t, "FloatEqual", float.String(), want)
			t.Run("ToDecimal", func(t *testing.T) {
				d := FloatToDecimal(float)
				assertEqual(t, "DecimalEqual", d.GetValue(), want)
			})
		})
	}

	// Test conversion to float64.
	for v, want := range map[string]float64{
		"123":         123,
		"123.45":      123.45,
		"-123.45":     -123.45,
		"+123.45":     123.45,
		"123.45e100":  1.2345e+102,
		"123.45e+10":  1.2345e+12,
		"123.45e01":   1234.5,
		"123.45e-10":  1.2345e-08,
		"123.45e-100": 1.2345e-98,
		"123.45E10":   1.2345e+12,
		"123.45e0":    123.45,
	} {
		decimalPb := &dpb.Decimal{Value: v}
		float, _, err := DecimalToFloat64(decimalPb)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Run("ToFloat64/"+v, func(t *testing.T) {
			assertEqual(t, "Float64Equal", float, want)
			t.Run("ToDecimal", func(t *testing.T) {
				d := Float64ToDecimal(float)
				assertEqual(t, "DecimalEqual", d.GetValue(), fmt.Sprintf("%f", want))
			})
		})
	}

	// Test the error case for float64 conversion.
	t.Run("ErrorCase", func(t *testing.T) {
		_, _, err := DecimalToFloat64(&dpb.Decimal{Value: "invalid"})
		if err == nil {
			t.Errorf("Expected error, got %v.", err)
		}
	})
}
