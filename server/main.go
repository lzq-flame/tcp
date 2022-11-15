package main

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/11/15 13:25
 **/

var matchString = "^[1-9]\\d*[\\+\\]^[1-9]\\d*"

var re = regexp.MustCompile(matchString)

func process(conn net.Conn) {
	// 函数执行完之后关闭连接
	defer conn.Close()
	// 输出主函数传递的conn可以发现属于*TCPConn类型, *TCPConn类型那么就可以调用*TCPConn相关类型的方法, 其中可以调用read()方法读取tcp连接中的数据
	fmt.Printf("服务端: %T\n", conn)
	for {
		var buf [128]byte
		// 将tcp连接读取到的数据读取到byte数组中, 返回读取到的byte的数目
		n, err := conn.Read(buf[:])
		if err != nil {
			// 从客户端读取数据的过程中发生错误
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("服务端收到客户端发来的数据：", recvStr)
		allString := re.FindAllString(recvStr, 1)
		if len(allString) == 0 {
			fmt.Println("0")
		} else {
			tmp := strings.Split(allString[0], "+")
			a, _ := strconv.Atoi(tmp[0])
			b, _ := strconv.Atoi(tmp[1])
			fmt.Printf("%d+%d=%d\n", a, b, a+b)
		}
		// 由于是tcp连接所以双方都可以发送数据, 下面接收服务端发送的数据这样客户端也可以收到对应的数据
		//inputReader := bufio.NewReader(os.Stdin)
		//s, _ := inputReader.ReadString('\n')
		//t := strings.Trim(s, "\r\n")
		//// 向当前建立的tcp连接发送数据, 客户端就可以收到服务端发送的数据
		//conn.Write([]byte(t))
	}
}

func main() {
	// 监听当前的tcp连接
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	fmt.Printf("服务端: %T=====\n", listen)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		fmt.Println("当前建立了tcp连接")
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		// 对于每一个建立的tcp连接使用go关键字开启一个goroutine处理
		go process(conn)
	}

}
