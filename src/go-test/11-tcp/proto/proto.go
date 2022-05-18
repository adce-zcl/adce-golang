package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型
	var length = int32(len(message))
	// 新建一个字节流buf
	var pkg = new(bytes.Buffer)
	// 写入消息头,小端编码
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	// 写入消息体，要把string转换成字节切片
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	// 返回字节切片和err
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度,但是不从流中取出来
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	// 创建一个新的缓冲区，内容是lengthByte
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	// 把byte切片转换为int32这么费劲？
	// binary二进制操作类似内存拷贝，把缓冲区的内容拷贝的length的地址空间去
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// Buffered返回缓冲中现有的可读取的字节数 总字节数=length+4
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}

	// 最后返回将字节切片转换成字符串
	// 从第四个字节开始取，说明peek确实没有把前四个字节拿走
	return string(pack[4:]), nil
}
