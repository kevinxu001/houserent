// organization

// @description 组织机构控制器

// @link

// @license     https://github.com/kevinxu001/houserent/blob/master/LICENSE

// @authors     kevinxu

package controllers

import (
	"github.com/kevinxu001/houserent/models/rbac"
)

type OrganizationController struct {
	CommonController
}

func (this *OrganizationController) Get() {
	ou := new(rbac.OrganizationUnit)
	ou.UnitName = "百年房产"
	this.Ctx.WriteString(ou.UnitName)
}
