package main

type config struct {
	//生成shellcode
	loader         string //构造shellcode的方式
	shellcode      string //bin格式shellcode文件
	os             string //windows,linux
	domain         string //签名域名
	hostObfuscator string //混淆远程加载shellcode的地址，干扰蓝队告警日志研判

	//加密 shellcode
	crypt bool
}
