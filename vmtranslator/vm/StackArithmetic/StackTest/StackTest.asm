@256
D=A
@SP
M=D

@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE856231
D;JEQ
@SP
A=M
M=0
@NEXT856231
0;JMP
(TRUE856231)
@SP
A=M
M=0
M=-1
(NEXT856231)
@SP
M=M+1

@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@16
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE70010
D;JEQ
@SP
A=M
M=0
@NEXT70010
0;JMP
(TRUE70010)
@SP
A=M
M=0
M=-1
(NEXT70010)
@SP
M=M+1

@16
D=A
@SP
A=M
M=D
@SP
M=M+1

@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE44701
D;JEQ
@SP
A=M
M=0
@NEXT44701
0;JMP
(TRUE44701)
@SP
A=M
M=0
M=-1
(NEXT44701)
@SP
M=M+1

@892
D=A
@SP
A=M
M=D
@SP
M=M+1

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE451583
D;JLT
@SP
A=M
M=0
@NEXT451583
0;JMP
(TRUE451583)
@SP
A=M
M=0
M=-1
(NEXT451583)
@SP
M=M+1

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@892
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE465399
D;JLT
@SP
A=M
M=0
@NEXT465399
0;JMP
(TRUE465399)
@SP
A=M
M=0
M=-1
(NEXT465399)
@SP
M=M+1

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE464069
D;JLT
@SP
A=M
M=0
@NEXT464069
0;JMP
(TRUE464069)
@SP
A=M
M=0
M=-1
(NEXT464069)
@SP
M=M+1

@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE854157
D;JGT
@SP
A=M
M=0
@NEXT854157
0;JMP
(TRUE854157)
@SP
A=M
M=0
M=-1
(NEXT854157)
@SP
M=M+1

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE743015
D;JGT
@SP
A=M
M=0
@NEXT743015
0;JMP
(TRUE743015)
@SP
A=M
M=0
M=-1
(NEXT743015)
@SP
M=M+1

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE383554
D;JGT
@SP
A=M
M=0
@NEXT383554
0;JMP
(TRUE383554)
@SP
A=M
M=0
M=-1
(NEXT383554)
@SP
M=M+1

@57
D=A
@SP
A=M
M=D
@SP
M=M+1

@31
D=A
@SP
A=M
M=D
@SP
M=M+1

@53
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
A=M
A=A-1
D=M
A=A-1
M=M+D
@SP
M=M-1

@112
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
A=M
A=A-1
D=M
A=A-1
M=M-D
@SP
M=M-1

@SP
A=M
A=A-1
M=-M

@SP
A=M
A=A-1
D=M
A=A-1
M=M&D
@SP
M=M-1

@82
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
A=M
A=A-1
D=M
A=A-1
M=M|D
@SP
M=M-1

@SP
A=M
A=A-1
M=!M