package gohttp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// RunServer run server
func RunServer(engine *gin.Engine) {
	addr := GetServerAddr()
	fmt.Println("server addr : ", addr)

	// https
	if strings.ToLower(os.Getenv("ServerSSLUse")) == "true" {
		pwdPath, err := os.Getwd()
		if err != nil {
			panic(fmt.Errorf("RunServer os.Getwd error : %v", err))
		}

		certFile := filepath.Join(pwdPath, os.Getenv("ServerSSLFileCert"))
		keyFile := filepath.Join(pwdPath, os.Getenv("ServerSSLFileKey"))

		if err := RunTLS(engine, addr, certFile, keyFile); err != nil {
			panic(fmt.Errorf("RunServer RunTLS(engine, addr, certFile, keyFile) error : %v", err))
		}
	}

	// server
	if err := Run(engine, addr); err != nil {
		panic(fmt.Errorf("RunServer Run(engine, addr) error : %v", err))
	}
}

// GetServerAddr server addr
// GetServerAddr default 30216
func GetServerAddr() string {
	port := os.Getenv("ServerPort")
	if len(port) > 0 {
		return ":" + port
	}
	return ":30216"
}

// RegisterServer register server
func RegisterServer() {
	ip := os.Getenv("ServerIP")
	if len(ip) < 1 {
		ip = getLocalIP()
	}

	address := ip + GetServerAddr()

	_ = address
}

// getLocalIp local ip
// getLocalIp default 127.0.0.1
func getLocalIP() string {
	conn, err := net.Dial("tcp", "163.com:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()

	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
