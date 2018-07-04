/**
 * Copyright 2018 godog Author. All Rights Reserved.
 * Author: Chuck1024
 */

package main

import (
	"github.com/xuyu/logging"
	"godog/net/tcplib"
	"godog"
	"net/http"
)

var App *godog.Application

func HandlerHttpTest(w http.ResponseWriter, r *http.Request) {
	logging.Debug("connected : %s",r.RemoteAddr)
	w.Write([]byte("test success!!!"))
}

func HandlerTcpTest(req tcplib.Packet) (rsp tcplib.Packet) {
	cReq := req.(*tcplib.TcpPacket)
	rsp = tcplib.NewCustomPacketWithSeq(cReq.Cmd, []byte("1024 hello."), cReq.Seq)
	return
}

func main() {
	App = godog.NewApplication("test")
	// Http
	App.AppHttp.AddHandlerFunc("/test", HandlerHttpTest)

	// Tcp
	App.AppTcpServer.AddTcpHandler(1024, HandlerTcpTest)

	err := App.Run()
	if err != nil {
		logging.Error("Error occurs, error = %s", err.Error())
		return
	}
}

// you can use command to test service that it is in another file <serviceTest.txt>.
