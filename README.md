# About

This program generates the __MTD data text file__ required for LHDN PCB submission (for example via online banking portal).

This is the improved version of [mypcb](https://github.com/choonsiong/mypcb) added:
- Support multiple input files, e.g. company with different office branches
- Support multiple employee entries

## Reference

- [MTD Data specification](http://lampiran1.hasil.gov.my/pdf/pdfam/FormatDataCP39BI.pdf)
- [Specification for MTD Calculations Using Computerised Calculation for 2019](http://lampiran1.hasil.gov.my/pdf/pdfam/Spesifikasi_Kaedah_Pengiraan_Berkomputer_PCB_2019.pdf)

## Usage

```
$ ls -al
total 5032
drwxr-xr-x   3 choonsiong  staff       96 Oct  5 18:41 ./
drwxr-x---+ 42 choonsiong  staff     1344 Oct  5 18:41 ../
-rwxr-xr-x   1 choonsiong  staff  2574530 Oct  5 18:38 gopcb*
 $ ./gopcb ~/etc/gopcb.json 
 mbp2022 18:45:01 87  ls
PCB_9003393910_9003393910_202209.txt  gopcb*
 $    
```