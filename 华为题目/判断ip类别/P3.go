package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 255 16  1 2 4 8 16
func main() {
	reader := bufio.NewReader(os.Stdin)
	result := [7]int{}
	for {
		input, err := reader.ReadString('\n')
		if err != nil || strings.TrimSuffix(input, "\n") == "" {
			break
		}
		addr := strings.Split(strings.TrimSuffix(input, "\n"), "~")
		if len(addr) != 2 {
			continue
		}
		addrIp := addr[0]
		addrMask := addr[1]

		if isIpValid(addrIp) {
			addrIpArr := strings.Split(addrIp, ".")
			if addrIpArr[0] == "0" || addrIpArr[0] == "127" {
				continue
			}
		}

		if !isMaskValid(addrMask) {
			//fmt.Println("maskErr", addrMask)
			result[5]++
			continue
		}

		if !isIpValid(addrIp) {
			result[5]++
			continue
		}

		ipType, isPrivate := getIpType(addrIp)
		switch ipType {
		case "A":
			result[0]++
		case "B":
			result[1]++
		case "C":
			//fmt.Println("C:", addr)
			result[2]++
		case "D":
			result[3]++
		case "E":
			result[4]++
		}

		if isPrivate {
			result[6]++
		}
	}
	fmt.Printf("%d %d %d %d %d %d %d", result[0], result[1], result[2], result[3], result[4], result[5], result[6])
}

func getIpType(addrIp string) (ipType string, isPrivate bool) {
	addrIpArr := strings.Split(addrIp, ".")
	firstNum, _ := strconv.Atoi(addrIpArr[0])
	secondNum, _ := strconv.Atoi(addrIpArr[1])
	if firstNum <= 126 && firstNum >= 1 {
		ipType = "A"
		if addrIpArr[0] == "10" {
			isPrivate = true
		}
	}

	if firstNum <= 191 && firstNum >= 128 {
		ipType = "B"
		if firstNum == 172 && secondNum >= 16 && secondNum <= 31 {
			isPrivate = true
		}
	}

	if firstNum <= 223 && firstNum >= 192 {
		ipType = "C"
		if firstNum == 192 && secondNum == 168 {
			isPrivate = true
		}
	}

	if firstNum <= 239 && firstNum >= 224 {
		ipType = "D"
	}

	if firstNum <= 255 && firstNum >= 240 {
		ipType = "E"
	}
	return
}

//255.255.0.245
func isMaskValid(addrMask string) bool {
	if addrMask == "0.0.0.0" || addrMask == "255.255.255.255" {
		return false
	}
	maskArr := strings.Split(addrMask, ".")
	if len(maskArr) != 4 {
		return false
	}
	maskStrBinary := ""
	for i := 0; i < 4; i++ {
		if maskArr[i] < "0" || maskArr[i] > "255" {
			return false
		}
		num, _ := strconv.Atoi(maskArr[i])
		maskBit := strconv.FormatInt(int64(num), 2)
		if len(maskBit) < 8 {
			tmp := ""
			for i := 1; i <= 8-len(maskBit); i++ {
				tmp += "0"
			}
			maskBit = tmp + maskBit
		}
		maskStrBinary += maskBit
	}
	//fmt.Println(maskStrBinary)
	return strings.IndexByte(maskStrBinary, '0') > strings.LastIndexByte(maskStrBinary, '1')
}

func isIpValid(addrIp string) bool {
	addrIpArr := strings.Split(addrIp, ".")
	if len(addrIpArr) != 4 {
		return false
	}
	for i, _ := range addrIpArr {
		num, err := strconv.Atoi(addrIpArr[i])
		if err != nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}
