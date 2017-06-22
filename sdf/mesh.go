//-----------------------------------------------------------------------------
/*

 */
//-----------------------------------------------------------------------------

package sdf

//-----------------------------------------------------------------------------

type Mesh struct {
	Triangles []*Triangle3
	box       *Box3
}

func NewMesh(triangles []*Triangle3) *Mesh {
	return &Mesh{triangles, nil}
}

//-----------------------------------------------------------------------------
