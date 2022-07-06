package main

import (
	"bytes"
	cr "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"os"
	"strings"
	"sync"

	//pb "github.com/cheggaaa/pb"

	"github.com/ethereum/go-ethereum/common"
	peer "github.com/ipfs/go-ipfs/source/go-libp2p-peer"
)

type people struct {
	Name       string
	Age        *big.Int
	Country    string
	Height     int
	Province   string
	Profession []string
	sign       map[string][]byte
	friends    map[string][]string
	status     chan error
}

type User struct {
	ID     int64
	Name   string
	Avatar string
}

type groupInfo struct {
	id         string
	peopleInfo sync.Map
}

type StringList struct {
	StringList []string
}

const (
	s             = 100
	GWei          = 1e9
	Token         = 1e18
	BASETIME      = time.RFC3339
	AddressLength = 20
)

func init() {
	fmt.Println("init")
}

var a = func() int {
	fmt.Println("var")
	return 0
}() // 会执行吗？ 会

var ErrNonceTooLow = errors.New("nonce too low")
var tmp = 99

//=======旋转加载start=======
var spinChars = "|/-\\"

type Spinner struct {
	message string
	i       int
}

func NewSpinner(message string) *Spinner {
	return &Spinner{message: message}
}

func (s *Spinner) Tick() {
	fmt.Printf("%s %c \r", s.message, spinChars[s.i])
	s.i = (s.i + 1) % len(spinChars)
}

func isTTY() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

//=======旋转加载end=========

