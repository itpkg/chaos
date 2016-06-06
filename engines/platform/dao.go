package platform

import "github.com/jinzhu/gorm"

type Dao struct {
	Db        *gorm.DB  `inject:""`
	Encryptor Encryptor `inject:""`
}
