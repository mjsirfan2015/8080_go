package cpu

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	S  = 4
	Z  = 3
	P  = 1
	C  = 0
	AC = 2
)

//Emulate functions to emulate cpu
type Emulate interface {
	Cpuinfo()
	GetFile(fname string, mloc uint16) int
	setfl(flag uint8, b bool)
	getfl(flag uint8) uint8
	excop(opcode uint8)
	setzsp(ans uint8)
	//getreg(rno uint8)uint16
	HL() uint16
	Debug()
	add(val uint16, cy bool, cneg bool)
	ana(val uint8)
	xra(val uint8)
	ora(val uint8)
	inr(reg *uint8)
	dcr(reg *uint8)
	jmp()
	push(*uint8, *uint8)
	pop(*uint8, *uint8)
	call()
	ret()
	dad(r1 *uint8, r2 *uint8)
	lxi(r1 *uint8, r2 *uint8)
	rst(m uint16)
	unimp(uint8)
}

//Cpu emulates 8080 cpu struct
type Cpu struct {
	A      uint8
	B      uint8
	C      uint8
	D      uint8
	E      uint8
	H      uint8
	L      uint8
	SP     uint16
	PC     uint16
	Memory [1 << 16]uint8
	cc     uint8
	IntEn  bool
}

//HL returns M
func (cpu *Cpu) HL() uint8 {
	var h, l uint16 = uint16(cpu.H), uint16(cpu.L)
	return cpu.Memory[h<<8|l]
}

//Debug for debug purposes
func (cpu *Cpu) Debug() {
	cpu.GetFile("invaders.h", 0x000)
	cpu.GetFile("invaders.g", 0x800)
	cpu.GetFile("invaders.f", 0x1000)
	cpu.GetFile("invaders.e", 0x1800)
	/*for i, x := range cpu.Memory[0xada:0xae2] {
		fmt.Printf("0x%x:0x%x\n", (0xada + i), x)
	}*/
	/*cpu.PC = 0x100
	cpu.Memory[368] = 0x7

	//Skip DAA test
	/*cpu.Memory[0x59c] = 0xc3 //JMP
	cpu.Memory[0x59d] = 0xc2
	cpu.Memory[0x59e] = 0x05
	*/
	i := 0
	for true {
		if cpu.PC == 0xade {
			i++
		}
		if cpu.PC == 0xade && i == 100 {

			break
		}
		cpu.excop()
		fmt.Printf("current loc: 0x%x %d\n", cpu.PC, i)
		//fmt.Printf("%d\n", cpu.PC)
	}
	fmt.Println(cpu.Memory[0x2400:0x4000])

}
func (cpu *Cpu) setzsp(ans uint8) {
	cpu.setfl(Z, ans == 0)
	cpu.setfl(S, (ans&0x80) == 0x80)
	cpu.setfl(P, parity(uint16(ans)))
}
func (cpu *Cpu) unimp(op uint8) {
	fmt.Printf("Error! opcode 0x%x Unimplemented!\n", op)
	fmt.Printf("PC value: 0x%x", cpu.PC)
	os.Exit(1)
}

func parity(x uint16) bool {
	y := x ^ (x >> 1)
	y = y ^ (y >> 2)
	y = y ^ (y >> 4)
	return !(y&1 == 1)
}

//Cpuinfo info abt cpu
func (cpu *Cpu) Cpuinfo() {
	fmt.Printf("A= %x\n", cpu.A)
	fmt.Printf("B= %x\n", cpu.B)
	fmt.Printf("C= %x\n", cpu.C)
	fmt.Printf("D= %x\n", cpu.D)
	fmt.Printf("E= %x\n", cpu.E)
	fmt.Printf("H= %x\n", cpu.H)
	fmt.Printf("L= %x\n", cpu.L)
	fmt.Printf("SP= %x\n", cpu.SP)
	fmt.Printf("PC= %x\n", cpu.PC)
	fmt.Printf("S=%d,Z=%d,C=%d,P=%d,AC=%d\n", cpu.getfl(S),
		cpu.getfl(Z), cpu.getfl(C), cpu.getfl(P), cpu.getfl(AC))
}

