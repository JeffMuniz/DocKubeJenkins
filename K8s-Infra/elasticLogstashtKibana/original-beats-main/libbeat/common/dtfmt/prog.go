// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package dtfmt

import (
	"errors"
)

type prog struct {
	p []byte
}

const (
	opNone              byte = iota
	opCopy1                  // copy next byte
	opCopy2                  // copy next 2 bytes
	opCopy3                  // copy next 3 bytes
	opCopy4                  // copy next 4 bytes
	opCopyShort              // [op, len, content[len]]
	opCopyLong               // [op, len1, len, content[len1<<8 + len]]
	opNum                    // [op, ft]
	opNumPadded              // [op, ft, digits]
	opExtNumPadded           // [op, ft, divExp, digits]
	opExtNumFractPadded      // [op, ft, divExp, digits, fractDigits]
	opZeros                  // [op, count]
	opTwoDigit               // [op, ft]
	opTextShort              // [op, ft]
	opTextLong               // [op, ft]
)

var pow10Table [10]int

func init() {
	x := 1
	for i := range pow10Table {
		pow10Table[i] = x
		x *= 10
	}
}

func (p prog) eval(bytes []byte, ctx *ctx) ([]byte, error) {
	for i := 0; i < len(p.p); {
		op := p.p[i]
		i++

		switch op {
		case opNone:

		case opCopy1:
			bytes = append(bytes, p.p[i])
			i++
		case opCopy2:
			bytes = append(bytes, p.p[i], p.p[i+1])
			i += 2
		case opCopy3:
			bytes = append(bytes, p.p[i], p.p[i+1], p.p[i+2])
			i += 3
		case opCopy4:
			bytes = append(bytes, p.p[i], p.p[i+1], p.p[i+2], p.p[i+3])
			i += 4
		case opCopyShort:
			l := int(p.p[i])
			i++
			bytes = append(bytes, p.p[i:i+l]...)
			i += l
		case opCopyLong:
			l := int(p.p[i])<<8 | int(p.p[i+1])
			i += 2
			bytes = append(bytes, p.p[i:i+l]...)
			i += l
		case opNum:
			ft := fieldType(p.p[i])
			i++
			v := getIntField(ft, ctx)
			bytes = appendUnpadded(bytes, v)
		case opNumPadded:
			ft, digits := fieldType(p.p[i]), int(p.p[i+1])
			i += 2
			v := getIntField(ft, ctx)
			bytes = appendPadded(bytes, v, digits)
		case opExtNumPadded:
			ft, divExp, digits := fieldType(p.p[i]), int(p.p[i+1]), int(p.p[i+2])
			div := pow10Table[divExp]
			i += 3
			v := getIntField(ft, ctx)
			bytes = appendPadded(bytes, v/div, digits)
		case opExtNumFractPadded:
			ft, divExp, digits, fractDigits := fieldType(p.p[i]), int(p.p[i+1]), int(p.p[i+2]), int(p.p[i+3])
			div := pow10Table[divExp]
			i += 4
			v := getIntField(ft, ctx)
			bytes = appendFractPadded(bytes, v/div, digits, fractDigits)
		case opZeros:
			digits := int(p.p[i])
			i++
			for x := 0; x < digits; x++ {
				bytes = append(bytes, '0')
			}
		case opTwoDigit:
			ft := fieldType(p.p[i])
			i++
			v := getIntField(ft, ctx)
			bytes = appendPadded(bytes, v%100, 2)
		case opTextShort:
			ft := fieldType(p.p[i])
			i++
			s, err := getTextFieldShort(ft, ctx)
			if err != nil {
				return bytes, err
			}
			bytes = append(bytes, s...)
		case opTextLong:
			ft := fieldType(p.p[i])
			i++
			s, err := getTextField(ft, ctx)
			if err != nil {
				return bytes, err
			}
			bytes = append(bytes, s...)
		default:
			return bytes, errors.New("unknown opcode")
		}
	}

	return bytes, nil
}

func makeProg(b ...byte) (prog, error) {
	return prog{b}, nil
}

func makeCopy(b []byte) (prog, error) {
	l := len(b)
	switch l {
	case 0:
		return prog{}, nil
	case 1:
		return makeProg(opCopy1, b[0])
	case 2:
		return makeProg(opCopy2, b[0], b[1])
	case 3:
		return makeProg(opCopy2, b[0], b[1], b[2])
	case 4:
		return makeProg(opCopy2, b[0], b[1], b[2], b[3])
	}

	if l < 256 {
		return prog{append([]byte{opCopyShort, byte(l)}, b...)}, nil
	}
	if l < (1 << 16) {
		l1 := byte(l >> 8)
		l2 := byte(l)
		return prog{append([]byte{opCopyLong, l1, l2}, b...)}, nil
	}

	return prog{}, errors.New("literal too long")
}
