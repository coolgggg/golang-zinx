package main

import (
	"fmt"
	"golang-zinx/znet"
	"io"
	"net"
	"time"
)

func main() {

	fmt.Println("client0 start...")

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit")
		return
	}

	for {
		//发送封包msg msg id 0
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMesPackage(0, []byte("zinx client0 test message")))
		if err != nil {
			fmt.Println("pack error", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error", err)
			return
		}

		//服务器给我们回复一个msg数据，msg id 1
		//先读取流中的head部分，得到id 和data len
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error", err)
			break
		}
		//将二进制的head拆包到msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack err", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			//msg里是有数据的，再根据data len 进行第二次读取，将data读出来
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data err", err)
				return
			}

			fmt.Println("-->recv server msg: id=", msg.Id,
				", len:", msg.DataLen,
				", data:", string(msg.Data),
			)
		}
		/*
			_, err := conn.Write([]byte("zinx v0.2"))
			if err != nil {
				fmt.Println("write conn err", err)
				return
			}

			buf := make([]byte, 512)
			cnt, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read buf error")
				return
			}

			fmt.Printf(" server call back : %s, cnt=%d\n", buf, cnt)
		*/

		time.Sleep(1 * time.Second)

	}

}
