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
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Description   string      `json:"description"`
	Count         int         `json:"count,omitempty"`
	Offset        int64       `json:"offset,omitempty"`
	Href          string      `json:"href"`
	Payload       interface{} `json:"payload"`
}

// Controller main controller structre
type Controller struct {
	beego.Controller
	Res Response
	Err error
}

// Serve main method for response
func (c *Controller) Serve() {

	c.Res.Href = c.Ctx.Input.URI()

	if c.Err != nil {
		if c.Res.StatusCode == 0 {
			c.Res.StatusCode = 403
		}

		// Set status message menjadi error ketika c.Err bukan sama dengan nil
		// agar frontend lebih mudah consume dan identifikasi error
		// karena hanya ada 2 pesan => ["Success","Error"]
		// dan pesan error saya masukkan ke description dibawah
		c.Res.StatusMessage = "Error"

		// Set Http response code agar header status code sama dengan body
		c.Ctx.Output.SetStatus(c.Res.StatusCode)

		// Beberapa error message masih mengembalikan raw format
		// seperti beberapa error message dari ORM Beego masih menggunakan
		// format <QuerySetter> ... sehingga ini harus dibuang
		re := regexp.MustCompile("<[a-zA-Z0-9]+>\\s*")

		// Saya menggunakan description untuk menampilkan pesan error
		c.Res.Description = re.ReplaceAllString(c.Err.Error(), "")

	} else {
		c.Res.StatusMessage = "Success"
		c.Res.StatusCode = 200
	}

	c.Data["json"] = c.Res
	c.ServeJSON()
}