//GetFile writes file into memory
func (cpu *Cpu) GetFile(fname string, mloc int) int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	for i, x := range data {
		cpu.Memory[mloc+i] = x

	}
	fmt.Println("End: ", mloc+len(data))
	return mloc + len(data)
	/*fmt.Println(cpu.Memory[mloc : mloc+100])
	fmt.Println("....................................................")
	fmt.Println(cpu.Memory[mloc+len(data)-100 : mloc+len(data)])*/
}
func (cpu *Cpu) setfl(flag uint8, b bool) {

	if b == true {
		cpu.cc |= (1 << flag)
	} else {
		cpu.cc &= ^(1 << flag)
	}
}
func (cpu *Cpu) getfl(flag uint8) uint8 {
	return (cpu.cc >> flag) & 1
}

func (cpu *Cpu) excop() {
	op := uint8(cpu.Memory[cpu.PC])
	switch op {
	case 0x0: //fmt.Printf("NOP")//NOP
	//cpu.PC+=1
	case 0x1: //LXI B,D16
		cpu.lxi(&cpu.B, &cpu.C)
	case 0x2: //STAX B
		cpu.Memory[uint16(cpu.B)<<8|uint16(cpu.C)] = cpu.A
	case 0x3: //INX B
		inx(&cpu.B, &cpu.C)
	case 0x4: //INR B
		cpu.inr(&cpu.B)
	case 0x5: //DCR B
		cpu.inr(&cpu.B)
	case 0x6: //MVI B, D8
		mov(&cpu.B, &cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0x7: //RLC
		var x = cpu.A
		cpu.setfl(C, (x&0x80) == 0x80) //set Highest bit
		cpu.A = ((x << 1) | (x&0x80)>>7)
	case 0x8:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0x9: //DAD B
		cpu.dad(cpu.B, cpu.C)
	case 0xa: //LDAX B
		cpu.A = cpu.Memory[uint16(cpu.B)<<8|uint16(cpu.C)]
	case 0xb: //DCX B
		dcx(&cpu.B, &cpu.C)
	case 0xc: //INR C
		cpu.inr(&cpu.C)
	case 0xd: //DCR C
		cpu.dcr(&cpu.C)
	case 0xe: //MVI C,D8
		mov(&cpu.C, &cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0xf: //RRC
		var x uint8 = cpu.A
		cpu.setfl(C, (x&1) == 1) //set Lowest bit
		cpu.A = ((x&1)<<7 | (x >> 1))
	case 0x10:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0x11: //LXI D,D16
		cpu.lxi(&cpu.D, &cpu.E)
	case 0x12: //STAX D
		cpu.Memory[uint16(cpu.D)<<8|uint16(cpu.E)] = cpu.A
	case 0x13: //INX D
		inx(&cpu.D, &cpu.E)
	case 0x14: //INR D
		cpu.inr(&cpu.D)
	case 0x15: //DCR D
		cpu.dcr(&cpu.D)
	case 0x16: //MVI D, D8
		mov(&cpu.D, &cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0x17: //RAL
		var x uint8 = cpu.A
		cpu.A = ((x << 1) | cpu.getfl(C))
		cpu.setfl(C, (x&0x80) == 0x80) //set Lowest bit
	case 0x18:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0x19: //DAD D
		cpu.dad(cpu.D, cpu.E)
	case 0x1a: //LDAX D
		cpu.A = cpu.Memory[uint16(cpu.D)<<8|uint16(cpu.E)]
	case 0x1b: //DCX D
		dcx(&cpu.D, &cpu.E)
	case 0x1c: //INR E
		cpu.inr(&cpu.E)
	case 0x1d: //DCR E
		cpu.dcr(&cpu.E)
	case 0x1e: //MVI E,D8
		mov(&cpu.E, &cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0x1f: //RAR
		var x uint8 = cpu.A
		cpu.A = (cpu.getfl(C)<<7 | (x >> 1))
		cpu.setfl(C, (x&1) == 1) //set Lowest bit
	case 0x20:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0x21: //LXI H,D16
		cpu.lxi(&cpu.H, &cpu.L)
	case 0x22: //SHLD adr
		var adr = uint16(cpu.Memory[cpu.PC+2])<<8 | uint16(cpu.Memory[cpu.PC+1])
		cpu.Memory[adr] = cpu.L
		cpu.Memory[adr+1] = cpu.H
		cpu.PC += 2
	case 0x23: //INX H
		inx(&cpu.H, &cpu.L)
	case 0x24: //INR H
		cpu.inr(&cpu.H)
	case 0x25: //DCR H
		cpu.dcr(&cpu.H)
	case 0x26: //MVI H,D8
		mov(&cpu.H, &cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0x27:
		cpu.unimp(op) //DAA
	//cpu.PC+=1
	case 0x28:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0x29: //DAD H
		cpu.dad(cpu.H, cpu.L)
	case 0x2a: //LHLD adr
		var adr = uint16(cpu.Memory[cpu.PC+2])<<8 | uint16(cpu.Memory[cpu.PC+1])
		cpu.L = cpu.Memory[adr]
		cpu.H = cpu.Memory[adr+1]
		cpu.PC += 2
	case 0x2b: //DCX H
		dcx(&cpu.H, &cpu.L)
	case 0x2c: //INR L
		cpu.inr(&cpu.L)
	case 0x2d: //DCR L
		cpu.dcr(&cpu.L)
	case 0x2e: //MVI L, D8
		mov(&cpu.L, &cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0x2f: //CMA
		cpu.A = ^cpu.A
	case 0x30:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0x31: //LXI SP, D16
		cpu.SP = uint16(cpu.Memory[cpu.PC+2])<<8 | uint16(cpu.Memory[cpu.PC+1])
		cpu.PC += 2
	case 0x32: //STA adr
		cpu.Memory[uint16(cpu.Memory[cpu.PC+2])<<8|uint16(cpu.Memory[cpu.PC+1])] = cpu.A
		cpu.PC += 2
	case 0x33: //INX SP
		cpu.SP++
	case 0x34: //INR M
		cpu.inr(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)])
	case 0x35: //DCR M
		cpu.dcr(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)])
	case 0x36: //MVI M,D8
		mov(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)], &cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0x37: //STC
		cpu.setfl(C, true)
	case 0x38:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0x39: //DAD SP
		var ans uint32
		var hl = uint32(cpu.H)<<8 | uint32(cpu.L)
		ans = hl + uint32(cpu.SP)
		cpu.setfl(C, (ans > 0xffff))
	case 0x3a: //LDA adr
		cpu.A = cpu.Memory[uint16(cpu.Memory[cpu.PC+2])<<8|uint16(cpu.Memory[cpu.PC+1])]
		cpu.PC += 2
	case 0x3b: //DCX SP
		cpu.SP--
	case 0x3c: //INR A
		cpu.inr(&cpu.A)
	case 0x3d: //DCR A
		cpu.dcr(&cpu.A)
	case 0x3e: //MVI A,D8
		mov(&cpu.A, &cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0x3f: //CMC
		var val = !(cpu.getfl(C) == 1)
		//fmt.Println(val)
		cpu.setfl(C, val)
		/*MOV B,B..A(0x40-47)*/
	case 0x40: //MOV B,B
		mov(&cpu.B, &cpu.B)
	case 0x41: //MOV B,C
		mov(&cpu.B, &cpu.C)
	case 0x42: //MOV B,D
		mov(&cpu.B, &cpu.D)
	case 0x43: //MOV B,E
		mov(&cpu.B, &cpu.E)
	case 0x44: //MOV B,H
		mov(&cpu.B, &cpu.H)
	case 0x45: //MOV B,L
		mov(&cpu.B, &cpu.L)
	case 0x46: //MOV B,M
		mov(&cpu.B, &cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)])
	case 0x47: //MOV B,A
		mov(&cpu.B, &cpu.A)
		/*MOV C,B..A(0x48-4f)*/
	case 0x48: //MOV C,B
		mov(&cpu.C, &cpu.B)
	case 0x49: //MOV C,C
		mov(&cpu.C, &cpu.C)
	case 0x4a: //MOV C,D
		mov(&cpu.C, &cpu.D)
	case 0x4b: //MOV C,E
		mov(&cpu.C, &cpu.E)
	case 0x4c: //MOV C,H
		mov(&cpu.C, &cpu.H)
	case 0x4d: //MOV C,L
		mov(&cpu.C, &cpu.L)
	case 0x4e: //MOV C,M
		mov(&cpu.C, &cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)])
	case 0x4f: //MOV C,A
		mov(&cpu.C, &cpu.A)
		/*MOV D,B..A(0x50-57)*/
	case 0x50: //MOV D,B
		mov(&cpu.D, &cpu.B)
	case 0x51: //MOV D,C
		mov(&cpu.D, &cpu.C)
	case 0x52: //MOV D,D
		mov(&cpu.D, &cpu.D)
	case 0x53: //MOV D,E
		mov(&cpu.D, &cpu.E)
	case 0x54: //MOV D,H
		mov(&cpu.D, &cpu.H)
	case 0x55: //MOV D,L
		mov(&cpu.D, &cpu.L)
	case 0x56: //MOV D,M
		mov(&cpu.D, &cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)])
	case 0x57: //MOV D,A
		mov(&cpu.D, &cpu.A)
		/*MOV E,B..A(0x58-5f)*/
	case 0x58: //MOV E,B
		mov(&cpu.E, &cpu.B)
	case 0x59: //MOV E,C
		mov(&cpu.E, &cpu.C)
	case 0x5a: //MOV E,D
		mov(&cpu.E, &cpu.D)
	case 0x5b: //MOV E,E
		mov(&cpu.E, &cpu.E)
	case 0x5c: //MOV E,H
		mov(&cpu.E, &cpu.H)
	case 0x5d: //MOV E,L
		mov(&cpu.E, &cpu.L)
	case 0x5e: //MOV E,M
		mov(&cpu.E, &cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)])
	case 0x5f: //MOV E,A
		mov(&cpu.E, &cpu.A)
		/*MOV H,B..A(0x60-67)*/
	case 0x60: //MOV H,B
		mov(&cpu.H, &cpu.B)
	case 0x61: //MOV H,C
		mov(&cpu.H, &cpu.C)
	case 0x62: //MOV H,D
		mov(&cpu.H, &cpu.D)
	case 0x63: //MOV H,E
		mov(&cpu.H, &cpu.E)
	case 0x64: //MOV H,H
		mov(&cpu.H, &cpu.H)
	case 0x65: //MOV H,L
		mov(&cpu.H, &cpu.L)
	case 0x66: //MOV H,M
		mov(&cpu.H, &cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)])
	case 0x67: //MOV H,A
		mov(&cpu.H, &cpu.A)
		/*MOV L,B..A(0x68-6f)*/
	case 0x68: //MOV L,B
		mov(&cpu.L, &cpu.B)
	case 0x69: //MOV L,C
		mov(&cpu.L, &cpu.C)
	case 0x6a: //MOV L,D
		mov(&cpu.L, &cpu.D)
	case 0x6b: //MOV L,E
		mov(&cpu.L, &cpu.E)
	case 0x6c: //MOV L,H
		mov(&cpu.L, &cpu.H)
	case 0x6d: //MOV L,L
		mov(&cpu.L, &cpu.L)
	case 0x6e: //MOV L,M
		mov(&cpu.L, &cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)])
	case 0x6f: //MOV L,A
		mov(&cpu.L, &cpu.A)
		/*MOV M,B..A(0x70-77)*/
	case 0x70: //MOV M,B
		mov(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)], &cpu.B)
	case 0x71: //MOV M,C
		mov(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)], &cpu.C)
	case 0x72: //MOV M,D
		mov(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)], &cpu.D)
	case 0x73: //MOV M,E
		mov(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)], &cpu.E)
	case 0x74: //MOV M,H
		mov(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)], &cpu.H)
	case 0x75: //MOV M,L
		mov(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)], &cpu.L)
	case 0x76: //HLT
		fmt.Println("Exiting...........")
		os.Exit(0)
	case 0x77: //MOV M,A
		mov(&cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)], &cpu.A)
		/*MOV A,B..A(0x78-7f)*/
	case 0x78: //MOV A,B
		mov(&cpu.A, &cpu.B)
	case 0x79: //MOV A,C
		mov(&cpu.A, &cpu.C)
	case 0x7a: //MOV A,D
		mov(&cpu.A, &cpu.D)
	case 0x7b: //MOV A,E
		mov(&cpu.A, &cpu.E)
	case 0x7c: //MOV A,H
		mov(&cpu.A, &cpu.H)
	case 0x7d: //MOV A,L
		mov(&cpu.A, &cpu.L)
	case 0x7e: //MOV A,M
		mov(&cpu.A, &cpu.Memory[uint16(cpu.H)<<8|uint16(cpu.L)])
	case 0x7f: //MOV A,A
		mov(&cpu.A, &cpu.A)

		/*ADD A..M(0x80-87)*/
	case 0x80: //ADD B
		cpu.add(uint16(cpu.B), false, false)
	case 0x81: //ADD C
		cpu.add(uint16(cpu.C), false, false)
	case 0x82: //ADD D
		cpu.add(uint16(cpu.D), false, false)
	case 0x83: //ADD E
		cpu.add(uint16(cpu.E), false, false)
	case 0x84: //ADD H
		cpu.add(uint16(cpu.H), false, false)
	case 0x85: //ADD L
		cpu.add(uint16(cpu.L), false, false)
	case 0x86: //ADD M
		cpu.add(uint16(cpu.HL()), false, false)
	case 0x87: //ADD A
		cpu.add(uint16(cpu.A), false, false)

		/*ADC A..M(0x88-8f)*/
	case 0x88: //ADC B
		cpu.add(uint16(cpu.B), true, false)
	case 0x89: //ADC C
		cpu.add(uint16(cpu.C), true, false)
	case 0x8a: //ADC D
		cpu.add(uint16(cpu.D), true, false)
	case 0x8b: //ADC E
		cpu.add(uint16(cpu.E), true, false)
	case 0x8c: //ADC H
		cpu.add(uint16(cpu.H), true, false)
	case 0x8d: //ADC L
		cpu.add(uint16(cpu.L), true, false)
	case 0x8e: //ADC M
		cpu.add(uint16(cpu.HL()), true, false)
	case 0x8f: //ADC A
		cpu.add(uint16(cpu.A), true, false)
		/*SUB A..M(0x90-97)*/
	case 0x90: //SUB B
		cpu.add(uint16(cpu.B), false, true)
	case 0x91: //SUB C
		cpu.add(uint16(cpu.C), false, true)
	case 0x92: //SUB D
		cpu.add(uint16(cpu.D), false, true)
	case 0x93: //SUB E
		cpu.add(uint16(cpu.E), false, true)
	case 0x94: //SUB H
		cpu.add(uint16(cpu.H), false, true)
	case 0x95: //SUB L
		cpu.add(uint16(cpu.L), false, true)
	case 0x96: //SUB M
		cpu.add(uint16(cpu.HL()), false, true)
	case 0x97: //SUB A
		cpu.add(uint16(cpu.A), false, true)

		/*SBB A..M(0x98-9f)*/
	case 0x98: //SBB B
		cpu.add(uint16(cpu.B), true, true)
	case 0x99: //SBB C
		cpu.add(uint16(cpu.C), true, true)
	case 0x9a: //SBB D
		cpu.add(uint16(cpu.D), true, true)
	case 0x9b: //SBB E
		cpu.add(uint16(cpu.E), true, true)
	case 0x9c: //SBB H
		cpu.add(uint16(cpu.H), true, true)
	case 0x9d: //SBB L
		cpu.add(uint16(cpu.L), true, true)
	case 0x9e: //SBB M
		cpu.add(uint16(cpu.HL()), true, true)
	case 0x9f: //SBB A
		cpu.add(uint16(cpu.A), true, true)

		/*ANA A..M(0xa0-a7)*/
	case 0xa0: //ANA B
		cpu.ana(cpu.B)
	case 0xa1: //ANA C
		cpu.ana(cpu.C)
	case 0xa2: //ANA D
		cpu.ana(cpu.D)
	case 0xa3: //ANA E
		cpu.ana(cpu.E)
	case 0xa4: //ANA H
		cpu.ana(cpu.H)
	case 0xa5: //ANA L
		cpu.ana(cpu.L)
	case 0xa6: //ANA M
		cpu.ana(cpu.HL())
	case 0xa7: //ANA A
		cpu.ana(cpu.A)
		/*XRA A..M(0xa8-af)*/
	case 0xa8: //XRA B
		cpu.xra(cpu.B)
	case 0xa9: //XRA C
		cpu.xra(cpu.C)
	case 0xaa: //XRA D
		cpu.xra(cpu.D)
	case 0xab: //XRA E
		cpu.xra(cpu.E)
	case 0xac: //XRA H
		cpu.xra(cpu.H)
	case 0xad: //XRA L
		cpu.xra(cpu.L)
	case 0xae: //XRA M
		cpu.xra(cpu.HL())
	case 0xaf: //XRA A
		cpu.xra(cpu.A)
		/*ORA A..M(0xb0-b7)*/
	case 0xb0: //ORA B
		cpu.ora(cpu.B)
	case 0xb1: //ORA C
		cpu.ora(cpu.C)
	case 0xb2: //ORA D
		cpu.ora(cpu.D)
	case 0xb3: //ORA E
		cpu.ora(cpu.E)
	case 0xb4: //ORA H
		cpu.ora(cpu.H)
	case 0xb5: //ORA L
		cpu.ora(cpu.L)
	case 0xb6: //ORA M
		cpu.ora(cpu.HL())
	case 0xb7: //ORA A
		cpu.ora(cpu.A)

		/*CMP A..M(0xb8-bf)*/
	case 0xb8: //CMP B
		cpu.cmp(cpu.B)
	case 0xb9: //CMP C
		cpu.cmp(cpu.C)
	case 0xba: //CMP D
		cpu.cmp(cpu.D)
	case 0xbb: //CMP E
		cpu.cmp(cpu.E)
	case 0xbc: //CMP H
		cpu.cmp(cpu.H)
	case 0xbd: //CMP L
		cpu.cmp(cpu.L)
	case 0xbe: //CMP M
		cpu.cmp(cpu.HL())
	case 0xbf: //CMP A
		cpu.cmp(cpu.A)
	case 0xc0: //RNZ
		if cpu.getfl(Z) != 1 {
			cpu.ret()
		} else {
			cpu.PC += 2
		}
	case 0xc1: //POP B
		cpu.pop(&cpu.B, &cpu.C)
	case 0xc2: //JNZ adr
		if cpu.getfl(Z) != 1 {
			cpu.jmp()
		} else {
			cpu.PC += 2
		}
	case 0xc3: //JMP adr
		cpu.jmp()
	case 0xc4: //CNZ adr
		if cpu.getfl(Z) != 1 {
			cpu.call()
		} else {
			cpu.PC += 2
		}
	case 0xc5: //PUSH B
		cpu.push(&cpu.B, &cpu.C)
	case 0xc6: //ADI D8
		cpu.add(uint16(cpu.Memory[cpu.PC+1]), false, false)
		cpu.PC++
	case 0xc7:
		//RST 0
		cpu.rst(0)
	case 0xc8: //RZ
		if cpu.getfl(Z) == 1 {
			cpu.ret()
		} else {
			cpu.PC += 2
		}
	case 0xc9: //RET
		cpu.ret()
	case 0xca: //JZ adr
		if cpu.getfl(Z) == 1 {
			cpu.jmp()
		} else {
			cpu.PC += 2
		}
	case 0xcb:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0xcc: //CZ adr
		if cpu.getfl(Z) == 1 {
			cpu.call()
		} else {
			cpu.PC += 2
		}
	case 0xcd: //CALL adr
		cpu.call()
	case 0xce: //ACI D8
		cpu.add(uint16(cpu.Memory[cpu.PC+1]), true, false)
		cpu.PC++
	case 0xcf: //RST 1
		cpu.rst(8)
	case 0xd0: //RNC
		if cpu.getfl(C) != 1 {
			cpu.ret()
		} else {
			cpu.PC += 2
		}
	case 0xd1: //POP D
		cpu.pop(&cpu.D, &cpu.E)
	case 0xd2: //JNC adr
		if cpu.getfl(C) != 1 {
			cpu.jmp()
		} else {
			cpu.PC += 2
		}
	case 0xd3: //OUT D8
		cpu.PC += 1
	case 0xd4: //CNC adr
		if cpu.getfl(C) != 1 {
			cpu.call()
		} else {
			cpu.PC += 2
		}
	case 0xd5: //PUSH D
		cpu.push(&cpu.D, &cpu.E)
	case 0xd6: //SUI D8
		cpu.add(uint16(cpu.Memory[cpu.PC+1]), false, true)
		cpu.PC++
	case 0xd7: //RST 2
		cpu.rst(16)
	case 0xd8: //RC
		if cpu.getfl(C) == 1 {
			cpu.ret()
		} else {
			cpu.PC += 2
		}
	case 0xd9:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0xda: //JC adr
		if cpu.getfl(C) == 1 {
			cpu.jmp()
		} else {
			cpu.PC += 2
		}
	case 0xdb: //IN D8
		cpu.PC += 1
	case 0xdc: //CC adr
		if cpu.getfl(C) == 1 {
			cpu.call()
		} else {
			cpu.PC += 2
		}
	case 0xdd:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0xde: //SBI D8
		cpu.add(uint16(cpu.Memory[cpu.PC+1]), true, true)
		cpu.PC++
	case 0xdf: //RST 3
		cpu.rst(24)
	case 0xe0: //RPO
		if cpu.getfl(P) != 1 {
			cpu.ret()
		} else {
			cpu.PC += 2
		}
	case 0xe1: //POP H
		cpu.pop(&cpu.H, &cpu.L)
	case 0xe2: //JPO adr
		if cpu.getfl(P) == 0 {
			cpu.jmp()
		} else {
			cpu.PC += 2
		}
	case 0xe3: //XTHL
		cpu.Memory[cpu.SP], cpu.H = cpu.H, cpu.Memory[cpu.SP]
		cpu.Memory[cpu.SP+1], cpu.L = cpu.L, cpu.Memory[cpu.SP+1]
	case 0xe4: //CPO adr
		if cpu.getfl(P) != 1 {
			cpu.call()
		} else {
			cpu.PC += 2
		}
	case 0xe5: //PUSH H
		cpu.push(&cpu.H, &cpu.L)
	case 0xe6: //ANI D8
		cpu.ana(cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0xe7: //RST 4
		cpu.rst(32)
	case 0xe8: //RPE
		if cpu.getfl(P) == 1 {
			cpu.ret()
		} else {
			cpu.PC += 2
		}
	case 0xe9: //PCHL
		cpu.PC = uint16(cpu.H)<<8 | uint16(cpu.L)
	case 0xea: //JPE adr
		if cpu.getfl(P) == 1 {
			cpu.jmp()
		} else {
			cpu.PC += 2
		}
	case 0xeb: //XCHG
		cpu.H, cpu.L, cpu.D, cpu.E = cpu.D, cpu.E, cpu.H, cpu.L
	case 0xec: //CPE adr
		if cpu.getfl(P) == 1 {
			cpu.call()
		} else {
			cpu.PC += 2
		}
	case 0xed:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0xee: //XRI D8
		cpu.xra(cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0xef: //RST 5
		cpu.rst(40)
	case 0xf0: //RP
		if cpu.getfl(S) == 0 {
			cpu.ret()
		} else {
			cpu.PC += 2
		}
	case 0xf1: //POP PSW
		cpu.pop(&cpu.A, &cpu.cc)
	case 0xf2: //JP adr
		if cpu.getfl(S) == 0 {
			cpu.jmp()
		} else {
			cpu.PC += 2
		}
	case 0xf3: //DI
		cpu.IntEn = false
	case 0xf4: //CP adr
		if cpu.getfl(S) == 0 {
			cpu.call()
		} else {
			cpu.PC += 2
		}
	case 0xf5: //PUSH PSW
		cpu.push(&cpu.A, &cpu.cc)
	case 0xf6: //ORI D8
		cpu.ora(cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0xf7: //RST 6
		cpu.rst(48)
	case 0xf8: //RM
		if cpu.getfl(S) == 1 {
			cpu.ret()
		} else {
			cpu.PC += 2
		}
	case 0xf9: //SPHL
		cpu.SP = uint16(cpu.H)<<8 | uint16(cpu.L)
	case 0xfa: //JM adr
		if cpu.getfl(S) == 1 {
			cpu.jmp()
		} else {
			cpu.PC += 2
		}
	case 0xfb: //EI
		cpu.IntEn = true
	case 0xfc: //CM adr
		if cpu.getfl(S) == 1 {
			cpu.call()
		} else {
			cpu.PC += 2
		}
	case 0xfd:
		cpu.unimp(op) //-
	//cpu.PC+=1
	case 0xfe: //CPI D8
		cpu.cmp(cpu.Memory[cpu.PC+1])
		cpu.PC++
	case 0xff: //RST 7
		cpu.rst(56)

	}
	cpu.PC++
}

func (cpu *Cpu) add(val uint16, adc bool, sub bool) {
	var c, ans uint16
	c = 0
	//fmt.Printf("neg=%t ans=%t",cneg,!(ans>0xff))
	if adc == true {
		c = uint16(cpu.getfl(C))
	}
	if sub == true {
		val = ^(val + c) + 1
	}
	ans = uint16(cpu.A) + val
	if sub == true {
		cpu.setfl(C, !(ans > 0xff))
	} else {
		cpu.setfl(C, ans > 0xff)
	}
	cpu.setfl(Z, (ans&0xff) == 0)
	cpu.setfl(S, (ans&0x80) != 0)
	cpu.setfl(P, parity(ans&0xff))
	//Note:Set AC
	cpu.A = uint8(ans & 0xff)
}

func (cpu *Cpu) ana(val uint8) {
	var ans uint8
	ans = cpu.A & val

	//reset carry flag
	cpu.setfl(C, false)
	cpu.setzsp(ans)
	//Note:Set AC
	cpu.A = ans
}

func (cpu *Cpu) xra(val uint8) {
	var ans uint8
	ans = cpu.A ^ val

	//reset carry flag
	cpu.setfl(C, false)
	cpu.setzsp(ans)
	//Note:Set AC
	cpu.A = ans
}

func (cpu *Cpu) ora(val uint8) {
	var ans uint8
	ans = cpu.A | val
	//reset carry flag
	cpu.setfl(C, false)
	cpu.setzsp(ans)
	//Note:Set AC
	cpu.A = ans
}

func (cpu *Cpu) cmp(val uint8) {
	var x uint8
	x = cpu.A - val
	//reset carry flag
	cpu.setfl(C, cpu.A < val)
	cpu.setzsp(x)
}

func inx(r1 *uint8, r2 *uint8) {
	var val uint16
	val = uint16(*r2) + 1
	*r1 += uint8(val >> 8)
	*r2 = uint8(val & 0xff)
}

func dcx(r1 *uint8, r2 *uint8) {
	var val uint16
	val = uint16(*r1)<<8 | uint16(*r2) - 1
	//fmt.Printf("::%x,%x,%x\n",*r1,*r2,val>>8)
	*r1 = uint8(val >> 8)
	*r2 = uint8(val & 0xff)
}

func mov(dst *uint8, src *uint8) {
	*dst = *src
}

func (cpu *Cpu) inr(reg *uint8) {
	*reg++
	cpu.setzsp(*reg)
	//set AC
}
func (cpu *Cpu) dcr(reg *uint8) {
	*reg--
	cpu.setzsp(*reg)
	//set AC
}

func (cpu *Cpu) jmp() {
	cpu.PC = uint16(cpu.Memory[cpu.PC+2])<<8 | uint16(cpu.Memory[cpu.PC+1]) - 1
}

func (cpu *Cpu) push(r1 *uint8, r2 *uint8) {
	cpu.Memory[cpu.SP-2] = *r2
	cpu.Memory[cpu.SP-1] = *r1
	cpu.SP -= 2
}

func (cpu *Cpu) pop(r1 *uint8, r2 *uint8) {
	*r2 = cpu.Memory[cpu.SP]
	*r1 = cpu.Memory[cpu.SP+1]
	cpu.SP += 2
}

func (cpu *Cpu) call() {
	var ret = cpu.PC + 2
	var m1, m2 uint8 = uint8((ret >> 8) & 0xff), uint8(ret & 0xff)
	cpu.push(&m1, &m2)
	cpu.jmp()
}

func (cpu *Cpu) ret() {
	cpu.PC = uint16(cpu.Memory[cpu.SP]) | (uint16(cpu.Memory[cpu.SP+1]) << 8)
	cpu.SP += 2
}

func (cpu *Cpu) dad(r1 uint8, r2 uint8) {
	var ans uint32
	var hl = uint32(cpu.H)<<8 | uint32(cpu.L)
	var r = uint32(r1)<<8 | uint32(r1)
	ans = hl + r
	cpu.H = uint8(ans >> 8)
	cpu.L = uint8(ans & 0xff)
	//fmt.Printf("%x\n", ans)
	cpu.setfl(C, (ans > 0xffff))
}

func (cpu *Cpu) lxi(r1 *uint8, r2 *uint8) {
	*r1 = cpu.Memory[cpu.PC+2]
	*r2 = cpu.Memory[cpu.PC+1]
	cpu.PC += 2
}

func (cpu *Cpu) rst(m uint16) {
	var ret = cpu.PC + 2
	var m1, m2 uint8 = uint8((ret >> 8) & 0xff), uint8(ret & 0xff)
	cpu.push(&m1, &m2)
	cpu.PC = m
}
