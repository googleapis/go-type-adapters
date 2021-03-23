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

	fpb "google.golang.org/genproto/googleapis/type/fraction"
)

func TestFractionRat(t *testing.T) {
	for _, test := range []struct {
		name  string
		num   int64
		denom int64
	}{
		{"Terminating Decimal", 1, 5},
		{"Non-terminating Decimal", 1, 3},
	} {
		t.Run(test.name, func(t *testing.T) {
			fraction := &fpb.Fraction{
				Numerator:   test.num,
				Denominator: test.denom,
			}
			rat := FractionToRat(fraction)
			t.Run("FractionToRat", func(t *testing.T) {
				assertEqual(t, "rat.Num", rat.Num().Int64(), test.num)
				assertEqual(t, "rat.Denom", rat.Denom().Int64(), test.denom)
			})
			t.Run("RatToFraction", func(t *testing.T) {
				fraction = RatToFraction(rat)
				assertEqual(t, "frac.Num", fraction.GetNumerator(), test.num)
				assertEqual(t, "frac.Denom", fraction.GetDenominator(), test.denom)
			})
		})
	}
}
