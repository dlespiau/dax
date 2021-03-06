package main

import (
	"github.com/dlespiau/dax"
	"github.com/dlespiau/dax/geometry"
	"github.com/dlespiau/dax/material"
)

type sceneGraphBasic struct {
	dax.Scene

	sg *dax.SceneGraph
}

func (s *sceneGraphBasic) Setup() {
	camera := dax.NewPerspectiveCamera(70, 800./600., 1, 1000)
	camera.SetPosition(0, 0, 600)
	s.SetCamera(camera)

	s.sg = dax.NewSceneGraph()

	box := geometry.NewBox(100, 100, 100)
	material := material.NewColor(&dax.Color{R: 1.0, G: 1.0, B: 1.0, A: 1.0})

	node1 := s.CreateActor(box, material)
	node1.TranslateX(-250)

	node2 := s.CreateActor(box, material)

	node3 := s.CreateActor(box, material)
	node3.TranslateX(+250)

	s.sg.AddChildren(node1, node2, node3)
}

func (s *sceneGraphBasic) Update(time float64) {
	for i, child := range s.sg.GetChildren() {
		node := child.(*dax.Node)
		node.RotateX(0.02 * float32(i+1))
		node.RotateY(0.04 * float32(i+1))
	}
}

func (s *sceneGraphBasic) Draw(fb dax.Framebuffer) {
	fb.Draw(s.sg)
}

var gfxScenegraphExample = Example{
	Category:    CategoryGraphics,
	Name:        "Scene Graph",
	Description: "Display a few rotating cubes",
	Scene:       &sceneGraphBasic{},
}
