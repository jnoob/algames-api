package controllers

import (
	"github.com/astaxie/beego"
)

// MorrisController is the controller for morris games
type MorrisController struct {
	beego.Controller
}

// Get
func (o *MorrisController) Get() {
	AesEncrypt()
}
