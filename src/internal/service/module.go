package service

import (
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/util"
)

type ModuleService interface {
	Install(req entity.Module) (*helper.ResponseModule, error)
}

type moduleservice struct {
}

func NewModuleService() ModuleService {
	return &moduleservice{}
}

func (mod *moduleservice) Install(req entity.Module) (*helper.ResponseModule, error) {
	response, err := util.InstallDepedency(req)
	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	return response, nil
}
