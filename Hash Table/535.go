package main

import (
	"fmt"
	"math/rand"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

type Codec struct {
	mapping map[string]string
}

func Constructor() Codec {
	return Codec{mapping: map[string]string{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	url := "http://tinyUrl/" + RandomString(5)
	this.mapping[url] = longUrl
	return url
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.mapping[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

func main() {
	longUrl := "https://leetcode.com/problems/design-tinyurl"
	obj := Constructor()
	url := obj.encode(longUrl)
	ans := obj.decode(url)
	fmt.Println(ans)

}
