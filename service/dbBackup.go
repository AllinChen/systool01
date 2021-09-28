package service

import (
	"github.com/AllinChen/systool01/config"
	"github.com/AllinChen/systool01/controller"
)

func DbBackup() (BKpath string, err error) {
	err, BKpath = controller.BackupMySqlDb(config.GlobalConf.Mysql.Host, config.GlobalConf.Mysql.Port, config.GlobalConf.Mysql.User, config.GlobalConf.Mysql.Pass, config.GlobalConf.Mysql.DatabaseName, config.GlobalConf.Mysql.TableName, config.GlobalConf.Mysql.SqlPath)
	return
}
