package xor

import (
	"fmt"
	"reflect"
	"testing"
)

func Example() {
	const (
		seed = "Hello World"
		key  = 97
	)

	encrypted := Encrypt(byte(key), []byte(seed))
	fmt.Printf("Encrypt=> key: %d, seed: %s, encryptedText: %v\n", key, seed, encrypted)

	decrypted := Decrypt(byte(key), encrypted)
	fmt.Printf("Decrypt=> key: %d, encryptedText: %v, DecryptedText: %s\n", key, encrypted, string(decrypted))

	// Output:
	// Encrypt=> key: 97, seed: Hello World, encryptedText: [41 4 13 13 14 65 54 14 19 13 5]
	// Decrypt=> key: 97, encryptedText: [41 4 13 13 14 65 54 14 19 13 5], DecryptedText: Hello World
}

var xorTestData = []struct {
	description string
	input       string
	key         int
	encrypted   string
}{
	{
		"Encrypt letter 'a' with key 0 makes no changes",
		"a",
		0,
		"a",
	},
	{
		"Encrypt letter 'a' with key 1",
		"a",
		1,
		"`",
	},
	{
		"Encrypt letter 'a' with key 10",
		"a",
		10,
		"k",
	},
	{
		"Encrypt 'hello world' with key 0 makes no changes",
		"hello world",
		0,
		"hello world",
	},
	{
		"Encrypt 'hello world' with key 1",
		"hello world",
		1,
		"idmmn!vnsme",
	},
	{
		"Encrypt 'hello world' with key 10",
		"hello world",
		10,
		"boffe*}exfn",
	},
	{
		"Encrypt full sentence with key 64",
		"the quick brown fox jumps over the lazy dog.",
		64,
		"4(%`15)#+`\"2/7.`&/8`*5-03`/6%2`4(%`,!:9`$/'n",
	},
	{
		"Encrypt a word with key 32 make the case swap",
		"abcdefghijklmNOPQRSTUVWXYZ",
		32,
		"ABCDEFGHIJKLMnopqrstuvwxyz",
	},
}

func TestXorCipherEncrypt(t *testing.T) {
	for _, test := range xorTestData {
		t.Run(test.description, func(t *testing.T) {
			encrypted := Encrypt(byte(test.key), []byte(test.input))
			if !reflect.DeepEqual(string(encrypted), test.encrypted) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %s, actual %s", test.encrypted, string(encrypted))
			}
		})
	}
}

func TestXorCipherDecrypt(t *testing.T) {
	for _, test := range xorTestData {
		t.Run(test.description, func(t *testing.T) {
			decrypted := Decrypt(byte(test.key), []byte(test.encrypted))

			if !reflect.DeepEqual(string(decrypted), test.input) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %s, actual %s", test.input, string(decrypted))
			}
		})
	}
}

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range xorTestData {
			cipherText := Encrypt(byte(test.key), []byte(test.input))
			Decrypt(byte(test.key), cipherText)
		}
	}
}

func BenchmarkOldEncrypt(b *testing.B) {

	// Encrypt encrypts with Xor encryption after converting each character to byte
	// The returned value might not be readable because there is no guarantee
	// which is within the ASCII range
	// If using other type such as string, []int, or some other types,
	// add the statements for converting the type to []byte.
	oldEncrypt := func(key byte, plaintext []byte) []byte {
		cipherText := []byte{}
		for _, ch := range plaintext {
			cipherText = append(cipherText, key^ch)
		}
		return cipherText
	}

	// Decrypt decrypts with Xor encryption
	oldDecrypt := func(key byte, cipherText []byte) []byte {
		plainText := []byte{}
		for _, ch := range cipherText {
			plainText = append(plainText, key^ch)
		}
		return plainText
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range xorTestData {
			cipherText := oldEncrypt(byte(test.key), []byte(test.input))
			oldDecrypt(byte(test.key), cipherText)
		}
	}
}
