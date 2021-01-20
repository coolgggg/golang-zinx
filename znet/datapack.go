package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/aceld/golang-zinx/utils"
	"github.com/aceld/golang-zinx/ziface"
)

//封包、拆包的具体模块
type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包的头的长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	//data len unit32 是4个字节 + id unit32 4个字节
	return 8
}

//封包方法 |data len|msg id| data|
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放byte字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//将data len写进 data buf 中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//将msg id 写进 data buf 中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//将data数据 写进 data buf中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil

}

//拆包方法 将head消息读出来，之后在根据head消息里的data长度读取
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//解压head信息，得到data len 和 msg id
	msg := &Message{}

	//读data len
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//读msg id
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断data len 是否已经超出了我们也许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data recvive!")
	}

	return msg, nil
}
