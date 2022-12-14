package handler

import (
	"errors"
	"strconv"
	"v0/internal/domain"
	"v0/internal/paciente"
	"v0/pkg/web"

	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{s}
}

func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentista, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("paciente not found"))
			return
		}
		web.Success(c, 200, dentista)
	}
}

func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente

		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validatePacienteEmptys(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *pacienteHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

func (h *pacienteHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("paciente not found"))
			return
		}

		var paciente domain.Paciente
		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validatePacienteEmptys(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *pacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Nome      string `json:"nome,omitempty"`
		Sobrenome string `json:"sobrenome,omitempty"`
		RG        string `json:"rg,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("paciente not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Paciente{
			Nome:      r.Nome,
			Sobrenome: r.Sobrenome,
			RG:        r.RG,
		}

		d, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

func validatePacienteEmptys(paciente *domain.Paciente) (bool, error) {
	switch {
	case paciente.Nome == "" || paciente.Sobrenome == "" || paciente.RG == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}
