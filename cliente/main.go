package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	// "strings"
)

var date_sep = "/"

func apply_time(thetime string) {
	cmd := exec.Command("cmd.exe", "/c", "time")
	cmd_stdin, err := cmd.StdinPipe()
	if err != nil {
		println("Fallo al ejecutar el comando stdinpipe")
		return
	}
	err = cmd.Start()
	if err != nil {
		println("Fallo al ejecutar el comando time")
		return
	}

	fmt.Fprintf(cmd_stdin, "%s\n", thetime)
	cmd_stdin.Close()
	err = cmd.Wait()
	if err != nil {
		println("Fallo en: ", err.Error())
	}
}

func apply_date(thedate string) {
	cmd := exec.Command("cmd.exe", "/c", "date")
	cmd_stdin, err := cmd.StdinPipe()
	if err != nil {
		println("Fallo al ejecutar el comando stdinpipe")
		return
	}
	err = cmd.Start()
	if err != nil {
		println("Fallo al ejecutar el comando date")
		return
	}

	fmt.Fprintf(cmd_stdin, "%s\n", thedate)
	cmd_stdin.Close()
	err = cmd.Wait()
	if err != nil {
		println("Fallo en: ", err.Error())
	}
}

func main() {
	var buff = make([]byte, 3333)
	cc, err := net.Dial("tcp4", os.Args[1])
	if err != nil {
		println("Fallo al conectar")
		return
	}
	n, err := cc.Write([]byte("GET_TIME"))
	if err != nil {
		println("Fallo al enviar datos")
		return
	}

	n, err = cc.Read(buff)
	if err != nil {
		println("Fallo al recibir datos")
		return
	}
	var s string = string(buff[:n])
	s = strings.ReplaceAll(s, "\n", " ")
	values := strings.Split(s, " ")
	println(strings.Join(values, "|"))
	thetime := values[0]
	thedate := values[2]
	date_elements := strings.Split(thedate, date_sep)
	year := date_elements[2]
	month := date_elements[1]
	day := date_elements[0]
	//
	apply_time(thetime)
	println("La Hora actualizada es: ", thetime)
	apply_date(day + date_sep + month + date_sep + year)
	println("La Fecha actualizada es: ", thedate)

}
