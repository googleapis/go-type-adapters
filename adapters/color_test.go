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
	"image/color"
	"testing"

	wpb "github.com/golang/protobuf/ptypes/wrappers"
	cpb "google.golang.org/genproto/googleapis/type/color"
)

func TestColorRGBA(t *testing.T) {
	for _, test := range []struct {
		name  string
		color *cpb.Color
		rgba  *color.RGBA
	}{
		{"White", &cpb.Color{Red: 1, Green: 1, Blue: 1}, &color.RGBA{R: 255, G: 255, B: 255, A: 255}},
		{"ExplicitAlpha", &cpb.Color{Red: 1, Green: 1, Blue: 1, Alpha: &wpb.FloatValue{Value: 1}}, &color.RGBA{R: 255, G: 255, B: 255, A: 255}},
		{"Red", &cpb.Color{Red: 1, Green: 0, Blue: 0}, &color.RGBA{R: 255, G: 0, B: 0, A: 255}},
		{"Float", &cpb.Color{Red: 0.5, Green: 0, Blue: 0}, &color.RGBA{R: 128, G: 0, B: 0, A: 255}},
		{"MultiFloat", &cpb.Color{Red: 0.25, Green: 0.5, Blue: 0}, &color.RGBA{R: 64, G: 128, B: 0, A: 255}},
		{"Black", &cpb.Color{Red: 0, Green: 0, Blue: 0}, &color.RGBA{R: 0, G: 0, B: 0, A: 255}},
		{"PartialAlpha", &cpb.Color{Red: 0, Green: 0.5, Blue: 0, Alpha: &wpb.FloatValue{Value: 0.5}}, &color.RGBA{R: 0, G: 128, B: 0, A: 128}},
		{"NoAlpha", &cpb.Color{Red: 0, Green: 0.5, Blue: 0, Alpha: &wpb.FloatValue{Value: 0}}, &color.RGBA{R: 0, G: 128, B: 0, A: 0}},
	} {
		t.Run(test.name, func(t *testing.T) {
			rgba := ColorToRGBA(test.color)
			t.Run("RGBA", func(t *testing.T) {
				assertEqual(t, "R", rgba.R, test.rgba.R)
				assertEqual(t, "G", rgba.G, test.rgba.G)
				assertEqual(t, "B", rgba.B, test.rgba.B)
				assertEqual(t, "A", rgba.A, test.rgba.A)
			})
			t.Run("Color", func(t *testing.T) {
				color := RGBAToColor(rgba)
				assertEqual(t, "Red", color.GetRed(), test.color.GetRed())
				assertEqual(t, "Green", color.GetGreen(), test.color.GetGreen())
				assertEqual(t, "Blue", color.GetBlue(), test.color.GetBlue())
				if test.color.Alpha == nil {
					assertEqual(t, "Alpha", color.GetAlpha().GetValue(), float32(1))
				} else {
					assertEqual(t, "Alpha", color.GetAlpha().GetValue(), test.color.GetAlpha().GetValue())
				}
			})
		})
	}
}
