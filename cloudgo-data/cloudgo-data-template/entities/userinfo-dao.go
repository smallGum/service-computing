package entities

import (
	"service-computing/cloudgo-data/cloudgo-data-template/mysqlt"
)

type userInfoDao struct {
	mysqlt.SQLTemplate
}

var userInfoInsertStmt = "INSERT userinfo SET username=?,departname=?,created=?"

// Save .
func (dao *userInfoDao) Save(u *UserInfo) error {
	return dao.Insert(userInfoInsertStmt, &u.UID, u.UserName, u.DepartName, u.CreateAt)
}

var userInfoQueryAll = "SELECT * FROM userinfo"
var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"
var userInfoCount = "SELECT count(*) FROM userinfo"

func getUserInfoMapper(ul *[]UserInfo) mysqlt.RowMapperCallback {
	return func(row mysqlt.RowScanner) error {
		u := UserInfo{}
		err := row.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
		if err != nil {
			return err
		}
		*ul = append(*ul, u)
		return nil
	}
}

func getUserInfoOnceMapper(u *UserInfo) mysqlt.RowMapperCallback {
	return func(row mysqlt.RowScanner) error {
		err := row.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
		return err
	}
}

// FindAll .
func (dao *userInfoDao) FindAll() ([]UserInfo, error) {
	ulist := make([]UserInfo, 0, 0)
	err := dao.Select(userInfoQueryAll, getUserInfoMapper(&ulist))
	return ulist, err
}

// FindByID .
func (dao *userInfoDao) FindByID(id int) (*UserInfo, error) {
	u := UserInfo{}
	err := dao.SelectOne(userInfoQueryByID, getUserInfoOnceMapper(&u), id)
	return &u, err
}

// Count .
func (dao *userInfoDao) Count() (int, error) {
	count := 0
	f := func(row mysqlt.RowScanner) error {
		err := row.Scan(&count)
		return err
	}
	err := dao.SelectOne(userInfoCount, f)
	return count, err
}
