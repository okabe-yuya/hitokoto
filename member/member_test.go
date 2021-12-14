package member

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("EXEC_MODE", "test")
	os.Setenv("MEMBER", "髙橋,加藤,田中,斉藤,山田,渡辺")

	code := m.Run()
	os.Exit(code)
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func Test_load(t *testing.T) {
	expect := []string { "髙橋","加藤","田中","斉藤","山田","渡辺" }
	loaded := load("MEMBER")
	fmt.Println(loaded)
	for i := 0; i < len(expect); i++ {
		if loaded[i] != expect[i] {
			t.Errorf("Mismatch loaded value %v != %v", loaded[i], expect[i])
		}
	}
}

func Test_trim(t *testing.T) {
	input := []string { "髙橋 ","加藤"," 田中","斉藤 ","山田","渡辺" }
	expect := []string { "髙橋","加藤","田中","斉藤","山田","渡辺" }
	trimed := trim(input)
	for i := 0; i < len(expect); i++ {
		if trimed[i] != expect[i] {
			t.Errorf("Mismatch trimed value %v != %v", trimed[i], expect[i])
		}
	}
}

func Test_shuffle(t *testing.T) {
	input := []string { "髙橋","加藤","田中","斉藤","山田","渡辺" }
	shuffled := shuffle(input)
	for _, v := range(shuffled) {
		if !contains(input, v) {
			t.Errorf("Not include %v in shuffled list", v)
		}
	}
}
