package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//只是负责测试拆包
func TestDataPack(t *testing.T) {
	//模拟的服务器
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err", err)
		return
	}

	//创建一个go 承载 负责从客户端处理业务
	go func() {
		//从客户端读取数据，拆包处理
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("server accept err", err)
			}

			go func(conn net.Conn) {
				//处理客户端的请求
				//-----》拆包的过程《-----
				//定义一个拆包的对象dp
				dp := NewDataPack()
				for {
					//第一次从conn读，把包的head读出来
					headData := make([]byte, dp.GetHeadLen())
					if _, err := io.ReadFull(conn, headData); err != nil {
						fmt.Println("read head error")
						return
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err", err)
						break
					}

					if msgHead.GetMsgLen() > 0 {
						//msg是有数据的，需要进行二次读取
						//第二次从conn读，根据head中的data len 再读取data内存
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())

						//根据data len的长度再次从io流中
						if _, err := io.ReadFull(conn, msg.Data); err != nil {
							fmt.Println("server unpack data err", err)
							return
						}

						//完整的一个消息已经读取完毕
						fmt.Println("----> recv msg id:", msg.Id,
							", data len:", msg.DataLen,
							", data =", string(msg.Data),
						)
					}
				}
			}(conn)
		}
	}()

	//模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err ", err)
	}

	//创建一个封包对象
	dp := NewDataPack()

	//模拟粘包过程，封装两个msg一同发送
	//封装一个msg1
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'z', 'i', 'n', 'x', '!'},
	}
	sendPack1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err", err)
		return
	}

	//封装第二个msg2
	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'n', 'i', 'h', 'a', 'o', '!', '!'},
	}
	sendPack2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 err", err)
		return
	}

	//将两个包粘在一起
	sendPack1 = append(sendPack1, sendPack2...)

	//一起发送
	conn.Write(sendPack1)

	//阻塞
	select {}

}
