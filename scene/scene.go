package scene

import "github.com/nilese1/Ascii-Rasterizer/vector"

type Scene struct {
	ScreenWidth      uint32
	ScreenHeight     uint32
	CamFOV           float32
	ViewPlaneHeight  float64

	SunDir           vector.Vec3
}
