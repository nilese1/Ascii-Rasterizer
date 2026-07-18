package rasterizer

import (
	sc "github.com/nilese1/Ascii-Rasterizer/scene"
	"github.com/nilese1/Ascii-Rasterizer/vector"
)

type Triangle struct {
	world_a vector.Vec3
	world_b vector.Vec3
	world_c vector.Vec3

	normal_vec vector.Vec3
}

func CreateTriangle(world_a vector.Vec3, world_b vector.Vec3, world_c vector.Vec3, normal vector.Vec3) Triangle {
	return Triangle{world_a, world_b, world_c, normal.Normalise()}
}

func (tri *Triangle) GetWorldCenter() vector.Vec3 {
	x := (tri.world_a.X + tri.world_b.X + tri.world_c.X) / 3
	y := (tri.world_a.Y + tri.world_b.Y + tri.world_c.Y) / 3
	z := (tri.world_a.Z + tri.world_b.Z + tri.world_c.Z) / 3

	return vector.Vec3{X: x, Y: y, Z: z}
}

func (tri *Triangle) GetNormal() vector.Vec3 {
	return tri.normal_vec
}

func (tri *Triangle) Rotate(rot_x float64, rot_y float64, rot_z float64) Triangle {
	new_a := tri.world_a.Rotate(rot_x, rot_y, rot_z)
	new_b := tri.world_b.Rotate(rot_x, rot_y, rot_z)
	new_c := tri.world_c.Rotate(rot_x, rot_y, rot_z)
	new_normal := tri.normal_vec.Rotate(rot_x, rot_y, rot_z)

	return CreateTriangle(new_a, new_b, new_c, new_normal)
}

func (tri *Triangle) Translate(translation vector.Vec3) Triangle {
	new_a := vector.Add(tri.world_a, translation)
	new_b := vector.Add(tri.world_b, translation)
	new_c := vector.Add(tri.world_c, translation)

	return CreateTriangle(new_a, new_b, new_c, tri.normal_vec)
}

func (tri *Triangle) Scale(scale_factor float64) Triangle {
	new_a := tri.world_a.Mul(scale_factor)
	new_b := tri.world_b.Mul(scale_factor)
	new_c := tri.world_c.Mul(scale_factor)

	return CreateTriangle(new_a, new_b, new_c, tri.normal_vec)
}

func (tri *Triangle) PointInTri(point vector.Vec2, scene *sc.Scene) bool {
	a := convertToScreen(tri.world_a, scene)
	b := convertToScreen(tri.world_b, scene)
	c := convertToScreen(tri.world_c, scene)

	ap := vector.Sub(point, a)
	bp := vector.Sub(point, b)
	cp := vector.Sub(point, c)

	ab := vector.Sub(b, a)
	bc := vector.Sub(c, b)
	ca := vector.Sub(a, c)

	ab_out := ab.Rot90()
	bc_out := bc.Rot90()
	ca_out := ca.Rot90()

	return vector.VecsSameDir(ap, ab_out) && vector.VecsSameDir(bp, bc_out) && vector.VecsSameDir(cp, ca_out)
}
