package main

import (
	"v0/cmd/handler"
	"v0/internal/consulta"
	"v0/internal/dentista"
	"v0/internal/paciente"
	"v0/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {

	db := store.InitDB()
	defer db.Close()

	dentistaRepo := dentista.NewRepository(db)
	dentistaService := dentista.NewService(dentistaRepo)
	dentistaHandler := handler.NewDentistaHandler(dentistaService)

	pacienteRepo := paciente.NewRepository(db)
	pacienteService := paciente.NewService(pacienteRepo)
	pacienteHandler := handler.NewPacienteHandler(pacienteService)

	consultaRepo := consulta.NewRepository(db)
	consultaService := consulta.NewService(consultaRepo)
	consultaHandler := handler.NewConsultaHandler(consultaService, pacienteService, dentistaService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentistas := r.Group("/dentista")
	{
		dentistas.GET(":id", dentistaHandler.GetByID())
		dentistas.POST("", dentistaHandler.Post())
		dentistas.DELETE(":id", dentistaHandler.Delete())
		dentistas.PUT(":id", dentistaHandler.Put())
		dentistas.PATCH(":id", dentistaHandler.Patch())
	}
	pacientes := r.Group("/paciente")
	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":id", pacienteHandler.Delete())
		pacientes.PUT(":id", pacienteHandler.Put())
		pacientes.PATCH(":id", pacienteHandler.Patch())
	}
	consultas := r.Group("/consulta")
	{
		consultas.GET(":id", consultaHandler.GetByID())
		consultas.POST("", consultaHandler.Post())
		consultas.DELETE(":id", consultaHandler.Delete())
		consultas.PUT(":id", consultaHandler.Put())
		consultas.PATCH(":id", consultaHandler.Patch())
	}

	r.Run(":8080")

}
