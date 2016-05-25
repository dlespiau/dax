package math

import (
	"testing"
)

func TestProject(t *testing.T) {
	t.Parallel()
	obj := &Vec3{1002, 960, 0}
	modelview := &Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 203, 1, 0, 1}
	projection := &Mat4{0.0013020833721384406, 0, 0, 0, -0, -0.0020833334419876337, -0, -0, -0, -0, -1, -0, -1, 1, 0, 1}
	initialX, initialY, width, height := 0, 0, 1536, 960
	win := Project(obj, modelview, projection, initialX, initialY, width, height)
	answer := &Vec3{1205.0000359117985, -1.0000501200556755, 0.5} // From glu.Project()

	if !win.EqualThreshold(answer, 1e-4) {
		var diff Vec3
		diff.SubOf(&win, answer)
		t.Errorf("Project does something weird, differs from expected by of %v", diff.Len())
	}

	objr := UnProject(&win, modelview, projection, initialX, initialY, width, height)
	if !objr.EqualThreshold(obj, 1e-4) {
		t.Errorf("UnProject(%v) != %v (got %v)", win, obj, objr)
	}
}

func TestLookAtV(t *testing.T) {
	t.Parallel()
	// http://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
	iden := Ident4()
	tests := []struct {
		Description     string
		Eye, Center, Up Vec3
		Expected        Mat4
	}{
		{
			"forward",
			Vec3{0, 0, 0},
			Vec3{0, 0, -1},
			Vec3{0, 1, 0},
			iden,
		},
		{
			"heading 90 degree",
			Vec3{0, 0, 0},
			Vec3{1, 0, 0},
			Vec3{0, 1, 0},
			Mat4{
				0, 0, -1, 0,
				0, 1, 0, 0,
				1, 0, 0, 0,
				0, 0, 0, 1,
			},
		},
		{
			"heading 180 degree",
			Vec3{0, 0, 0},
			Vec3{0, 0, 1},
			Vec3{0, 1, 0},
			Mat4{
				-1, 0, 0, 0,
				0, 1, 0, 0,
				0, 0, -1, 0,
				0, 0, 0, 1,
			},
		},
		{
			"attitude 90 degree",
			Vec3{0, 0, 0},
			Vec3{0, 0, -1},
			Vec3{1, 0, 0},
			Mat4{
				0, 1, 0, 0,
				-1, 0, 0, 0,
				0, 0, 1, 0,
				0, 0, 0, 1,
			},
		},
		{
			"bank 90 degree",
			Vec3{0, 0, 0},
			Vec3{0, -1, 0},
			Vec3{0, 0, -1},
			Mat4{
				1, 0, 0, 0,
				0, 0, 1, 0,
				0, -1, 0, 0,
				0, 0, 0, 1,
			},
		},
	}

	threshold := Pow(10, -2)
	for _, c := range tests {
		if r := LookAtV(&c.Eye, &c.Center, &c.Up); !r.EqualThreshold(&c.Expected, threshold) {
			t.Errorf("%v failed: LookAtV(%v, %v, %v) != %v (got %v)", c.Description, c.Eye, c.Center, c.Up, c.Expected, r)
		}

		if r := LookAt(c.Eye[0], c.Eye[1], c.Eye[2], c.Center[0], c.Center[1], c.Center[2], c.Up[0], c.Up[1], c.Up[2]); !r.EqualThreshold(&c.Expected, threshold) {
			t.Errorf("%v failed: LookAt(%v, %v, %v) != %v (got %v)", c.Description, c.Eye, c.Center, c.Up, c.Expected, r)
		}
	}
}

