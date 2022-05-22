package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/masayoshi4649/GamePartySys/src/routing/routeauth"
	"github.com/masayoshi4649/GamePartySys/src/routing/routebooking"
	"github.com/masayoshi4649/GamePartySys/src/routing/routeroot"
)

// Initialize ... "common"
func Initialize(r *gin.Engine) {
	r.StaticFile("/favicon.ico", "view/favicon.ico")
	r.LoadHTMLGlob("view/**/*.html")
}

// Common ... https://example.com/common*
func Common(r *gin.Engine) {
	d := "/common"

	// 静的配置
	r.Static(d+"/css", "view"+d+"/css")
	r.Static(d+"/js", "view"+d+"/js")

	// ルーティング
	gr := r.Group(d)
	{
		gr.GET("/", routebooking.Index)
	}
}

// Root ... https://example.com/*
func Root(r *gin.Engine) {
	d := "/"

	// 静的配置
	r.Static(d+"/css", "view"+"/root"+"/css")
	r.Static(d+"/js", "view"+"/root"+"/js")

	// ルーティング
	gr := r.Group(d)
	{
		gr.GET("/", routeroot.Index)

	}
}

// Booking ... https://example.com/booking/*
func Booking(r *gin.Engine) {
	d := "/booking"

	// 静的配置
	r.Static(d+"/css", "view"+d+"/css")
	r.Static(d+"/js", "view"+d+"/js")

	// ルーティング
	gr := r.Group(d)
	{
		gr.GET("/", routebooking.Index)

		// https://example.com/booking/api/*
		api := gr.Group("/api")
		api.POST("/search", routebooking.APISearch)
	}
}

// Auth ... https://example.com/auth/*
func Auth(r *gin.Engine) {
	d := "/auth"

	// 静的配置
	r.Static(d+"/css", "view"+d+"/css")
	r.Static(d+"/js", "view"+d+"/js")

	// ルーティング
	gr := r.Group(d)
	{
		gr.GET("/login", routeauth.Login)
		gr.GET("/oauth", routeauth.Oauth)
		gr.GET("/logout", routeauth.Logout)
	}
}
