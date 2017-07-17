// 예제 4.37  슬라이스에 다른 슬라이스의 원소 추가하기

// 두 개의 정수를 갖는 슬라이스를 두 개 생성한다.
s1 := []int{1, 2}
s2 := []int{3, 4}

// 두 개의 슬라이스를 결합하고 그 결과를 출력한다.
fmt.Printf("%v\n", append(s1, s2...))


// 결과
// [1 2 3 4]