func TestOrtho(t *testing.T) {
	t.Parallel()
	iden := Ident4()
	tests := []struct {
		Left, Right, Bottom, Top, Near, Far float32
		Expected                            Mat4
	}{
		{
			-1.0, 1.0, -1.0, 1.0, 1.0, -1.0,
			iden,
		}, {
			-10.0, 10.0, -10.0, 10.0, 0.0, 100.0,
			Mat4{0.1, 0.0, 0.0, 0.0, 0.0, 0.1, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, 0.0, 0.0, -1.0, 1.0},
		}, {
			0.0, 10.0, 0.0, 10.0, 0.0, 100.0,
			Mat4{0.2, 0.0, 0.0, 0.0, 0.0, 0.2, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, -1.0, -1.0, -1.0, 1.0},
		},
	}

	for _, c := range tests {
		if r := Ortho(c.Left, c.Right, c.Bottom, c.Top, c.Near, c.Far); !r.EqualThreshold(&c.Expected, 1e-4) {
			t.Errorf("Ortho(%v, %v, %v, %v, %v, %v) != %v (got %v)", c.Left, c.Right, c.Bottom, c.Top, c.Near, c.Far, c.Expected, r)
		}
	}
}

func TestOrtho2D(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Left, Right, Bottom, Top float32
		Expected                 Mat4
	}{
		{
			-1.0, 1.0, -1.0, 1.0,
			Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -1, 0, 0, 0, 0, 1},
		}, {
			-10.0, 10.0, -10.0, 10.0,
			Mat4{0.1, 0.0, 0.0, 0.0, 0.0, 0.1, 0.0, 0.0, 0.0, 0.0, -1.0, 0.0, 0.0, 0.0, 0.0, 1.0},
		}, {
			0.0, 10.0, 0.0, 10.0,
			Mat4{0.2, 0.0, 0.0, 0.0, 0.0, 0.2, 0.0, 0.0, 0.0, 0.0, -1.0, 0.0, -1.0, -1.0, 0.0, 1.0},
		},
	}

	for _, c := range tests {
		if r := Ortho2D(c.Left, c.Right, c.Bottom, c.Top); !r.EqualThreshold(&c.Expected, 1e-4) {
			t.Errorf("Ortho2D(%v, %v, %v, %v) != %v (got %v)", c.Left, c.Right, c.Bottom, c.Top, c.Expected, r)
		}
	}
}

func TestPerspective(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Fovy, Aspect,
		Near, Far float32
		Expected Mat4
	}{
		{
			DegToRad(450.0), 1.0, -1.0, 1.0,
			Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, -1, 0, 0, 1, 0},
		}, {
			DegToRad(45.0), 4.0 / 3.0, 0.1, 100.0,
			Mat4{1.810660, 0.0, 0.0, 0.0, 0.0, 2.4142134, 0.0, 0.0, 0.0, 0.0, -1.002002, -1.0, 0.0, 0.0, -0.2002002, 0.0},
		}, {
			DegToRad(90.0), 16.0 / 9.0, -1.0, 1.0,
			Mat4{0.562500, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
		},
	}

	for _, c := range tests {
		if r := Perspective(c.Fovy, c.Aspect, c.Near, c.Far); !r.EqualThreshold(&c.Expected, 1e-4) {
			t.Errorf("Perspective(%v, %v, %v, %v) != %v (got %v)", c.Fovy, c.Aspect, c.Near, c.Far, c.Expected, r)
		}
	}
}

func TestFrustum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Left, Right,
		Bottom, Top,
		Near, Far float32
		Expected Mat4
	}{
		{
			-1.0, 1.0, -1.0, 1.0, 1.0, 2.0,
			Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -3, -1, 0, 0, -4, 0},
		},
		// TODO: more tests
	}

	for _, c := range tests {
		if r := Frustum(c.Left, c.Right, c.Bottom, c.Top, c.Near, c.Far); !r.EqualThreshold(&c.Expected, 1e-4) {
			t.Errorf("Frustum(%v, %v, %v, %v, %v, %v) != %v (got %v)", c.Left, c.Right, c.Bottom, c.Top, c.Near, c.Far, c.Expected, r)
		}
	}
}
