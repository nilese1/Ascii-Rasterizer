package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/nilese1/Ascii-Rasterizer/mesh"
	"github.com/nilese1/Ascii-Rasterizer/rasterizer"
	sc "github.com/nilese1/Ascii-Rasterizer/scene"
	"github.com/nilese1/Ascii-Rasterizer/vector"
)

var (
	MODEL_TRANSLATION = vector.Vec3{X: 0, Y: 0, Z: 5}
	MODEL_SCALE       = 1.5
)

func triInPixel(scene *sc.Scene, pixel_x uint32, pixel_y uint32, tris []rasterizer.Triangle) (bool, rasterizer.Triangle) {
	point := vector.Vec2{X: float64(pixel_x), Y: float64(pixel_y)}

	hit := false
	var nearest_tri rasterizer.Triangle

	for _, i := range tris {
		if i.PointInTri(point, scene) {
			if !hit || nearest_tri.GetWorldCenter().Z < nearest_tri.GetWorldCenter().Z {
				nearest_tri = i
			}

			hit = true
		}
	}

	return hit, nearest_tri
}

func randAngle(max_angle float64) float64 {
	rand_f := rand.Float64()

	return rand_f * max_angle
}

func rotateModel(model *mesh.Model, translation vector.Vec3) {
	rot_x := randAngle(0.05)
	rot_y := randAngle(0.05)
	rot_z := randAngle(0.05)

	inverse_translate := translation.Mul(-1)

	model.Translate(inverse_translate)
	model.Rotate(rot_x, rot_y, rot_z)
	model.Translate(translation)
}

func renderModel(model *mesh.Model, scene *sc.Scene) [][]pixel {
	screen := make([][]pixel, scene.ScreenHeight)
	for i := range screen {
		screen[i] = make([]pixel, scene.ScreenWidth)
	}

	for y := range scene.ScreenHeight {
		for x := range scene.ScreenWidth {
			p := pixel{0, 0, 0, 0}

			hits, tri := triInPixel(scene, x, y, model.Triangles)
			if hits {
				normal := tri.GetNormal()
				light := (1 + vector.Dot3(&scene.SunDir, &normal)) * 0.5
				brightness_adjusted := model.Colour.Mul(light)

				p.r = int(brightness_adjusted.X)
				p.g = int(brightness_adjusted.Y)
				p.b = int(brightness_adjusted.Z)

				p.light = light
			}

			screen[y][x] = p
		}
	}

	return screen
}

func spinningObject(file_path string, scene *sc.Scene) {
	model := mesh.ParseModel(file_path)

	model.Scale(MODEL_SCALE)
	model.Translate(MODEL_TRANSLATION)

	var printTime time.Duration

	for {
		start := time.Now()

		rotateModel(&model, MODEL_TRANSLATION)
		screen := renderModel(&model, scene)

		renderTime := time.Since(start)

		totalTime := renderTime + printTime

		frametimeHeader := fmt.Sprintf("Render Time: %v", renderTime)
		printtimeHeader := fmt.Sprintf("Print Time: %v", printTime)
		totaltimeHeader := fmt.Sprintf("Total Time: %v", totalTime)
		fpsHeader := fmt.Sprintf("FPS: %v", float32(time.Second/totalTime))

		start = time.Now()
		printScreen(screen, scene, frametimeHeader, printtimeHeader, totaltimeHeader, fpsHeader)
		printTime = time.Since(start)
	}
}

func cleanup(scene *sc.Scene) {
	showCursor()
	moveCursor(scene.ScreenHeight, false)
}

func main() {
	scene := &sc.Scene{
		ScreenWidth:     100,
		ScreenHeight:    40,
		CamFOV:          math.Pi / 3,
		ViewPlaneHeight: math.Tan(math.Pi / 6),

		SunDir: vector.Vec3{X: 1, Y: 0.2, Z: -0.3}.Normalise(),
	}

	defer cleanup(scene)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	hideCursor()
	go spinningObject("models/suzanne", scene)

	<-sig
}
