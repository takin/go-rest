// Package nst Main Controller untuk memudahkan proses
// Pembuatan response pada API Controller yang
// mengimplementasikan Beego Framework
// Created By: Syamsul Muttaqin @2019
package nst

import (
	"regexp"

	"github.com/astaxie/beego"
)

// Response standar response model
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Controller main controller structre
type Controller struct {
	beego.Controller
	Res Response
	Err error
}

// Serve main method for response
func (c *Controller) Serve() {

	if c.Err != nil {
		if c.Res.Status == 0 {
			c.Res.Status = 403
		}

		// Set Http response code agar header status code sama dengan body
		c.Ctx.Output.SetStatus(c.Res.Status)

		// Beberapa error message masih mengembalikan raw format
		// seperti beberapa error message dari ORM Beego masih menggunakan
		// format <QuerySetter> ... sehingga ini harus dibuang
		re := regexp.MustCompile("^[<(a-zA-Z0-9)>]+\\s*")
		c.Res.Message = re.ReplaceAllString(c.Err.Error(), "")
	} else {
		c.Res.Message = "Success"
		c.Res.Status = 200
	}

	c.Data["json"] = c.Res
	c.ServeJSON()
}
