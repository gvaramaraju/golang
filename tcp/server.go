package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	defer lis.Close()
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Waiting for connection")
		conn, err := lis.Accept()
		fmt.Println(conn.RemoteAddr().String(), " Connected")
		if err != nil {
			panic(err)
		}
		// go handleConnection(conn)
		go handleHTTPrequest(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		dataBytes, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		data := strings.TrimSpace(string(dataBytes))

		fmt.Println(data)
		fmt.Println(len(data))
		if data == "exit" {
			conn.Close()
			fmt.Println(conn.RemoteAddr().String(), " Disconnected")
			break
		}

	}
}

func handleHTTPrequest(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	text := scanner.Text()
	httpSchema := strings.Split(text, " ")
	switch httpSchema[0] {
	case "GET":
		fmt.Println("Received Get request")
		httpHeaders := getRequestHeaders(*scanner)
		fmt.Println(httpHeaders)
	}
	// scanner.Split(bufio.ScanLines)
	// for scanner.Scan() != false {
	// 	fmt.Println("Scanning request")
	// 	text := scanner.Text()
	// 	fmt.Println(text)
	// 	if len(text) == 0 {
	// 		fmt.Println("Inside length = 0")
	// 		err := scanner.Err()
	// 		fmt.Println("Error :: ", err)
	// 		scanner
	// 		if err == io.EOF {
	// 			fmt.Println("Error EOF", err)
	// 			break
	// 		}
	// 		continue
	// 	}
	// }
	// var buf bytes.Buffer
	// io.Copy(&buf, conn)
	// data := buf.String()
	// fmt.Println(data)
	// scanner := bufio.NewScanner(conn)
	// for {
	// 	eof := scanner.Scan()
	// 	fmt.Println("EOF ::", eof)
	// 	str := scanner.Text()
	// 	if eof == false {
	// 		break
	// 	}
	// 	if str == "" {
	// 		continue
	// 	}
	// 	fmt.Println(str)
	// }
	// for scanner.Scan() != false {
	// 	str := scanner.Text()
	// 	if len(str) == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(str)
	// }

}

func getRequestHeaders(scanner bufio.Scanner) *httpHeaders {
	headers := make(map[string]string)
	httpHeaders := &httpHeaders{
		&headers,
	}
	for scanner.Scan() {
		header := scanner.Text()
		fmt.Println("Header :: ", header)
		if header == "" {
			break
		}
		headerSplit := strings.Split(header, ":")
		(*httpHeaders.headers)[headerSplit[0]] = headerSplit[1]

	}
	return httpHeaders
}

type httpHeaders struct {
	headers *map[string]string
}
