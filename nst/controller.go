// Package controllers Main Controller untuk memudahkan proses
// Pembuatan response pada API Controller yang
// mengimplementasikan Beego Framework
// Created By: Syamsul Muttaqin @2019
package nst

import (
	"github.com/astaxie/beego"
)

// Response standar response model
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// Controller main controller structre
type Controller struct {
	beego.Controller
	Res Response
}

// ServeSuccess method untuk mengembalikan response sukses
func (c *Controller) ServeSuccess() {
	c.Res.Message = "Success"
	c.Res.Status = 200
	c.Serve()
}

// ServeError Response formatter untuk error response
func (c *Controller) ServeError() {
	if c.Res.Status == 0 {
		c.Res.Status = 403
	}
	if c.Res.Message == "" {
		c.Res.Message = "Unknown Error Occurs"
	}
	c.Serve()
}

// Serve main method for response
func (c *Controller) Serve() {
	c.Data["json"] = c.Res
	c.ServeJSON()
}
