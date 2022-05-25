# About

This is the improved version of [mypcb](https://github.com/choonsiong/mypcb) added:
- Support multiple input files, e.g. company with different office branches
- Support multiple employee entries

## Usage

Run the command with two input files:
```
15:57:44 220 gopcb on  main +8 -2 ➜ ls
LICENSE    README.md  cmd/       data/      go.mod     go.sum     private/
15:57:48 220 gopcb on  main +8 -2 ➜ ./cmd/gopcb data/test1.json data/test2.json 
15:58:02 220 gopcb on  main +8 -2 ➜
```

Two output files are generated based on the company number and branch number:
```
15:58:04 220 gopcb on  main +8 -2 ➜ cat PCB_00000test1_000branch1_202205.txt 
H00000test1000branch1202205000010000000002000005000000002
D00000010011FOO BAR                                                     ABC12345    ABC12345    P123        9900050000000250009001      
D00000010022ALICE                                                       ABC999      ABC999      P999        9900050000000250009002      
15:58:10 220 gopcb on  main +8 -2 ➜ cat PCB_00000test1_000branch2_202205.txt
H00000test1000branch2202205000100000000003000050000000003
D00000010031DANNY LIM                                                   ABC12345    ABC12345    P123        9900500000002500009003      
D00000010042JACKY CHANG                                                 ABC999      ABC999      P999        9900400000002000009003      
D00000010052WENDY LOW                                                   ABC999      ABC999      P999        9900100000000500009005      
15:58:15 220 gopcb on  main +8 -2 ➜
```