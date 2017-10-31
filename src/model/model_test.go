package model

import (
	"testing"
	"fmt"
)

func TestQuery(t *testing.T) {
	rst, err := Query()
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range rst {
		fmt.Println(p)
	}
}

func TestQueryOne(t *testing.T) {
	person, err := QueryOne(1)
	if err != nil {
		t.Fatal(err)
	}
	
	fmt.Println(person)
	fmt.Println(&person)
}

func TestAdd(t *testing.T) {
	p := &Person{
		Name:   "lxl",
		Age:    28,
		Remark: "小213",
	}
	err := Add(p)
	if err != nil {
		t.Fatal(err)
	}
	
	fmt.Println(p, "插入成功!")
}

func TestDelete(t *testing.T) {
	id := 4
	person, err := QueryOne(id)
	if err != nil {
		t.Fatal(err)
	}
	
	err = Delete(id)
	if err != nil {
		t.Fatal(err)
	}
	
	fmt.Println(person, "删除成功!")
}

func TestUpdate(t *testing.T) {
	err:=Update(3,"小浪蹄子")
	if err != nil {
		t.Fatal(err)
	}
	
	person, err := QueryOne(3)
	if err != nil {
		t.Fatal(err)
	}
	
	fmt.Println(person,"修改成功！")
}

func BenchmarkQueryOne(b *testing.B) {

}

