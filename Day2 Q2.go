package main

import (
	"math/rand"
	"sync"
	"time"
)

type student struct {
	rating   int
	waitTime float32
}

func main() {
	studentList := make([]student, 200)
	var wg sync.WaitGroup

	setValues := func(s *student) {
		s.rating = rand.Intn(10) + 1
		s.waitTime = float32(rand.Intn(10) / 10)
	}
	getRating := func(s student) int {
		time.Sleep(time.Second * time.Duration(s.waitTime))
		return s.rating
	}

	wg.Add(200)
	for i := 0; i < 200; i++ {
		go setValues(&studentList[i])
		wg.Done()
	}
	wg.Wait()

	totalRating := 0
	wg.Add(200)
	for i := 0; i < 200; i++ {
		//println(getRating(studentList[i]))
		totalRating += getRating(studentList[i])
		wg.Done()
	}
	wg.Wait()

	averageRating := totalRating / 200
	println("average rating is ", averageRating)

}
