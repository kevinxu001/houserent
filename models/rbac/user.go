// user

// @description 用户模型定义

// @link

// @license     https://github.com/kevinxu001/houserent/blob/master/LICENSE

// @authors     kevinxu

package rbac

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	UserName string `orm:"size(50)" form:"UserName" valid:"Required;MaxSize(50)"`
	PassWord string `orm:"size(40)" form:"PassWord" valid:"Required;MaxSize(40)"`
	Mobile   string `orm:"size:(20)" form:"Mobile" valid:"Required;Mobile"`
	Phone    string `orm:"size:(20)" form:"Phone" valid:"Phone"`
	IdCard   string `orm:"size:(20)" form:"IdCard" valid:"Required;Match(/\d{17}[\dXx]{1}/)"`
	Status   uint8  `orm:"default(0)" form:"Status" valid:"Range(0，1)"`
}

// func (o *OrganizationUnit) TableName() string {
// 	return "organization_unit"
// }

// func (o *OrganizationUnit) TableIndex() [][]string {
// 	return [][]string{
// 		[]string{"Id","UnitName"},
// 	}
// }

// func (o *OrganizationUnit) TableUnique() [][]string {
// 	return [][]string{
// 		[]string{"UnitName"}
// 	}
// }

func init() {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("tableprefix"), new(User))
}
