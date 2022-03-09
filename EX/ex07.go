package main

func main() {
	testMapReadWriteDiffKey()
}

func testMapReadWriteDiffKey() {
	m := make(map[int]int)
	// 不停的写
	go func() {
		for {
			m[100] = 100
		}

	}()

	// 不停的读取
	go func() {
		for {
			_ = m[100]
		}

	}()

	for {

	}
}
