// go2hal v0.2.0
// Copyright (c) 2016 Patrick Moule
// License: MIT

package hal

import "github.com/pmoule/go2hal/hal/relationtype"

// Resource is the root element of a HAL document.
// A Resource can
// - have links - AddLink(LinkRelation)
// - have CURIEs - AddCurieLinks([]*LinkObject)
// - embed other resources - AddResource(ResourceRelation)
type Resource interface {
	Data() PropertyMap
	Links() NamedMap
	EmbeddedResources() NamedMap
	AddLink(LinkRelation)
	AddResource(ResourceRelation)
	AddCurieLinks([]*LinkObject)
}

type resourceObject struct {
	data     PropertyMap `json:"-"`
	links    links `json:"_links,omitempty"`
	embedded embeddedResources `json:"_embedded,omitempty"`
}

// NewResourceObject initializes a valid Resource.
func NewResourceObject() Resource {
	return &resourceObject{data: PropertyMap{}, links: links{}, embedded: embeddedResources{}}
}

func (r *resourceObject) Data() PropertyMap {
	return r.data
}

func (r *resourceObject) Links() NamedMap {
	return r.links.ToMap()
}

func (r *resourceObject) EmbeddedResources() NamedMap {
	return r.embedded.ToMap()
}

func (r *resourceObject) ToMap() NamedMap {
	propertyMap := PropertyMap{}

	mappers := []mapper{&r.links, &r.embedded}

	for _, mapper := range mappers {
		namedMap := mapper.ToMap()

		if len(namedMap.Content) > 0 {
			propertyMap[namedMap.Name] = namedMap.Content
		}
	}

	for key, val := range r.data {
		propertyMap[key] = val
	}

	return NamedMap{Name: "root", Content: propertyMap}
}

func (r *resourceObject) AddCurieLinks(linkObjects []*LinkObject) {
	rel, _ := NewLinkRelation(relationtype.CURIES)
	rel.SetLinks(linkObjects)
	r.AddLink(rel)
}

func (r *resourceObject) AddLink(rel LinkRelation) {
	r.links[rel.Name()] = rel
}

func (r *resourceObject) AddResource(rel ResourceRelation) {
	r.embedded[rel.Name()] = rel
}