package main

import (
	"encoding/json"
	"fmt"
	"gorm-gen/gen"
	"gorm-gen/mylog"
	"gorm-gen/utils"
	"io"
	"net/http"
	"time"
)

func main() {

	utils.InitGlobalConfig()
	gen.InitMetaQueryers()
	initHttpServer(utils.Global.Server.Port)
}

func initHttpServer(port int) {
	http.HandleFunc("/gen", genFile)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Start server :%d", port)
	err := s.ListenAndServe()
	fmt.Printf("Start server err:%v", err)

}

func genFile(w http.ResponseWriter, req *http.Request) {

	b, err := io.ReadAll(req.Body)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	t := &gen.TableTemplate{}
	err = json.Unmarshal(b, t)

	if err != nil {
		mylog.Logger.Error(err)
		io.WriteString(w, err.Error())
		return
	}

	err = utils.GetFirestErr(t)
	if err != nil {
		mylog.Logger.Error(err)
		io.WriteString(w, err.Error())
		return
	}

	mq, err := gen.GetMetaQueryer(t.DBCode)

	if err != nil {
		mylog.Logger.Error(err)
		io.WriteString(w, err.Error())
		return
	}
	err = gen.GenerateFile(t, mq)
	if err != nil {
		mylog.Logger.Error(err)
		io.WriteString(w, err.Error())
		return
	}
	mylog.Logger.Infof("Generate Code OK , table_name:%s", t.TableName)

	io.WriteString(w, "Generate Code OK")
}
