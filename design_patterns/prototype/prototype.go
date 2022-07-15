/**
 * 原型模式
 */
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
	// 修改earthSea 的 level 属性
	var earthSea EarthSea
	earthSea.level = 10
	t.sea = earthSea
	return t.sea
}

func (t TerrainFactory) getPlains() Plains {
	return t.plains
}

func main() {
	factory := NewFactory(EarthSea{}, MarsPlains{})
	fmt.Println(factory.getSea().Say())
	fmt.Println(factory.getPlains().Say())
}
