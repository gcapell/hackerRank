package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type transaction struct {
	target uint16
	x      uint32
	set    bool // or xor
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	txnLog := make([]transaction, 0, m)
	s := bufio.NewScanner(os.Stdin)
	for j := 0; j < m; j++ {
		s.Scan()
		line := s.Text()
		fields := strings.Fields(line[2:])
		switch line[0] {
		case '1':
			txnLog = append(txnLog, transaction{
				set:    true,
				x:      atoi(fields[0]),
				target: bitmap(fields[1]),
			})
		case '2':
			txnLog = append(txnLog, transaction{
				set:    false,
				x:      atoi(fields[0]),
				target: bitmap(fields[1]),
			})
		case '3':
			output(txnLog, bitmap(fields[0]))
		}
	}
}

func bitmap(bits string) uint16 {
	var reply uint16
	for j, b := range bits {
		if b == '1' {
			reply |= 1 << uint16(j)
		}
	}
	return reply
}

func output(txnLog []transaction, t uint16) {
	var n, xor uint32
	seenXOR := false

	for j := len(txnLog) - 1; j >= 0; j-- {
		op := txnLog[j]
		if op.target&t != t {
			continue
		}
		if op.set {
			n = op.x
			break
		} else {
			xor ^= op.x
			seenXOR = true
		}
	}
	if seenXOR {
		n ^= xor
	}
	fmt.Println(n)
}

func atoi(s string) uint32 {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err, s)
	}
	return uint32(n)
}
