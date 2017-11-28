package entities

type userInfoDao DaoSource

// Save insert new user information to database
func (dao *userInfoDao) Save(u *UserInfo) error {
	_, err := dao.Insert(u)
	checkErr(err)
	return err
}

var userInfoQueryByID = "uid = ?"

// FindAll find information of all users
func (dao *userInfoDao) FindAll() []UserInfo {
	ulist := make([]UserInfo, 0, 0)
	err := dao.Find(&ulist)
	checkErr(err)

	return ulist
}

// FindByID find information of a user by his id
func (dao *userInfoDao) FindByID(id int) *UserInfo {
	u := UserInfo{}
	_, err := dao.Where(userInfoQueryByID, id).Get(&u)
	checkErr(err)

	return &u
}
