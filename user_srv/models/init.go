package models

import (
	"fmt"
	"time"
	"user_srv/db"

	"github.com/jinzhu/gorm"
)

func init() {
	RegisterCallbacks()
	AutoMigrateModels()
}

// RegisterCallbacks 注册自定义的gorm回调钩子
func RegisterCallbacks() {
	db.Client().Callback().Delete().Replace("gorm:delete", deleteCallback)
}

// AutoMigrateModels 自动生成模型对应的表结构
func AutoMigrateModels() {
	// Migrate the schema
	db.Client().AutoMigrate(&User{})
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt")
		deletedField, hasDeletedField := scope.FieldByName("Deleted")
		idField, _ := scope.FieldByName("ID")

		if !scope.Search.Unscoped && hasDeletedAtField && !hasDeletedField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else if !scope.Search.Unscoped && hasDeletedAtField && hasDeletedField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v, %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedField.DBName),
				scope.AddToVars(idField.Field.Interface()),
				scope.Quote(deletedAtField.DBName),
				scope.AddToVars(time.Now().Format("2006-01-02 15:04:05")),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
