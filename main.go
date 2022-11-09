package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strconv"
)

func encrypt(m string, e, n *big.Int) []string {
	var result []string

	for i := 0; i < len(m); i++ {
		pw := new(big.Int).Exp(big.NewInt(int64([]rune(m)[i])), e, n)
		result = append(result, pw.Text(10))
	}

	return result
}

func decrypt(c []string, d, n *big.Int) ([]int64, []string) {
	var result []int64
	var textResult []string

	for i := 0; i < len(c); i++ {
		cn, _ := strconv.Atoi(c[i])
		pw := new(big.Int).Exp(big.NewInt(int64(cn)), d, n)
		result = append(result, pw.Int64())
		textResult = append(textResult, string(pw.Int64()))
	}

	return result, textResult
}

func generateKeyPair(p, q *big.Int) (e, d, n *big.Int) {
	n = new(big.Int).Mul(p, q)
	fn := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	var ePoss []int64
	for i := 0; int64(i) < n.Int64(); i++ {
		if GCD(fn.Int64(), int64(i)) == 1 {
			ePoss = append(ePoss, int64(i))
		}
	}

	eRand, _ := rand.Int(rand.Reader, big.NewInt(int64(len(ePoss)-1)))
	e = big.NewInt(ePoss[eRand.Int64()])

	for i := 0; int64(i) < n.Int64(); i++ {
		ed := new(big.Int).Mul(e, big.NewInt(int64(i)))
		if new(big.Int).Mod(ed, fn).Cmp(big.NewInt(1)) == 0 {
			d = big.NewInt(int64(i))
		}
	}

	return e, d, n
}

func webservice() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()
		if !query.Has("type") {
			fmt.Fprintln(writer, "specify type: enc, dec, gen")
			return
		}

		typ := query.Get("type")
		switch typ {
		case "enc":
			break
		case "dec":
			break
		case "gen":
			break
		}
	})

	http.HandleFunc("/exit", func(writer http.ResponseWriter, request *http.Request) {
		os.Exit(0)
	})

	fmt.Println("starting server at port 80")
	err := http.ListenAndServe(":80", nil)
	fmt.Println(err)
}

func commandline() {
	/*p, _ := rand.Prime(rand.Reader, 10)
	q, _ := rand.Prime(rand.Reader, 10)*/
	p := big.NewInt(97)
	q := big.NewInt(103)
	e, d, n := generateKeyPair(p, q)
	fmt.Println(e, d, n)

	message := "BAUM"

	fmt.Println("Encrypting...")

	fmt.Println(message)
	fmt.Println([]rune(message))
	encrypted := encrypt(message, e, n)
	fmt.Println(encrypted)

	fmt.Println("Decrypting...")

	decryptedNum, decrypted := decrypt(encrypted, d, n)
	fmt.Println(decryptedNum)
	fmt.Println(decrypted)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("specify running mode as program args: api, cmd")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "api":
		webservice()
		break
	case "cmd":
		commandline()
		break
	}
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
