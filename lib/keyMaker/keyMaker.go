package keyMaker

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Token struct {
	keyRunes []rune
	sync.RWMutex
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz-_~")

func Generate() string {
	t := Token{keyRunes: make([]rune, 0, 20)}
	t.Lock()

	defer time.Sleep(1 * time.Nanosecond)
	defer t.Unlock()

	now := strconv.FormatInt(time.Now().Unix(), 10)
	timeRune := []rune(now)
	ns := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(ns)

	for _, r := range timeRune {
		t.keyRunes = append(t.keyRunes, letterRunes[r1.Intn(len(letterRunes))])
		t.keyRunes = append(t.keyRunes, r)
	}
	return string(t.keyRunes)
}
