package database

type UserModel struct {
	ID        uint
	Telephone string `gorm:"unique"`
	Password  string
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (*UserModel) Get(id uint) (res UserModel, err error) {
	err = database.Model(&UserModel{}).First(&res, id).Error
	return
}
func (*UserModel) Create(item *UserModel) error {
	return database.Model(&UserModel{}).Create(item).Error
}

func (*UserModel) Delete(id uint) error {
	return database.Delete(&UserModel{}, id).Error
}

func (*UserModel) Update(item *UserModel) error {
	return database.Save(item).Error
}

func (*UserModel) GetByTelephone(telephone string) (res UserModel, err error) {
	err = database.Model(&UserModel{}).Where("telephone = ?", telephone).First(&res).Error
	return
}

func (UserModel) TableName() string {
	return "user"
}
