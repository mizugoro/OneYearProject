package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	ID   int
	Name string
	Age  int
}

func main() {
	//DBを開く、ファイルが存在しない場合生成される
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")

	//終わったら閉じる
	defer DbConnection.Close()

	//DB作成　SQLコマンド
	cmd := `CREATE TABLE IF NOT EXISTS person(
		ID iINT,
		Name STRING,
		Age INT)`

	//実行結果は返ってこないため _ を利用する
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}

}
