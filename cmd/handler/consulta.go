package handler

import (
	"errors"
	"strconv"
	"strings"
	"v0/internal/consulta"
	"v0/internal/dentista"
	"v0/internal/domain"
	"v0/internal/paciente"
	"v0/pkg/web"

	"github.com/gin-gonic/gin"
)

type consultaHandler struct {
	s                consulta.Service
	paciente_service paciente.Service
	dentista_service dentista.Service
}

func NewConsultaHandler(s consulta.Service, p_s paciente.Service, d_s dentista.Service) *consultaHandler {
	return &consultaHandler{
		s:                s,
		paciente_service: p_s,
		dentista_service: d_s,
	}
}

func (h *consultaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		consulta, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("consulta not found"))
			return
		}
		web.Success(c, 200, consulta)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateConsultaEmptys(consulta *domain.Consulta) (bool, error) {
	switch {
	case consulta.Data == "" || consulta.Descricao == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// validateExpiration valida que la fecha de expiracion sea valida
func validateExpiration(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

func (h *consultaHandler) Post() gin.HandlerFunc {
	type Request struct {
		Descricao  string `json:"descricao" binding:"required"`
		Data       string `json:"data" binding:"required"`
		PacienteId int    `json:"pacienteId" binding:"required"`
		DentistaId int    `json:"dentistaId" binding:"required"`
	}
	return func(c *gin.Context) {
		var r Request

		err := c.ShouldBindJSON(&r)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		valid, err := validateExpiration(r.Data)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		paciente, err := h.paciente_service.GetByID(r.PacienteId)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		dentista, err := h.dentista_service.GetByID(r.DentistaId)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}

		consulta := domain.Consulta{
			Data:      r.Data,
			Descricao: r.Descricao,
			Paciente:  paciente,
			Dentista:  dentista,
		}

		valid, err = validateConsultaEmptys(&consulta)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		p, err := h.s.Create(consulta)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *consultaHandler) Delete() gin.HandlerFunc {
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

func (h *consultaHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("consulta not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var consulta domain.Consulta
		err = c.ShouldBindJSON(&consulta)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateConsultaEmptys(&consulta)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateExpiration(consulta.Data)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, consulta)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *consultaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Descricao  string `json:"descricao omitempty" `
		Data       string `json:"data omitempty" `
		PacienteId int    `json:"pacienteId omitempty" `
		DentistaId int    `json:"dentistaId omitempty" `
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
			web.Failure(c, 404, errors.New("consulta not found"))
			return
		}

		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		update := domain.Consulta{
			Data:      r.Data,
			Descricao: r.Descricao,
		}

		if update.Data != "" {
			valid, err := validateExpiration(update.Data)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
