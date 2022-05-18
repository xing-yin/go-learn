package main

import (
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/17
 * @Desc:
 */

type Animal1 struct {
	AnimalID int64     // 默认列名`animal_id`
	Birthday time.Time // 默认列名 `birthday`
	Age      int64     // 默认列名`age`
}

type Animal2 struct {
	AnimalID int64     `gorm:"column:animalID:primarykey"` // 将列名设为 `animalID` 并指定主键 id
	Birthday time.Time `gorm:"column:birthday"`            // 将列名设为 `birthday`
	Age      int64     `gorm:"column:age"`                 // 将列名设为 `age`
}

// Animal3 【推荐】方式
type Animal3 struct {
	gorm.Model
	AnimalID int64     `gorm:"column:animalID"` // 将列名设为 `animalID`
	Birthday time.Time `gorm:"column:birthday"` // 将列名设为 `birthday`
	Age      int64     `gorm:"column:age"`      // 将列名设为 `age`
}

func main() {

}
