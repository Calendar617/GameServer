package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBProcessor struct {
	db *sql.DB
}

func CreateDBProcessor() *DBProcessor {
	procssor := &DBProcessor{}
	procssor.Open()
	return procssor
}

func (dbp *DBProcessor) Open() {
	db, err := sql.Open("mysql", "root:830726@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("MySql Server Init Success!")
	dbp.db = db
}

func (dbp *DBProcessor) QueryRole(id uint64) (result bool, pos [3]float32) {
	fmt.Printf("query playerinfo id = %d\n", id)
	var name string
	var posX, posY, posZ float32 = 0.0, 0.0, 0.0
	err := dbp.db.QueryRow("select Name, PosX, PosY, PosZ from playerinfo where ID = ?", id).Scan(&name, &posX, &posY, &posZ)
	if err != nil {
		fmt.Printf("playerinfo id = %d not exist!\n", id)
		return false, [3]float32{0.0, 0.0, 0.0}
	}
	fmt.Printf("query playerinfo id = %d name = %s\n", id, name)
	return true, [3]float32{posX, posY, posZ}
}

func (dbp *DBProcessor) CreateRole(id uint64, name string, pos [3]float32) bool {
	currtime := time.Now()
	_, err := dbp.db.Exec("insert into playerinfo(ID, Name, PosX, PosY,PosZ,CreateTime) values (?,?,?,?,?,?)", id, name, pos[0], pos[1], pos[2], currtime)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("insert playerinfo id = %d success!\n", id)
	return true
}

func (dbp *DBProcessor) UpdateRole(id uint64, pos [3]float32) bool {
	_, err := dbp.db.Exec("update playerinfo set PosX = ?, PosY = ? ,PosZ =? where ID =?", pos[0], pos[1], pos[2], id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("update playerinfo id = %d success!\n", id)
	return true
}

func (dbp *DBProcessor) Close() {
	dbp.db.Close()
}
