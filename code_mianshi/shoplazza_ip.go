package code_mianshi

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// 将string ip转换为数字, 先实现ipv4
func strIpV4ToNumber(ipv4 string) (uint32, error) {
	ipSplit := strings.Split(ipv4, ".")
	if len(ipSplit) != 4 {
		return 0, fmt.Errorf("strIpV4ToNumber()失败, ipv4段数不足或超出长度4: %s", ipv4)
	}
	var ipInt uint32
	for i, subIpStr := range ipSplit {
		subIp, _ := strconv.ParseUint(subIpStr, 10, 32)
		ipInt |= uint32(subIp) << (24 - 8*i)
	}
	return ipInt, nil
}

func uint32ToStrIp(ip uint32) string {
	bs := *(*[4]byte)(unsafe.Pointer(&ip))
	sb := strings.Builder{}
	// 要倒过来转换
	// 涉及到大端、小端的问题，strIpV4ToNumber()解析时会根据大端存储(低位地址存储ip高位段)
	// 但是按这里直接读取内存的话，会按照小端解析
	for i := range 4 {
		subIp := strconv.Itoa(int(bs[3-i]))
		sb.WriteString(subIp)
		if i != 3 {
			sb.WriteString(".")
		}
	}
	return sb.String()
}

func uint32ToStrIp2(ip uint32) string {
	sb := strings.Builder{}
	for i := range 4 {
		offset := 24 - 8*i
		sub := ip & (0xFF << offset)
		sub >>= offset
		sb.WriteString(strconv.Itoa(int(sub)))
		if i < 3 {
			sb.WriteString(".")
		}
	}
	return sb.String()
}
