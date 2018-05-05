//-----------------------------------------------------------------------------
/*

Demonstration for Parametric Box/Case

*/
//-----------------------------------------------------------------------------

package main

import . "github.com/deadsy/sdfx/sdf"

//-----------------------------------------------------------------------------

func tab1() {
	tp := BoxTabParms{
		Wall:        3,
		Length:      20,
		Hole:        3.0,
		Orientation: "tl",
		Clearance:   0.05,
	}
	RenderSTL(BoxTab3D(&tp), 400, "tab.stl")
}

func box1() {

	bp := PanelBoxParms{
		Size:       V3{100.0, 60.0, 200.0},
		Wall:       3.0,
		Panel:      3.0,
		Rounding:   5.0,
		FrontInset: 5.0,
		BackInset:  5.0,
		Clearance:  0.05,
		Hole:       2.0,
		SideTabs:   "TbtbtbT",
	}

	box := PanelBox3D(&bp)

	RenderSTL(box[0], 400, "panel.stl")
	RenderSTL(box[1], 400, "top.stl")
	RenderSTL(box[2], 400, "bottom.stl")
}

//-----------------------------------------------------------------------------

func main() {
	//tab1()
	box1()
}

//-----------------------------------------------------------------------------