Extract malicious PE/DLL from dotnettojscript **VBA** payloads. 

^.^ helping u fight the bad guys during malware analysis

#### Instalation(windows/linux).
```bash
$ git clone https://github.com/pyperanger/rev_dotnettojscript.git # or download from github :)
$ cd rev_dotnettojscript && go build . 
```

Windows: Portable version avaiable :) 


#### Usage

Extract HEX payload from VBA code, cast to binary(ex: xxd -r -p payload.cer). 

Input is the serialized binary 

```bash
Usage of ./rev_dotnettojscript:
  -f string
     Filename (the serialized binary)
  -o string
     Output`
 ```


##### After

Now you have the PE/DLL :) tip: use dotPeek or another decompiler to read the new binary

Base64 payload is the same logic.. thx