func main() {
	// var s string
	// s = "c7463268d6b2ad969b3457d9e89c5f34e9debe445417fbf50299fe471a28f2c7"
	// fmt.Println("half:", len(s))

	// n, err := fileSize4("/home/zl/Downloads/gitkraken-amd64.deb")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("size:", n)

	// t := UnixToTime(1515527488)
	// fmt.Println(t)

	// timeToStrAndStrToTime()

	// s := "0xdD870fA1b7C4700F2BD7f44238821C26f7392148"
	// res := "0x"
	// element := "2"
	// for i := 0; i < len(s)-2; i++ {
	// 	res += element
	// }
	// fmt.Println(res)

	// s = "0x0000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc40000000000000000000000001111111111111111111111111111111111111111"
	// fmt.Println("a:", s[:66], "\nb:", s[66:130], "\nc:", s[130:])
	// fmt.Println(len("0x2bffc7ed0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000222221111111111111111111111111111112222200000000000000000000000000000000000000000000000000000000000000057171717171000000000000000000000000000000000000000000000000000000"))
	// s = "0x693ec85e000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000057171717171000000000000000000000000000000000000000000000000000000"
	// for i := 0; i < len(s); {
	// 	if i == 0 {
	// 		fmt.Println(s[:10])
	// 		i += 10
	// 		continue
	// 	}
	// 	fmt.Println(s[i : i+64])
	// 	i += 64
	// }

	// a := "0x97187881c653225714D70F1f9349078cC6ED25f1"
	// id, _ := GetIDFromAddress(a)
	// fmt.Println(id)

	// id := "8MJReG4CcrdshmReHAzjrSeZUYYKY1"
	// addr, _ := GetAddressFromID(id)
	// fmt.Println("addr:", addr.String())

	// chalfrequency := stShare(3, 7)
	// fmt.Println(chalfrequency)
	// s := []string{"aaaaaaaaaaaaa", "bbbbbbbbbbbbbb", "vvvvvvvvvvvvvvvvv"}
	// fmt.Println(s)

	// a := big.NewInt(100020003000)
	// s := a.String()
	// fmt.Println(s)

	// InvalidAddr := "0x0000000000000000000000000000000000000000"
	// addr := common.HexToAddress(InvalidAddr)
	// fmt.Println(addr)

	// splitKidPid()

	// var s sync.Map
	// for i := 0; i < 10; i++ {
	// 	s.Store(i, i*10)
	// }

	// s.Range(func(key, value interface{}) bool {
	// 	if key.(int)%2 == 0 {
	// 		fmt.Println(key, value)
	// 		return true
	// 	}
	// 	if key.(int) == 5 {
	// 		fmt.Println("false")
	// 		return false
	// 	}
	// 	fmt.Println(key, "go on")
	// 	return true
	// })

	// p := test("zl")
	// fmt.Println(p, "\n", "p.friends:", p.friends)
	// p.friends["zl"] = append(p.friends["zl"], "111")
	// p.friends["zl"] = append(p.friends["zl"], "222")
	// p.friends["zl"] = append(p.friends["zl"], "111")
	// p.friends["a"] = append(p.friends["zl"], "111")
	// p.friends["b"] = append(p.friends["zl"], "111")
	// p.friends["c"] = append(p.friends["zl"], "111")
	// p.friends["d"] = append(p.friends["zl"], "111")
	// p.friends["e"] = append(p.friends["zl"], "111")
	// fmt.Println("p.friends:", p.friends)
	// delete(p.friends, "a")
	// fmt.Println("p.friends:", p.friends)
	// for i, k := range p.friends["mmm"] {
	// 	fmt.Println(i, k)
	// }

	// p1 := new([]string)
	// fmt.Println(p1, "\n", *p1)

	// p2 := make([]string, 1)
	// fmt.Println(p2)

	//testLock()
	// s := "0x7828e9d3df82e69351d51b8c7a93dbe8a71fecfc70c9ab6c998b4a0105e7f8a7"
	// key := "keeper"
	// hash := crypto.Keccak256Hash([]byte(key))

	// fmt.Println(hash.Hex())
	// fmt.Println(s == hash.Hex())

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", sayHello)

	// server := &http.Server{
	// 	Handler: mux,
	// }

	// listener, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	panic(err)
	// }

	// err = server.Serve(listener)
	// if err != nil {
	// 	panic(err)
	// }

	//旋转加载start
	// flag.Parse()
	// s := NewSpinner("working...")
	// for i := 0; i < 100; i++ {
	// 	if isTTY() {
	// 		s.Tick()
	// 	}
	// 	time.Sleep(100 * time.Millisecond)
	// }
	//旋转加载end

	// //进度条1（下载文件）start
	// req, _ := http.NewRequest("GET", "https://dl.google.com/go/go1.14.2.src.tar.gz", nil)
	// resp, _ := http.DefaultClient.Do(req)
	// defer resp.Body.Close()

	// f, _ := os.OpenFile("go1.14.2.src.tar.gz", os.O_CREATE|os.O_WRONLY, 0644)
	// defer f.Close()

	// bar := progressbar.DefaultBytes(
	// 	-1, //如果设置为-1就是一个旋转器
	// 	"putting object...",
	// )
	// io.Copy(io.MultiWriter(f, bar), resp.Body)
	// //进度条1（下载文件）end

	// bar := progressbar.Default(-1, "uploading...")
	// for i := 0; i < 100; i++ {
	// 	bar.Add(1)
	// 	time.Sleep(30 * time.Millisecond)
	// }
	// bar.Finish()
	// bar := pb.StartNew(0)
	// for i := 0; i < 100; i++ {
	// 	bar.Increment()
	// 	time.Sleep(30 * time.Millisecond)
	// }
	// bar.Finish()

	//产生20个随机数
	// for i := 0; i < 20; i++ {
	// 	r1, r2, err := makeRandomInt()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(r1, " ", r2)
	// }

	s1 := make([]int, 5)
	s2 := make([]int, 5)

	for i := 0; i < 5; i++ {
		s1[i] = i
		s2[i] = s1[i]
	}

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	s2[0] = 100

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println(cap(s1))
}

func makeRandomInt() (int, int, error) {
	// 伪随机
	// 生成随机种子
	rand.Seed(time.Now().Unix())
	// 产生[0,100)范围的伪随机数
	result := rand.Intn(100)

	// 真随机数
	result2, err := cr.Int(cr.Reader, big.NewInt(100))
	if err != nil {
		return 0, 0, err
	}

	return result, int(result2.Int64()), nil
}

