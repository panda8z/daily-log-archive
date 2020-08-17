package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Packet 二进制封包后的格式
type Packet struct {
	Size uint16
	Body []byte
}

// 将 []byte 数据写入 dataWriter
func writePacket(dataWriter io.Writer, data []byte) error {

	// 准备一个字节数组缓冲
	var buffer bytes.Buffer

	// 将 size 写入缓冲
	if err := binary.Write(&buffer, binary.LittleEndian, uint16(len(data))); err != nil {
		return err
	}

	// 写入包体数据
	if _, err := buffer.Write(data); err != nil {
		return err
	}

	// 获得写入的完整数据
	out := buffer.Bytes()

	// 写入 dataWriter
	if _, err := dataWriter.Write(out); err != nil {
		return err
	}
	return nil
}

// 从 dataReader 中读取封包
func readPacket(dataReader io.Reader) (pkt Packet, err error) {
	// Packet 的Size 为 uint16类型 占两个字节.
	var sizeBuffer = make([]byte, 2)

	// 持续读取两个字节,直到读到为止. 数据不足时回持续阻塞.
	_, err = io.ReadFull(dataReader, sizeBuffer)
	// 发生错误就返回
	if err != nil {
		return
	}
	// 使用 bytes.Reader 读取 sizeBuffer中的数据
	sizeReader := bytes.NewReader(sizeBuffer)

	// 读取小端的 uint16 作为 Packet 的 Size
	err = binary.Read(sizeReader, binary.LittleEndian, &pkt.Size)
	// 发生错误就返回
	if err != nil {
		return
	}

	// 给 Packet封包实例分配包体大小
	pkt.Body = make([]byte, pkt.Size)
	// 读取包体大小.
	_, err = io.ReadFull(dataReader, pkt.Body)
	return
}
