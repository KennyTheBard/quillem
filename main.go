package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	adv "./adversary"
	db "./database"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	N := 10
	b := 5

	f, err := os.OpenFile("data.csv", os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 100; i++ {

		id := strconv.Itoa(i)
		age := strconv.Itoa(r.Intn(90))
		dgx := strconv.Itoa(r.Intn(N))

		if _, err = f.WriteString(id + ",John Doe," + age + "," + dgx + "\n"); err != nil {
			panic(err)
		}
	}

	f.Close()

	src := db.LoadDatabase("data.csv", db.BasicQueryService(), db.BasicResponseService())

	W := make([]int, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N-i; j++ {
			W = append(W, src.Query(db.DGX, j, j+i))
		}
	}

	fmt.Print("W=")
	fmt.Println(W)

	S := adv.Attack(W, b, N)
	for _, s := range S {
		fmt.Println(s)
	}
}
