package main

import (
	"fmt"
	"strconv"
)

type Sea interface {
	Say() string
}

type EarthSea struct {
	level uint
}

func (e EarthSea) Say() string {
	return "EarthSea" + strconv.Itoa(int(e.level))
}

type MarsSea struct{}

func (m MarsSea) Say() string {
	return "MarsSea"
}

type Plains interface {
	Say() string
}

type EarthPlains struct{}

func (ep EarthPlains) Say() string {
	return "earthPlains"
}

type MarsPlains struct{}

func (mp MarsPlains) Say() string {
	return "marsPlains"
}

type TerrainFactory struct {
	sea    Sea
	plains Plains
}

func NewFactory(sea Sea, plains Plains) *TerrainFactory {
	return &TerrainFactory{sea, plains}
}

func (t TerrainFactory) getSea() Sea {
	var sea Sea
	sea = t.sea
	return sea
}

func (t TerrainFactory) getPlains() Plains {
	return t.plains
}

func main() {
	factory := NewFactory(EarthSea{2}, MarsPlains{})
	fmt.Println(factory.getSea().Say())
	fmt.Println(factory.getPlains().Say())
}
