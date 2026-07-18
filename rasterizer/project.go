package rasterizer

import (
	sc "github.com/nilese1/Ascii-Rasterizer/scene"
	"github.com/nilese1/Ascii-Rasterizer/vector"
)

func convertToViewPlane(point vector.Vec3, viewPlaneHeight float64) vector.Vec2 {
	view_plane_y := point.Y / point.Z
	fraction_up_plane := view_plane_y / viewPlaneHeight

	view_plane_x := point.X / point.Z * 2 //multiply by 2 because ASCII char width is half of char height
	fraction_along_plane := view_plane_x / viewPlaneHeight

	return vector.Vec2{X: fraction_along_plane, Y: fraction_up_plane}
}

func convertToScreen(point vector.Vec3, scene *sc.Scene) vector.Vec2 {
	view_plane_point := convertToViewPlane(point, scene.ViewPlaneHeight)

	scale := float64(scene.ScreenHeight) / 2

	scaled_y := view_plane_point.Y * scale
	scaled_x := view_plane_point.X * scale

	offset_y := float64(scene.ScreenHeight)/2 - scaled_y
	offset_x := float64(scene.ScreenWidth)/2 + scaled_x

	return vector.Vec2{X: offset_x, Y: offset_y}
}
