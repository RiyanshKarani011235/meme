package concept

type Concept struct {
	parent   *Concept
	children []*Concept
}

type ConceptTree struct {
	root     *Concept
	concepts map[string]*Concept
}

func NewConceptTree() *ConceptTree {
	c := &Concept{nil, make([]*Concept, 0)}
	return &ConceptTree{c, map[string]*Concept{"Concept": c}}
}
