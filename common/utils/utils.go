package utils

import (
	"easychat/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//传输者
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //传输时使用的缓存
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		err = errors.New("conn read len err")
		return
	}
	//根据读到的长度，转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("conn read data err")
		return
	}

	//将收到的数据反序列化成message类型
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		err = errors.New("message unmarshal err")
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给客户端
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	//6.2 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("send data len err=", err)
		return
	}

	//发送消息实体
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("send data err=", err)
		return
	}
	return
}
