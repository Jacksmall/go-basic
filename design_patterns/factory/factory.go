package main

import "fmt"

type ApptEncoder interface {
	encode() string
}

type BloggerApptEncoder struct{}

func (b BloggerApptEncoder) encode() string {
	return "这是BloggerApptEncoder 结构体 ApptEncoder 接口的实现"
}

type MegaApptEncoder struct{}

func (m MegaApptEncoder) encode() string {
	return "这是MegaApptEncoder 结构体 ApptEncoder 接口的实现"
}

type CommonsManager interface {
	GetHeader() string
	GetApptEncoder() ApptEncoder
	GetBottom() string
}

type TCommonsManager struct {
	cm CommonsManager
}

// NewTCommonsManager creates a new TCommonsManager 新建工厂实例
func NewTCommonsManager(t string) TCommonsManager {
	var cm CommonsManager
	switch t {
	case "blogger":
		cm = BloggerCommonsManager{}
	case "mega":
		cm = MegaCommonsManager{}
	}
	return TCommonsManager{cm: cm}
}

// BloggerCommonsManager 结构体去实现CommonsManager
type BloggerCommonsManager struct{}

func (bc BloggerCommonsManager) GetHeader() string {
	return "这是blogger的commons getHeader的实现"
}

func (bc BloggerCommonsManager) GetApptEncoder() ApptEncoder {
	return BloggerApptEncoder{}
}

func (bc BloggerCommonsManager) GetBottom() string {
	return "这是blogger的commons getBottom的实现"
}

// MegaCommonsManager 结构体去实现CommonsManager
type MegaCommonsManager struct{}

func (mc MegaCommonsManager) GetHeader() string {
	return "这是mega的commons getHeader的实现"
}

func (mc MegaCommonsManager) GetApptEncoder() ApptEncoder {
	return MegaApptEncoder{}
}

func (mc MegaCommonsManager) GetBottom() string {
	return "这是mega的commons getBottom的实现"
}

func main() {
	tc := NewTCommonsManager("mega")
	fmt.Println(tc.cm.GetHeader())
	fmt.Println(tc.cm.GetApptEncoder().encode())
	fmt.Println(tc.cm.GetBottom())
}
