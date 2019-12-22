package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var reg = []byte{'B', 'C', 'D', 'E', 'H', 'L', 'M', 'A'}
	var fmts = []int8{'e'} //,'g','f','e'}
	fname := "invaders"
	for _, x := range fmts {
		file, err := os.Open(fname + "." + string(x))
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("filename: %s.%c\nSize: %d bytes\n", fname, x, len(data))
		fmt.Println("*************************************************\n\n")
		var opbytes, ind int
		opbytes, ind = 0, 0
		//print(len(data))
		i := 0
		for ind+opbytes < len(data) {
			ind += opbytes
			fmt.Printf("0x%02x(%d)[%x]: ", ind, ind, data[ind])
			switch data[ind] {
			case 0x00:
				fmt.Printf("NOP")
				opbytes = 1

			case 0x01:
				fmt.Printf("LXI B,#$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0x02:
				fmt.Printf("STAX B")
				opbytes = 1

			case 0x03:
				fmt.Printf("INX B")
				opbytes = 1

			case 0x04:
				fmt.Printf("INR B")
				opbytes = 1

			case 0x05:
				fmt.Printf("DCR B")
				opbytes = 1

			case 0x06:
				fmt.Printf("MVI B,#$%02x", data[ind+1])
				opbytes = 2

			case 0x07:
				fmt.Printf("RLC")
				opbytes = 1

			case 0x08:
				fmt.Printf("ERROR!")
				break

			case 0x09:
				fmt.Printf("DAD B")
				opbytes = 1

			case 0x0a:
				fmt.Printf("LDAX B")
				opbytes = 1

			case 0x0b:
				fmt.Printf("DCX B")
				opbytes = 1

			case 0x0c:
				fmt.Printf("INR C")
				opbytes = 1

			case 0x0d:
				fmt.Printf("DCR C")
				opbytes = 1

			case 0x0e:
				fmt.Printf("MVI C,#$%02x", data[ind+1])
				opbytes = 2

			case 0x0f:
				fmt.Printf("RRC")
				opbytes = 1

			case 0x10:
				fmt.Printf("ERROR!")
				break

			case 0x11:
				fmt.Printf("LXI D,#$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0x12:
				fmt.Printf("STAX D")
				opbytes = 1

			case 0x13:
				fmt.Printf("INX D")
				opbytes = 1

			case 0x014:
				fmt.Printf("INR D")
				opbytes = 1

			case 0x15:
				fmt.Printf("DCR D")
				opbytes = 1

			case 0x16:
				fmt.Printf("MVI D,#$%02x", data[ind+1])
				opbytes = 2

			case 0x17:
				fmt.Printf("RAL")
				opbytes = 1

			case 0x18:
				fmt.Printf("ERROR")
				break

			case 0x19:
				fmt.Printf("DAD d")
				opbytes = 1

			case 0x1a:
				fmt.Printf("LDAX D")
				opbytes = 1

			case 0x1b:
				fmt.Printf("DCX D")
				opbytes = 1

			case 0x1c:
				fmt.Printf("INR E")
				opbytes = 1

			case 0x1d:
				fmt.Printf("DCR E")
				opbytes = 1

			case 0x1e:
				fmt.Printf("MVI E,#$%02x", data[ind+1])
				opbytes = 2

			case 0x1f:
				fmt.Printf("RAR")
				opbytes = 1

			case 0x20:
				fmt.Printf("ERROR")
				break

			case 0x21:
				fmt.Printf("LXI H,#$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0x22:
				fmt.Printf("SHLD #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0x23:
				fmt.Printf("INX H")
				opbytes = 1

			case 0x24:
				fmt.Printf("INR H")
				opbytes = 1

			case 0x25:
				fmt.Printf("DCR H")
				opbytes = 1

			case 0x26:
				fmt.Printf("MVI H,#$%02x", data[ind+1])
				opbytes = 2

			case 0x27:
				fmt.Printf("DAA")
				opbytes = 1

			case 0x28:
				fmt.Printf("ERROR")
				break

			case 0x29:
				fmt.Printf("DAD H")
				opbytes = 1

			case 0x2a:
				fmt.Printf("LHLD #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0x2b:
				fmt.Printf("DCX H")
				opbytes = 1

			case 0x2c:
				fmt.Printf("INR L")
				opbytes = 1

			case 0x2d:
				fmt.Printf("DCR L")
				opbytes = 1

			case 0x2e:
				fmt.Printf("MVI L,#$%02x", data[ind+1])
				opbytes = 2

			case 0x2f:
				fmt.Printf("CMA")
				opbytes = 1

			case 0x30:
				fmt.Printf("ERROR")
				break

			case 0x31:
				fmt.Printf("LXI SP,#$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0x32:
				fmt.Printf("STA adr,#$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0x33:
				fmt.Printf("INX SP")
				opbytes = 1

			case 0x34:
				fmt.Printf("INR M")
				opbytes = 1

			case 0x35:
				fmt.Printf("DCR M")
				opbytes = 1

			case 0x36:
				fmt.Printf("MVI M,#$%02x", data[ind+1])
				opbytes = 2

			case 0x37:
				fmt.Printf("STC")
				opbytes = 1

			case 0x38:
				fmt.Printf("ERROR")
				break

			case 0x39:
				fmt.Printf("DAD SP")
				opbytes = 1

			case 0x3a:
				fmt.Printf("LDA #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0x3b:
				fmt.Printf("DCX SP")
				opbytes = 1

			case 0x3c:
				fmt.Printf("INR A")
				opbytes = 1

			case 0x3d:
				fmt.Printf("DCR A")
				opbytes = 1

			case 0x3e:
				fmt.Printf("MVI A,#$%02x", data[ind+1])
				opbytes = 2

			case 0x3f:
				fmt.Printf("CMC")
				opbytes = 1

			case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47:
				fmt.Printf("MOV B,%c", reg[(data[ind]-0x40)])
				opbytes = 1

			case 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f:
				fmt.Printf("MOV C,%c", reg[(data[ind]-0x48)])
				opbytes = 1

			case 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57:
				fmt.Printf("MOV D,%c", reg[(data[ind]-0x50)])
				opbytes = 1
			case 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f:
				fmt.Printf("MOV E,%c", reg[(data[ind]-0x58)])
				opbytes = 1

			case 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67:
				fmt.Printf("MOV H,%c", reg[(data[ind]-0x60)])
				opbytes = 1

			case 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f:
				fmt.Printf("MOV L,%c", reg[(data[ind]-0x68)])
				opbytes = 1

			case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75:
				fmt.Printf("MOV M,%c", reg[(data[ind]-0x70)])
				opbytes = 1

			case 0x76:
				fmt.Printf("HLT")
				opbytes = 1

			case 0x77:
				fmt.Printf("MOV M,A")
				opbytes = 1

			case 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f:
				fmt.Printf("MOV A,%c", reg[(data[ind]-0x78)])
				opbytes = 1

			case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87:
				fmt.Printf("ADD %c", reg[(data[ind]-0x80)])
				opbytes = 1

			case 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f:
				fmt.Printf("ADC %c", reg[(data[ind]-0x88)])
				opbytes = 1

			case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97:
				fmt.Printf("SUB %c", reg[(data[ind]-0x90)])
				opbytes = 1

			case 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f:
				fmt.Printf("SBB %c", reg[(data[ind]-0x98)])
				opbytes = 1

			case 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7:
				fmt.Printf("ANA %c", reg[(data[ind]-0xa0)])
				opbytes = 1

			case 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf:
				fmt.Printf("XRA %c", reg[(data[ind]-0xa8)])
				opbytes = 1

			case 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7:
				fmt.Printf("ORA %c", reg[(data[ind]-0xb0)])
				opbytes = 1

			case 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbe, 0xbf:
				fmt.Printf("CMP %c", reg[(data[ind]-0xb8)])
				opbytes = 1

			case 0xc0:
				fmt.Printf("RNZ")
				opbytes = 1

			case 0xc1:
				fmt.Printf("POP B")
				opbytes = 1

			case 0xc2:
				fmt.Printf("JNZ #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xc3:
				fmt.Printf("JMP #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xc4:
				fmt.Printf("CNZ #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xc5:
				fmt.Printf("PUSH B")
				opbytes = 1

			case 0xc6:
				fmt.Printf("ADI #$%02x", data[ind+1])
				opbytes = 2

			case 0xc7:
				fmt.Printf("RST")
				opbytes = 1

			case 0xc8:
				fmt.Printf("RZ")
				opbytes = 1

			case 0xc9:
				fmt.Printf("RET")
				opbytes = 1

			case 0xca:
				fmt.Printf("JZ #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xcb:
				fmt.Printf("ERROR!")
				break

			case 0xcc:
				fmt.Printf("CZ #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xcd:
				fmt.Printf("CALL #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xce:
				fmt.Printf("ACI #$%02x", data[ind+1])
				opbytes = 2

			case 0xcf:
				fmt.Printf("RST")
				opbytes = 1

			case 0xd0:
				fmt.Printf("RNC")
				opbytes = 1

			case 0xd1:
				fmt.Printf("POP D")
				opbytes = 1

			case 0xd2:
				fmt.Printf("JNC #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xd3:
				fmt.Printf("OUT #$%02x", data[ind+1])
				opbytes = 2

			case 0xd4:
				fmt.Printf("CNC #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xd5:
				fmt.Printf("PUSH D")
				opbytes = 1

			case 0xd6:
				fmt.Printf("SUI #$%02x", data[ind+1])
				opbytes = 2

			case 0xd7:
				fmt.Printf("RST")
				opbytes = 1

			case 0xd8:
				fmt.Printf("RC")
				opbytes = 1

			case 0xd9:
				fmt.Printf("ERROR!")
				break

			case 0xda:
				fmt.Printf("JC #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xdb:
				fmt.Printf("IN #$%02x", data[ind+1])
				opbytes = 2

			case 0xdc:
				fmt.Printf("CC #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xdd:
				fmt.Printf("ERROR!")
				break

			case 0xde:
				fmt.Printf("SBI #$%02x", data[ind+1])
				opbytes = 2

			case 0xdf:
				fmt.Printf("RST")
				opbytes = 1

			case 0xe0:
				fmt.Printf("RPO")
				opbytes = 1

			case 0xe1:
				fmt.Printf("POP H")
				opbytes = 1

			case 0xe2:
				fmt.Printf("JPO #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xe3:
				fmt.Printf("XTHL")
				opbytes = 1

			case 0xe4:
				fmt.Printf("CPO #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xe5:
				fmt.Printf("PUSH H")
				opbytes = 1

			case 0xe6:
				fmt.Printf("ANI #$%02x", data[ind+1])
				opbytes = 2

			case 0xe7:
				fmt.Printf("RST")
				opbytes = 1

			case 0xe8:
				fmt.Printf("RPE")
				opbytes = 1

			case 0xe9:
				fmt.Printf("PCHL")
				opbytes = 1

			case 0xea:
				fmt.Printf("JPE #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xeb:
				fmt.Printf("XCHG")
				opbytes = 1

			case 0xec:
				fmt.Printf("CPE #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xed:
				fmt.Printf("ERROR!")
				break

			case 0xee:
				fmt.Printf("XRI #$%02x", data[ind+1])
				opbytes = 2

			case 0xef:
				fmt.Printf("RST")
				opbytes = 1

			case 0xf0:
				fmt.Printf("RP")
				opbytes = 1

			case 0xf1:
				fmt.Printf("POP PSW")
				opbytes = 1

			case 0xf2:
				fmt.Printf("JP #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xf3:
				fmt.Printf("DI")
				opbytes = 1

			case 0xf4:
				fmt.Printf("CP #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xf5:
				fmt.Printf("PUSH PS")
				opbytes = 1

			case 0xf6:
				fmt.Printf("ORI #$%02x", data[ind+1])
				opbytes = 2

			case 0xf7:
				fmt.Printf("RST")
				opbytes = 1

			case 0xf8:
				fmt.Printf("RM")
				opbytes = 1

			case 0xf9:
				fmt.Printf("SPHL")
				opbytes = 1

			case 0xfa:
				fmt.Printf("JM #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 13

			case 0xfb:
				fmt.Printf("EI")
				opbytes = 1

			case 0xfc:
				fmt.Printf("CM #$%02x%02x", data[ind+2], data[ind+1])
				opbytes = 3

			case 0xfd:
				fmt.Printf("ERROR!")
				break

			case 0xfe:
				fmt.Printf("CPI #$%02x", data[ind+1])
				opbytes = 2

			case 0xff:
				fmt.Printf("RST")
				opbytes = 1
			default:
				opbytes = 1
				fmt.Print("Unimp")
			}
			//fmt.Printf("\n%d %d %x",ind,opbytes,data[ind])
			fmt.Println()
			i += 1
		}
	}
}
