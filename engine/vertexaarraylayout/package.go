package vertexaarraylayout

type VertexArrayElement struct {
	V_Type uint32;
	V_Count uint32;
	V_Normalised bool
}

type VertexArrayLayout struct {}

func (e *VertexArrayElement) New() {
	e.V_Count = e.V_Count + 1
}


