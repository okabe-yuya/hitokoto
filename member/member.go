package member

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

func Get() []string {
	member := load("N2i_MEMBER")
	member = trim(member)
	member = shuffle(member)
	return member
}

func load(envName string) []string {
	memberStr := os.Getenv(envName)
	return strings.Split(memberStr, ",")
}

func trim(memmber []string) []string {
	trimed := make([]string, len(memmber))
	for _, v := range memmber {
		t := strings.TrimSpace(v)
		trimed = append(trimed, t)
	}
	return trimed
}

func shuffle(member []string) []string {
	final := make([]string, len(member))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(member))
	for i, v := range perm {
		final[v] = member[i]
	}
	return final
}