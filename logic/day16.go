package logic

import (
	"adventofcode/general"
	"fmt"
	"log"
	"math"
)

func Day16Task1() int {
	log.Println("Day 16 task")

	input, err := general.ReadLines("assets/day16.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// visited
	// path tracking
	// nodes
	// data := []string{
	// 	"start-A",
	// 	"start-b",
	// 	"A-c",
	// 	"A-b",
	// 	"b-d",
	// 	"A-end",
	// 	"b-end",
	// }

	data := input[0]

	bits := hexToBitArray(data)
	packet, _ := readPacket(bits, 0)
	sum := sumVersions(packet)
	sum2 := eval(packet)

	fmt.Println(sum)
	fmt.Println(sum2)

	return 0
}

func eval(packet interface{}) int64 {

	switch p := packet.(type) {
	case Literal:
		return p.Value
	case Operator:
		switch p.TypeID {
		case 0:
			sum := int64(0)
			for _, p2 := range p.Packets {
				sum += eval(p2)
			}
			return sum
		case 1:
			product := int64(1)
			for _, p2 := range p.Packets {
				product *= eval(p2)
			}
			return product
		case 2:
			min := int64(math.MaxInt64)
			for _, p2 := range p.Packets {
				v := eval(p2)
				if v < min {
					min = v
				}
			}
			return min
		case 3:
			max := int64(0)
			for _, p2 := range p.Packets {
				v := eval(p2)
				if v > max {
					max = v
				}
			}
			return max

		case 7:
			v1 := eval(p.Packets[0])
			v2 := eval(p.Packets[1])
			if v1 == v2 {
				return 1
			} else {
				return 0
			}
		case 6:
			v1 := eval(p.Packets[0])
			v2 := eval(p.Packets[1])
			if v1 < v2 {
				return 1
			} else {
				return 0
			}
		case 5:
			v1 := eval(p.Packets[0])
			v2 := eval(p.Packets[1])
			if v1 > v2 {
				return 1
			} else {
				return 0
			}
		}

	}
	return -1
}
func sumVersions(packet interface{}) int64 {
	var sum int64
	switch p := packet.(type) {
	case Literal:
		sum += p.Version
	case Operator:
		sum += p.Version
		for _, p2 := range p.Packets {
			sum += sumVersions(p2)
		}
	}
	return sum

}
func hexToBitArray(s string) []byte {
	out := make([]byte, 0, 4*len(s))
	for _, c := range s {
		switch c {
		case '0':
			out = append(out, []byte{0, 0, 0, 0}...)
		case '1':
			out = append(out, []byte{0, 0, 0, 1}...)
		case '2':
			out = append(out, []byte{0, 0, 1, 0}...)
		case '3':
			out = append(out, []byte{0, 0, 1, 1}...)
		case '4':
			out = append(out, []byte{0, 1, 0, 0}...)
		case '5':
			out = append(out, []byte{0, 1, 0, 1}...)
		case '6':
			out = append(out, []byte{0, 1, 1, 0}...)
		case '7':
			out = append(out, []byte{0, 1, 1, 1}...)
		case '8':
			out = append(out, []byte{1, 0, 0, 0}...)
		case '9':
			out = append(out, []byte{1, 0, 0, 1}...)
		case 'A':
			out = append(out, []byte{1, 0, 1, 0}...)
		case 'B':
			out = append(out, []byte{1, 0, 1, 1}...)
		case 'C':
			out = append(out, []byte{1, 1, 0, 0}...)
		case 'D':
			out = append(out, []byte{1, 1, 0, 1}...)
		case 'E':
			out = append(out, []byte{1, 1, 1, 0}...)
		case 'F':
			out = append(out, []byte{1, 1, 1, 1}...)
		}

	}
	return out
}

type Literal struct {
	Version int64
	TypeID  int64
	Value   int64
}

type Operator struct {
	Version  int64
	TypeID   int64
	LengthID int64
	Length   int64
	Packets  []interface{}
}

func readPacket(data []byte, startpos int) (l interface{}, c int) {
	var count int
	n := startpos
	version, count := readBits(data, n, 3)
	n += count

	typeID, count := readBits(data, n, 3)
	n += count

	switch typeID {
	case 4:
		value, count := readNumber(data, n)
		n += count

		return Literal{Version: version,
			TypeID: typeID,
			Value:  value}, n - startpos
	default:
		lengthID, count := readBits(data, n, 1)
		n += count
		if lengthID == 0 {
			length, count := readBits(data, n, 15)
			n += count
			op := Operator{
				Version:  version,
				TypeID:   typeID,
				LengthID: lengthID,
				Length:   length,
				Packets:  nil,
			}

			subpacketStart := n
			for int64(n-subpacketStart) < length {
				packet, count := readPacket(data, n)
				n += count
				op.Packets = append(op.Packets, packet)
			}
			return op, n - startpos
		} else {
			length, count := readBits(data, n, 11)
			n += count

			op := Operator{
				Version:  version,
				TypeID:   typeID,
				LengthID: lengthID,
				Length:   length,
				Packets:  nil,
			}

			for i := int64(0); i < length; i++ {
				packet, count := readPacket(data, n)
				n += count
				op.Packets = append(op.Packets, packet)
			}
			return op, n - startpos
		}
		return nil, 0
	}

}

func readNumber(data []byte, startpos int) (out int64, count int) {
	//var out int64
	for {
		part, _ := readBits(data, startpos, 5)
		out <<= 4
		out |= int64(part & 0x0f)

		count += 5
		if part&0x10 == 0 {
			break
		}
		startpos += 5
	}
	return out, count
}
func readBits(data []byte, startpos, count int) (out int64, c int) {
	//var out int64
	for _, i := range data[startpos : startpos+count] {
		out <<= 1
		out |= int64(i)
	}
	return out, count
}
