// Copyright (c) 2022, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/goki/mat32"
	vk "github.com/goki/vulkan"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/vgpu/vgpu"
	"github.com/goki/vgpu/vphong"
	"github.com/goki/vgpu/vshape"
)

func init() {
	// must lock main thread for gpu!  this also means that vulkan must be used
	// for gogi/oswin eventually if we want gui and compute
	runtime.LockOSThread()
}

var TheGPU *vgpu.GPU

func OpenImage(fname string) image.Image {
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		fmt.Printf("image: %s\n", err)
	}
	gimg, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
	}
	return gimg
}

func main() {
	glfw.Init()
	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())
	vk.Init()

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	window, err := glfw.CreateWindow(1280, 960, "vPhong Test", nil, nil)
	vgpu.IfPanic(err)

	// note: for graphics, require these instance extensions before init gpu!
	winext := window.GetRequiredInstanceExtensions()
	gp := vgpu.NewGPU()
	gp.AddInstanceExt(winext...)
	gp.Debug = true
	gp.Config("vPhong test")
	TheGPU = gp

	// gp.PropsString(true) // print

	surfPtr, err := window.CreateWindowSurface(gp.Instance, nil)
	if err != nil {
		log.Println(err)
		return
	}
	sf := vgpu.NewSurface(gp, vk.SurfaceFromPointer(surfPtr))

	fmt.Printf("format: %s\n", sf.Format.String())

	ph := &vphong.Phong{}
	sy := &ph.Sys
	sy.InitGraphics(sf.GPU, "vphong.Phong", &sf.Device)
	sy.ConfigRenderPass(&sf.Format, vgpu.Depth32)
	sf.SetRenderPass(&sy.RenderPass)
	ph.ConfigSys()
	sy.SetRasterization(vk.PolygonModeFill, vk.CullModeBackBit, vk.FrontFaceCounterClockwise, 1.0)

	destroy := func() {
		vk.DeviceWaitIdle(sf.Device.Device)
		ph.Destroy()
		sf.Destroy()
		gp.Destroy()
		window.Destroy()
		glfw.Terminate()
	}

	/////////////////////////////
	// Lights

	dark := color.RGBA{50, 50, 50, 50}
	ph.AddAmbientLight(vphong.NewGoColor3(dark))
	ph.AddDirLight(vphong.NewGoColor3(color.White), mat32.Vec3{0, 0, -1})
	// ph.AddPointLight(vphong.NewGoColor3(color.White), mat32.Vec3{-3, 3, -6}, .001, .001)
	// ph.AddSpotLight(vphong.NewGoColor3(color.White), mat32.Vec3{0, 5, 0}, mat32.Vec3{0, -1, 0}, 15, 45, .01, .01)

	/////////////////////////////
	// Meshes

	floor := vshape.NewPlane(mat32.Y, 20, 20)
	nVtx, nIdx := floor.N()
	ph.AddMesh("floor", nVtx, nIdx, false)

	cube := vshape.NewBox(1, 1, 1)
	nVtx, nIdx = cube.N()
	ph.AddMesh("cube", nVtx, nIdx, false)

	/////////////////////////////
	// Textures

	imgFiles := []string{"ground.png", "wood.png", "teximg.jpg"}
	imgs := make([]image.Image, len(imgFiles))
	for i, fnm := range imgFiles {
		imgs[i] = OpenImage(fnm)
		if i == 0 {
			ph.AddTexture(fnm, vphong.NewTexture(imgs[i], mat32.Vec2{20, 20}, mat32.Vec2{0, 0}))
		} else {
			ph.AddTexture(fnm, vphong.NewTexture(imgs[i], mat32.Vec2{1, 1}, mat32.Vec2{0, 0}))
		}
	}

	/////////////////////////////
	// Colors

	blue := color.RGBA{0, 0, 255, 255}
	blueTr := color.RGBA{0, 0, 128, 128}
	red := color.RGBA{255, 0, 0, 255}
	redTr := color.RGBA{128, 0, 0, 128}
	ph.AddColor("blue", vphong.NewColors(blue, color.Black, color.White, 100, 1))
	ph.AddColor("blueTr", vphong.NewColors(blueTr, color.Black, color.White, 30, 1))
	ph.AddColor("red", vphong.NewColors(red, color.Black, color.White, 30, 1))
	ph.AddColor("redTr", vphong.NewColors(redTr, color.Black, color.White, 30, 1))

	/////////////////////////////
	// Camera / Mtxs

	// This is the standard camera view projection computation
	campos := mat32.Vec3{0, 1, 5}
	view := vphong.CameraViewMat(campos, mat32.Vec3{0, 0, 0}, mat32.Vec3Y)

	aspect := float32(sf.Format.Size.X) / float32(sf.Format.Size.Y)
	var prjn mat32.Mat4
	prjn.SetVkPerspective(45, aspect, 0.01, 100)

	var model1 mat32.Mat4
	model1.SetIdentity()
	model1.SetRotationY(0.5)
	ph.AddMtxs("mtx1", &model1, view, &prjn)

	var model2 mat32.Mat4
	model2.SetIdentity()
	model2.SetTranslation(-2, 0, 0)
	ph.AddMtxs("mtx2", &model2, view, &prjn)

	var floortx mat32.Mat4
	floortx.SetTranslation(0, -2, -2)
	ph.AddMtxs("floortx", &floortx, view, &prjn)

	/////////////////////////////
	//  Config!

	ph.Config()

	/////////////////////////////
	//  Set Mesh values

	vtxAry, normAry, texAry, _, idxAry := ph.MeshFloatsByName("floor")
	floor.Set(vtxAry, normAry, texAry, idxAry)
	ph.ModMeshByName("floor")

	vtxAry, normAry, texAry, _, idxAry = ph.MeshFloatsByName("cube")
	cube.Set(vtxAry, normAry, texAry, idxAry)
	ph.ModMeshByName("cube")

	ph.Sync()

	updateMats := func() {
		view = vphong.CameraViewMat(campos, mat32.Vec3{0, 0, 0}, mat32.Vec3Y)
		ph.SetMtxsName("mtx1", &model1, view, &prjn)
		ph.SetMtxsName("mtx2", &model2, view, &prjn)
		ph.SetMtxsName("floortx", &floortx, view, &prjn)
		ph.Sync()
	}

	render1 := func() {
		ph.UseColorName("blue")
		ph.UseMtxsName("floortx")
		ph.UseMeshName("floor")
		// ph.UseNoTexture()
		ph.UseTextureName("ground.png")
		ph.Render()

		ph.UseColorName("blue")
		ph.UseMtxsName("mtx1")
		ph.UseMeshName("cube")
		// ph.UseTextureName("teximg.jpg")
		ph.UseNoTexture()
		ph.Render()

		ph.UseColorName("redTr")
		ph.UseMtxsName("mtx2")
		ph.UseMeshName("cube")
		ph.UseTextureName("teximg.jpg")
		// ph.UseNoTexture()
		ph.Render()
	}

	frameCount := 0
	stTime := time.Now()

	renderFrame := func() {
		idx := sf.AcquireNextImage()
		cmd := sy.CmdPool.Buff
		descIdx := 0 // if running multiple frames in parallel, need diff sets
		sy.ResetBeginRenderPass(cmd, sf.Frames[idx], descIdx)

		fcr := frameCount % 10
		_ = fcr

		campos.X = float32(frameCount) * 0.01
		campos.Z = 5 - float32(frameCount)*0.03
		updateMats()
		render1()

		// switch {
		// case fcr < 3:
		// 	scaleImg(fcr)
		// case fcr < 6:
		// 	copyImg(fcr - 3)
		// default:
		// 	fillRnd()
		// }
		frameCount++

		sy.EndRenderPass(cmd)

		sf.SubmitRender(cmd) // this is where it waits for the 16 msec
		sf.PresentImage(idx)

		eTime := time.Now()
		dur := float64(eTime.Sub(stTime)) / float64(time.Second)
		if dur > 10 {
			fps := float64(frameCount) / dur
			fmt.Printf("fps: %.0f\n", fps)
			frameCount = 0
			stTime = eTime
		}
	}

	glfw.PollEvents()
	renderFrame()
	glfw.PollEvents()

	exitC := make(chan struct{}, 2)

	fpsDelay := time.Second / 50
	fpsTicker := time.NewTicker(fpsDelay)
	for {
		select {
		case <-exitC:
			fpsTicker.Stop()
			destroy()
			return
		case <-fpsTicker.C:
			if window.ShouldClose() {
				exitC <- struct{}{}
				continue
			}
			glfw.PollEvents()
			renderFrame()
		}
	}
}