func change(he *User) {
	he.ID = 2
}

func testSPrintf(s int) {
	res := fmt.Sprintf("%.3f M", float64(s))
	fmt.Println(res)

	res = fmt.Sprint(s)
	fmt.Println(res)

	res = fmt.Sprintln(s)
	fmt.Println(res)
}

func testSlice() {
	s := make([]int, 1, 3)
	fmt.Println(s[0])
	// s[1] = 1   //wrong
	// s[2] = 2  //wrong
	for i := 1; i < 6; i++ {
		s = append(s, i)
		fmt.Println(s[i])
	}
}

func testMap() {
	m := make(map[int]string, 1)
	fmt.Println(m[0])
	for i := 1; i < 6; i++ {
		m[i] = strconv.Itoa(i)
		fmt.Println(m[i])
	}
}

func app1() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func formatStorageSize(storageSize int64) (size string) {
	if storageSize < 1024 {
		return fmt.Sprintf("%.3fMB", float64(storageSize)/float64(1))
	} else if storageSize < (1024 * 1024) {
		return fmt.Sprintf("%.3fGB", float64(storageSize)/float64(1024))
	} else if storageSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.3fTB", float64(storageSize)/float64(1024*1024))
	} else if storageSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.3fPB", float64(storageSize)/float64(1024*1024*1024))
	} else {
		return fmt.Sprintf("%.3fEB", float64(storageSize)/float64(1024*1024*1024*1024))
	}
}

func formatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.3fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.3fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.3fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.3fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.3fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.3fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	fmt.Println("&t: ", &t)
	return c
}

func testBiBao() {
	a := 5
	b := func() func() {
		c := 10
		return func() {
			fmt.Printf("a,c: %d,%d \n", a, c)
			a *= 3
		}
	}
	b()()
	println(a)
}

func chanTest() {
	ch := make(chan string, 10)
	var arr []string
	var wg0 sync.WaitGroup
	go func() {
		wg0.Add(1)
		for i := range ch {
			arr = append(arr, i)
		}
		wg0.Done()
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			str := getStr()
			ch <- str
		}()
	}

	wg.Wait()
	close(ch)
	wg0.Wait()
	fmt.Println("arr:", arr, "len(arr):", len(arr))

	chtest := make(chan string, 3)
	chtest <- "hello"
	chtest <- "world"
	chtest <- "!"
	close(chtest)
	for i := range chtest {
		fmt.Println("i:", i)
	}
}

func getStr() string {
	strs := []string{"a", "b", "c", "d", "e"}
	i := rand.Intn(len(strs))
	return strs[i]
}

