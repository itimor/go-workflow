package models

import (
	"fmt"
	"go-workflow/backend/models/sys"
	"go-workflow/backend/models/workflow"
	"go-workflow/backend/models/workflowform"
	"time"

	"iris-ticket/backend/config"
	"iris-ticket/backend/models/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pelletier/go-toml"
)

func InitDB() {
	var gdb *gorm.DB
	var err error
	GetEnv := "test"
	// if GetEnv() == "test" {
	// } else {}

	dbconfig := config.Conf.Get(GetEnv + ".database").(*toml.Tree)
	driver := dbconfig.Get("driver").(string)
	connect := dbconfig.Get("connect").(string)
	gdb, err = gorm.Open(driver, connect)
	if err != nil {
		panic(err)
	}
	gdb.SingularTable(true)
	// gormconfig := config.Conf.Get("gorm").(*toml.Tree)
	gdb.DB().SetMaxIdleConns(150)
	gdb.DB().SetMaxOpenConns(50)
	gdb.DB().SetConnMaxLifetime(time.Duration(7200) * time.Second)
	db.DB = gdb
}

func Migration() {
	println("初始化数据库")
	// sys
	fmt.Println("db Menu 初始化：", db.DB.AutoMigrate(new(sys.Menu)).Error)
	fmt.Println("db User 初始化：", db.DB.AutoMigrate(new(sys.User)).Error)
	fmt.Println("db RoleMenu 初始化：", db.DB.AutoMigrate(new(sys.RoleMenu)).Error)
	fmt.Println("db Role 初始化：", db.DB.AutoMigrate(new(sys.Role)).Error)
	fmt.Println("db UserRole 初始化：", db.DB.AutoMigrate(new(sys.UserRole)).Error)
	fmt.Println("db OauthToken 初始化：", db.DB.AutoMigrate(new(sys.OauthToken)).Error)
	// workflow
	fmt.Println("db CaseType 初始化：", db.DB.AutoMigrate(new(workflow.CaseType)).Error)
	fmt.Println("db CaseTypeStep 初始化：", db.DB.AutoMigrate(new(workflow.CaseTypeStep)).Error)
	fmt.Println("db Case 初始化：", db.DB.AutoMigrate(new(workflow.Case)).Error)
	fmt.Println("db CaseCaseType 初始化：", db.DB.AutoMigrate(new(workflow.CaseCaseType)).Error)
	fmt.Println("db CaseOpera 初始化：", db.DB.AutoMigrate(new(workflow.CaseOpera)).Error)
	fmt.Println("db CaseStep 初始化：", db.DB.AutoMigrate(new(workflow.CaseStep)).Error)
	fmt.Println("db CaseType 初始化：", db.DB.AutoMigrate(new(workflow.CaseType)).Error)
	// workflowform
	fmt.Println("db Deploy 初始化：", db.DB.AutoMigrate(new(workflowform.Deploy)).Error)
}
