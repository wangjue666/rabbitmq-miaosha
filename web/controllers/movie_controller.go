package controllers

import (
	"rabbitmq-miaosha/repositories"
	"rabbitmq-miaosha/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
	Ctx iris.Context
}

func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManger(movieRepository)
	MovieResult := movieService.ShowMovieName()

	return mvc.View{
		Name: "movie/index.html",
		Data: MovieResult,
	}
}

func (c *MovieController) GetHello() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManger(movieRepository)
	MovieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: MovieResult,
	}
}
