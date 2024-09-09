package dao

import (
	"go-mod-work/dao/model"
)

var ArgoCD argoCD

type argoCD struct{}

// GetArgoCDList get ArgoCD instance list
func (a *argoCD) GetArgoCDList() (argoCDInstances []model.ArgoCDInstance, err error) {
	if err := db.Find(&argoCDInstances).Error; err != nil {
		return nil, err
	}
	return
}
