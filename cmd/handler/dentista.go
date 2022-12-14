package handler

import (
	"errors"
	"strconv"
	"v0/internal/dentista"
	"v0/internal/domain"
	"v0/pkg/web"

	"github.com/gin-gonic/gin"
)

type dentistaHandler struct {
	s dentista.Service
}

func NewDentistaHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{s}
}

func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentista, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, dentista)
	}
}

func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentista domain.Dentista

		err := c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(dentista)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *dentistaHandler) Delete() gin.HandlerFunc {
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

func (h *dentistaHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentista not found"))
			return
		}

		var dentista domain.Dentista
		err = c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, dentista)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

func (h *dentistaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Nome      string `json:"nome,omitempty"`
		Sobrenome string `json:"sobrenome,omitempty"`
		Matricula uint   `json:"matricula,omitempty"`
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
			web.Failure(c, 404, errors.New("dentista not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentista{
			Nome:      r.Nome,
			Sobrenome: r.Sobrenome,
			Matricula: r.Matricula,
		}

		d, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

func validateEmptys(dentista *domain.Dentista) (bool, error) {
	switch {
	case dentista.Nome == "" || dentista.Sobrenome == "":
		return false, errors.New("fields can't be empty")
	case dentista.Matricula <= 0:
		if dentista.Matricula <= 0 {
			return false, errors.New("quantity must be greater than 0")
		}
	}
	return true, nil
}
