// Code generated by "goki generate ./..."; DO NOT EDIT.

package szalloc

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:       "goki.dev/vgpu/v2/szalloc.Idxs",
	Doc:        "Idxs contains the indexes where a given item image size is allocated\nthere is one of these per each ItemSizes",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"PctSize", &gti.Field{Name: "PctSize", Type: "mat32.Vec2", Doc: "percent size of this image relative to max size allocated", Directives: gti.Directives{}}},
		{"GpIdx", &gti.Field{Name: "GpIdx", Type: "int", Doc: "group index", Directives: gti.Directives{}}},
		{"ItemIdx", &gti.Field{Name: "ItemIdx", Type: "int", Doc: "item index within group (e.g., Layer)", Directives: gti.Directives{}}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "goki.dev/vgpu/v2/szalloc.SzAlloc",
	Doc:        "SzAlloc manages allocation of sizes to a spec'd maximum number\nof groups.  Used for allocating texture images to image arrays\nunder the severe constraints of only 16 images.\nOnly a maximum of MaxItemsPerGp items can be allocated per grouping.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"On", &gti.Field{Name: "On", Type: "bool", Doc: "true if configured and ready to use", Directives: gti.Directives{}}},
		{"MaxGps", &gti.Field{Name: "MaxGps", Type: "image.Point", Doc: "maximum number of groups in X and Y dimensions", Directives: gti.Directives{}}},
		{"MaxNGps", &gti.Field{Name: "MaxNGps", Type: "int", Doc: "maximum number of groups = X * Y", Directives: gti.Directives{}}},
		{"MaxItemsPerGp", &gti.Field{Name: "MaxItemsPerGp", Type: "int", Doc: "maximum number of items per group -- constraint is enforced in addition to MaxGps", Directives: gti.Directives{}}},
		{"ItemSizes", &gti.Field{Name: "ItemSizes", Type: "[]image.Point", Doc: "original list of item sizes to be allocated", Directives: gti.Directives{}}},
		{"UniqSizes", &gti.Field{Name: "UniqSizes", Type: "[]image.Point", Doc: "list of all unique sizes -- operate on this for grouping", Directives: gti.Directives{}}},
		{"UniqSzMap", &gti.Field{Name: "UniqSzMap", Type: "map[image.Point]int", Doc: "map of all unique sizes, with group index as value", Directives: gti.Directives{}}},
		{"UniqSzItems", &gti.Field{Name: "UniqSzItems", Type: "[]int", Doc: "indexes into UniqSizes slice, ordered by ItemSizes indexes", Directives: gti.Directives{}}},
		{"GpSizes", &gti.Field{Name: "GpSizes", Type: "[]image.Point", Doc: "list of allocated group sizes", Directives: gti.Directives{}}},
		{"GpAllocs", &gti.Field{Name: "GpAllocs", Type: "[][]int", Doc: "allocation of image indexes by group -- first index is group, second is list of items for that group", Directives: gti.Directives{}}},
		{"ItemIdxs", &gti.Field{Name: "ItemIdxs", Type: "[]*Idxs", Doc: "allocation image value indexes to image indexes", Directives: gti.Directives{}}},
		{"XSizes", &gti.Field{Name: "XSizes", Type: "[]int", Doc: "sorted list of all unique sizes", Directives: gti.Directives{}}},
		{"YSizes", &gti.Field{Name: "YSizes", Type: "[]int", Doc: "sorted list of all unique sizes", Directives: gti.Directives{}}},
		{"GpNs", &gti.Field{Name: "GpNs", Type: "image.Point", Doc: "number of items in each dimension group (X, Y)", Directives: gti.Directives{}}},
		{"XGpIdxs", &gti.Field{Name: "XGpIdxs", Type: "[]int", Doc: "list of x group indexes", Directives: gti.Directives{}}},
		{"YGpIdxs", &gti.Field{Name: "YGpIdxs", Type: "[]int", Doc: "list of y group indexes", Directives: gti.Directives{}}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})