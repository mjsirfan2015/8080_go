reg=['B','C','D','E','H','L','Memory[uint(H)<<8|uint16(L)]','A']
com="mov(reg1,reg2)"
st=0x40
f=open("case",'w')
c='MOV'
for i in range(len(reg)):
    rg1=''
    if i==6:rg1='M'
    else:rg1=reg[i]
    f.write(f"/*%s %s,B..A(0x%x-%x)*/\n"%(c,rg1,st,st+7))
    for j in range(len(reg)):
        rg2=reg[j]
        if j==6:rg2='M'
        f.write(f"\tcase 0x%x://%s %s,%s\n"%(st+j,c,rg1,rg2))
        f.write(f"\t\t%s\n"%(com.replace("reg1,reg2",f"&cpu.%s,&cpu.%s"%(reg[i],reg[j]))))
    st+=8
f.close()