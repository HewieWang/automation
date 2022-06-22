package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/PuerkitoBio/goquery"
    "log"
    "bytes"
    "encoding/json"
)
type Server struct {
    State  string `json:"State,string"`
    County string `json:"County,string"`
    Place  string `json:"Place,string"`
    Code   string `json:"Code,string"`
}
type Serverslice struct {
    Servers []Server `json:"zipcode"`
}

func main() {

    db, _ := sql.Open("mysql", "root:root@(127.0.0.1)/caiji")
    defer db.Close()
    err := db.Ping()
    if err != nil {
        log.Println("数据库连接失败")
        return
    } else {
        log.Println("数据库连接成功")
    }

    //多行查询
    rows, _ := db.Query("select data from uszipcode2")
    var data string
    var res Serverslice
    for rows.Next() { //循环显示所有的数据
        rows.Scan(&data)
        doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(data))
        if err != nil {
          log.Fatal(err)
        }

        doc.Find(".units.noletters .unit").Each(func(i int, s *goquery.Selection) {
          place := s.Find(".place").Text()
          code:=s.Find(".code").Text()
          // fmt.Printf("%d: %s : %s\n", i, place,code)
          res.Servers = append(res.Servers, Server{Place: place, Code:code})
        })
    }

    b, err := json.Marshal(res)
    if err != nil {
        fmt.Println("JSON ERR:", err)
    }
    fmt.Println(string(b))
}
