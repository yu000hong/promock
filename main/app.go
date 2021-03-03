package main

import (
	"database/sql"
	"encoding/json"
	"log"
)

var host2app map[string]*App

// App 应用
type App struct {
	ID				int			`json:"id"`
	Host			string		`json:"host"`
	Name			string		`json:"name"`
	Params			[]Param		`json:"params"`
	CreateTime		int			`json:"create_time"`
	UpdateTime		int			`json:"update_time"`
}

// AddApp 新增App
func AddApp(db *sql.DB, app *App) {
	if _, ok := host2app[app.Host]; ok {
		//TODO throw exceptio
		return
	}
	db.Exec("INSERT INTO `app`(`host`,`name`,`params`,`create_time`) VALUES(?,?,?,UNIX_TIMESTAMP())", 
		app.Host, app.Name, ToJSON(app.Params))
	//TODO fill ID into app
	host2app[app.Host] = app
}

// DelApp 删除App
func DelApp(db *sql.DB, id int) {
	db.Exec("DELTE FROM `app` WHERE `id`=?", id)
	for host, app := range host2app {
		if app.ID == id {
			delete(host2app, host)
		}
	}
}

// UpdateApp 更新App
func UpdateApp(db *sql.DB, app *App) {
	db.Exec("UPDATE `app` SET `name`=?,`host`=?,`params`=?,`update_time`=UNIX_TIMESTAMP() WHERE `id`=?", 
		app.Name, app.Host, ToJSON(app.Params), app.ID)
	for h, a := range host2app {
		if app.ID == a.ID {
			delete(host2app, h)
			host2app[app.Host] = app
		}
	}
}

// InitApps 初始化
func InitApps(db *sql.DB) {
	host2app = make(map[string]*App)
	rows, err := db.Query("SELECT `id`,`name`,`host`,`params`,`create_time`,`update_time` FROM `app`")
	if err != nil {
		log.Fatal("Error when init apps")
	}
	for rows.Next() {
		app := &App{}
		var params string
		err = rows.Scan(&app.ID, &app.Name, &app.Host, &params, &app.CreateTime, &app.UpdateTime)
		if err != nil {
			log.Fatal("Error when fetching rows from app")
		}
		if err = json.Unmarshal([]byte(params), &app.Params); err != nil {
			panic("error when unmarshalling params")
		}
	}
}
