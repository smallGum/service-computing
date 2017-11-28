package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save insert new user information to database
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	session := mydb.NewSession()
	err := session.Begin()
	defer session.Close()

	dao := userInfoDao{mydb}
	err = dao.Save(u)

	if err == nil {
		session.Commit()
	} else {
		session.Rollback()
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	dao := userInfoDao{mydb}
	return dao.FindAll()
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	dao := userInfoDao{mydb}
	return dao.FindByID(id)
}
