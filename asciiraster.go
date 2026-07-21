package asciiraster

import (
	"github.com/nilese1/Ascii-Rasterizer/rasterizer"
	sc "github.com/nilese1/Ascii-Rasterizer/scene"
	"github.com/nilese1/Ascii-Rasterizer/vector"
)

var (
	MODEL_TRANSLATION = vector.Vec3{X: 0, Y: 0, Z: 5}
	MODEL_SCALE       = 1.5
)

// this needs some optimization
func RenderModel(model *rasterizer.Model, scene *sc.Scene) [][]pixel {
	screen := make([][]pixel, scene.SceneHeight)
	for i := range screen {
		screen[i] = make([]pixel, scene.SceneWidth)
	}

	for y := range scene.SceneHeight {
		for x := range scene.SceneWidth {
			p := pixel{0, 0, 0, 0}

			hits, tri := scene.GetTriInPixel(x, y, model.Triangles)
			if hits {
				normal := tri.GetNormal()
				light := (1 + vector.Dot3(&scene.SunDir, &normal)) * 0.5
				brightness_adjusted := model.Colour.Scale(light)

				p.r = int(brightness_adjusted.GetR())
				p.g = int(brightness_adjusted.GetB())
				p.b = int(brightness_adjusted.GetG())

				p.light = light
			}

			screen[y][x] = p
		}
	}

	return screen
}

func Cleanup(scenes []sc.Scene) {
	showCursor()
	moveCursor(sc.GetTotalHeight(scenes), false)
}

// TODO: more robust error handling with parsing
func LoadObjFile(filepath string) *rasterizer.Model {
	model := rasterizer.ParseModel(filepath)

	model.Scale(MODEL_SCALE)
	model.Translate(MODEL_TRANSLATION.X, MODEL_TRANSLATION.Y, MODEL_TRANSLATION.Z)

	return model
}
