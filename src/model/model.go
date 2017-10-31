package model

import (
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mssql", "server=127.0.0.1;port=1433;user id=sa;password=123456;encrypt=disable;database=db_xiye")
	if err != nil {
		log.Printf("连接数据库错误：%s\n", err)
	}
	
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
}

type Person struct {
	Id     int
	Name   string
	Age    int
	Remark string
}

func Query() ([]*Person, error) {
	rst := make([]*Person, 0)
	
	rows, err := db.Query("select  Id,name,age,remark from [dbo].[tb_test] ")
	if err != nil {
		log.Println("查询失败：", err)
		return rst, err
	}
	defer rows.Close()
	
	for rows.Next() {
		p := &Person{}
		err = rows.Scan(&p.Id, &p.Name, &p.Age, &p.Remark)
		if err != nil {
			return nil, err
		}
		
		rst = append(rst, p)
	}
	
	return rst, nil
}

func QueryOne(id int) (*Person, error) {
	person := &Person{}
	rows, err := db.Query("select  Id,name,age,remark from [dbo].[tb_test] where ID = ? ", id)
	if err != nil {
		log.Println("查询失败：", err)
		return person, err
	}
	defer rows.Close()
	
	for rows.Next() {
		err = rows.Scan(&person.Id, &person.Name, &person.Age, &person.Remark)
		if err != nil {
			return nil, err
		}
	}
	
	return person, nil
}

func Add(p *Person) (error) {
	_, err := db.Exec(`insert into tb_test(name,age,remark)values(?,?,?)`, p.Name, p.Age, p.Remark)
	if err != nil {
		log.Println("插入数据失败：", err)
		return err
	}
	
	return nil
}

func Delete(id int) ( error) {
	_, err := db.Exec("delete from [dbo].[tb_test] where ID = ? ", id)
	if err != nil {
		log.Println("删除失败：", err)
		return  err
	}
	
	return  nil
}

func Update(id int,remark string) (error) {
	_, err := db.Exec(`update  tb_test set remark = ? where id = ?`, remark,id)
	if err != nil {
		log.Println("修改数据失败：", err)
		return err
	}
	
	return nil
}
