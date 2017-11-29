package entities

import "service-computing/cloudgo-data/cloudgo-data-template/mysqlt"

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := userInfoDao{mysqlt.NewSQLTemplate(tx)}
	err = dao.Save(u)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	dao := userInfoDao{mysqlt.NewSQLTemplate(mydb)}
	ulist, err := dao.FindAll()
	checkErr(err)
	return ulist
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	dao := userInfoDao{mysqlt.NewSQLTemplate(mydb)}
	u, err := dao.FindByID(id)
	checkErr(err)
	return u
}

// Count .
func (*UserInfoAtomicService) Count() int {
	dao := userInfoDao{mysqlt.NewSQLTemplate(mydb)}
	c, err := dao.Count()
	checkErr(err)
	return c
}