func getUserInfo() *User {
	return &User{ID: 13746731, Name: "ABCDEF", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, ZL!"))
}

func testLock() {
	var tsy sync.Mutex

	tsy.Lock()

	tsy.Unlock()
	tsy.Unlock()
}

func DisorderArray(array []string) []string {
	var temp string
	var num int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(array) - 1; i >= 0; i-- {
		num = r.Intn(i + 1)
		fmt.Println(num)
		temp = array[i]
		array[i] = array[num]
		array[num] = temp
	}

	return array
}

func DisorderArrays(arrays [][]string) [][]string {
	var tmp []string
	var num int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(arrays) - 1; i >= 0; i-- {
		num = r.Intn(i + 1) //0 <= num < i+1
		tmp = arrays[i]
		arrays[i] = arrays[num]
		arrays[num] = tmp
	}
	return arrays
}

//downloadTest中20个文件的md5sum值
// var md5s = [20]string{"276fdc15319051c8d47796941090d4d4", "823f651b12a578e54ee905292294bd69", "e0f9c53c09c4cf6a0ff5345d85789e02", "8852a88676442dcd39fefacfd7e874c8",
// 	"f11a773f2817e2c73dde68ebc0f0d46b", "231eb4500b0dd53200fb746201858bf8", "842bb0ad17783e1ad19fabcc3dbf7df9", "2507fb13b6c7c5a059edfc32d3d3f5bb", "84b23d45c81444a8b482a9b5036e0793",
// 	"7962cad79739601515da8e0e4f79e8dd", "90aa9007bd9f8967fb3c1bbb537b1dd3", "e3f5c59e0a2d32895b869625fa84d1b7", "6d663156843ee606afa48188de66545a", "a63d250eab9607b9f2ed067624330a34",
// 	"40f87d13b8772bbbcdb4fedbc6e27d36", "284e66771c917fe86455495c60210307", "b87e38b9ff18e20ec234e42394a42a3b", "30664d01ffcb00540c3f02e870abf259", "9123231e636a4e5798df009fc1f9edba",
// 	"1d20804f795d1e0f9619c6f380acff48"}

func splitKidPid() {
	IDLength := 30
	metaValue := "8MKX58Ko5vBeJUkfgpkig53jZzwqoW8MGRZbvn8caS431icB2P1uT74B3EHh8MKEFiuo6pyUPwY6JX4UpGxNsHTPRn8MKFCw85uA7XoUtvAvgUt8XHCzST4e8MK5ewoUDg71TomfzKtKn36p6nB9ma8MHSTe6UWXsBYLgFKdEU3VnrnVPRUp/8MHRMZDG6X98SbTWPtLoNZUmArXPzY8MHuP2hhSwEkBmYDFpqbpWzETFQhhQ8MJSfW5oZrs1pyHkhTeFB3ZLUH47fk8MH2URvvhyzd5q9fMNZiB8jMbW4iDw8MJWXckRWs3yQpHVacxWAHAQjfQHue8MKYA1cVjtXt3Xs5d7jT9LZ8HG4e9Y8MGpV2fUoVBZmHVgbC6BffQA1WL3EW8MGiorvtvxGkT37xrJNnbNSnxwa23B8MGXyysTdfqL9q9Fpx9TuBAhSSh6Ed8MGQZ5phgtTgF6g78GBrk95fqmcpA58MGgwwPerdPg4Deb7NyS7iCETA7dFV"
	splitedMeta := strings.Split(metaValue, "/")
	if len(splitedMeta) != 2 {
		return
	}

	kcount := 0
	pcount := 0
	var ks, ps []string
	var has bool
	keepers := splitedMeta[0]
	for i := 0; i < len(keepers)/IDLength; i++ {
		kid := keepers[i*IDLength : (i+1)*IDLength]
		_, err := peer.IDB58Decode(kid)
		if err != nil {
			continue
		}

		has = false

		for _, k := range ks {
			if kid == k {
				has = true
				break
			}
		}

		if !has {
			ks = append(ks, kid)
			kcount++
		}
	}

	providers := splitedMeta[1]
	for i := 0; i < len(providers)/IDLength; i++ {
		kid := providers[i*IDLength : (i+1)*IDLength]
		_, err := peer.IDB58Decode(kid)
		if err != nil {
			continue
		}

		has = false

		for _, k := range ps {
			if kid == k {
				has = true
				break
			}
		}

		if !has {
			ps = append(ps, kid)
			pcount++
		}
	}
	fmt.Println("ks:", kcount, " ps:", pcount)
	fmt.Println(ks)
	fmt.Println(ps)
	return
}

// func test2() []people {
// 	s := make([]people, 0)
// 	for i := 0; i < 5; i++ {
// 		tmp := people{
// 			name: strconv.Itoa(i),
// 		}
// 		s = append(s, tmp)
// 	}

// 	fmt.Println(s)
// 	return s
// }

// func printPeople(p people) {
// 	fmt.Println(p)
// }

// func getPeople(s string) (*people, error) {
// 	p := &people{
// 		name: s,
// 		age:  big.NewInt(24),
// 	}
// 	return p, errors.New("no people")
// }

func stShare(start, end int64) int {
	chalFrequency := 0

	var chalMap sync.Map
	for i := 0; i < 10; i++ {
		chalMap.Store(int64(i), i)
	}

	chalMap.Range(func(k, value interface{}) bool {
		// remove paid challenges
		key := k.(int64)
		if key >= start && key < end {
			chalFrequency++
		}

		return true
	})
	return chalFrequency
}

func GetAddressFromID(id string) (address common.Address, err error) {
	ID, err := peer.IDB58Decode(id)
	if err != nil {
		return common.Address([AddressLength]byte{}), err
	}
	addressByte := []byte(ID)[2:] //因为前两位表示multihash的hash type和hash length
	address = bytesToAddress(addressByte)
	return address, nil
}

func bytesToAddress(b []byte) common.Address {
	var a common.Address
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
	return a
}

func GetIDFromAddress(address string) (id string, err error) {
	addressByte, err := decodeHex(address)
	if err != nil {
		return "", err
	}
	//目前id用的是keccak_256哈希，所以hash type和hash length是[27 20],如果以后更改hash，此处需手动更改值
	var a [22]byte
	a[0] = 27
	a[1] = 20
	copy(a[2:], addressByte)
	ID := peer.ID(string(a[:]))
	id = peer.IDB58Encode(ID)
	return id, nil
}

func decodeHex(hexStr string) (addressByte []byte, err error) {
	addressByte, err = hex.DecodeString(hexStr[2:])
	if err != nil {
		return addressByte, err
	}
	if len(addressByte) != AddressLength {
		return addressByte, errors.New("length is error")
	}
	return addressByte, nil
}

func intToBytes(n int64) []byte {
	s1 := make([]byte, 0)
	buf := bytes.NewBuffer(s1)

	// 数字转 []byte, 网络字节序为大端字节序
	binary.Write(buf, binary.BigEndian, n)
	fmt.Println(buf.Bytes())
	return buf.Bytes()
}

func bytesToInt(n []byte) int64 {
	buf := bytes.NewBuffer(n)
	var s2 int64
	binary.Read(buf, binary.BigEndian, &s2)
	fmt.Println(s2)
	return s2
}

func UnixToTime(timeStamp int64) time.Time {
	return time.Unix(timeStamp, 0).In(time.Local)
}

func timeToStrAndStrToTime() {
	t := time.Now().Unix()
	fmt.Println("t:", t)

	tt := time.Unix(t, 0).Format(BASETIME)
	fmt.Println("t to string:", tt)

	ttt, err := time.Parse(BASETIME, tt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("string to t:", ttt.Unix())
}

func test(s string) people {
	p := people{}
	p.friends = make(map[string][]string)
	p.Name = s

	p.status = make(chan error)
	go func() {
		time.Sleep(1 * time.Second)
		err := errors.New("this is a error!")
		fmt.Println("len(chan):", len(p.status))
		p.status <- err
		fmt.Println("len(chan):", len(p.status))
	}()

	return p
}

func FormatWei(i *big.Int) (result string) {
	f := new(big.Float).SetInt(i)
	res, _ := f.Float64()
	switch {
	case res >= Token:
		result = fmt.Sprintf("%.02f Token", res/Token)
	case res >= GWei:
		result = fmt.Sprintf("%.02f Gwei", res/GWei)
	default:
		result = fmt.Sprintf("%d Wei", i.Int64())
	}
	return
}

func fileSize1(path string) (int, error) {
	file, err := os.Open(path)
	if err == nil {
		sum := 0
		buf := make([]byte, 2014)
		for {
			n, err := file.Read(buf)
			sum += n
			if err == io.EOF {
				break
			}
		}
		return sum, nil
	}
	return 0, err
}

func fileSize2(path string) (int, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	return len(contents), nil
}

func fileSize3(path string) (int64, error) {
	file, err := os.Open(path)
	if err == nil {
		fi, err := file.Stat()
		if err == nil {
			return fi.Size(), nil
		}
	}
	return 0, err
}

func fileSize4(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err == nil {
		return fi.Size(), nil
	}
	return 0, err
}
