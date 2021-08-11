# 5장 Go의 타입 시스템
## 5.1 구조체를 선언하는 방법
예를 들어, 구조체가 있다고 치자.

``` Go
type user struct {
    name        string
    email       string
    ext         int
    privileged  bool
}
```

user 구조체의 변수를 선언하는 방법은 여러가지가 있다.

1. 가장 기본적인 방법으로 하나씩 값을 할당
``` Go
var user1 user // {"", "", 0, false} 로 초기화
user1.email="kim@kim.com"
```

2. 구조체 리터럴을 사용하여 디폴트값으로 생성
``` Go
user2 := user {} // {"", "", 0, false} 로 초기화
```
3. 구조체 리터럴을 사용하여 간략하게 생성
``` Go
user3 := user {"kim", "kim@kim.com", "123", "true"}
```
4. 구조체 리터럴을 사용하되 필드를 타이핑함으로써 알아볼 수 있게 생성
``` Go
user4 := user {
    name:       "kim", 
    email:      "kim@kim.com", 
    ext:        "123", 
    privileged: "true"
}
```
5. 구조체안에 구조체 생성 (기괴)
``` go
fred := admin {
    person: user {
        name:       "kim", 
        email:      "kim@kim.com", 
        ext:        "123", 
        privileged: "true"
    },
    level: "super",
}
```

## 5.2 메서드와 리시버

파라미터를 수신자(receiver, 리시버)라고 부르고, 수신자가 정의된 함수를 메서드(method)라고 부르기로 약속했단다.

``` 함수 + 리시버 = 메서드 ```

``` Go
func (u user) notify() {
    fmt.Printf("알림: %s<%s>\n", u.name, u.email)
}
```



## 5.3 타입을 바라보는 눈은 기본형(primitive)과 비기본형(non-primitive)이다.

책에서는 primitive를 "근원적인", "기반의", "기본적인"라고 스까쓰고있다.


## 5.4 인터페이스
### 5.4.1 인터페이스 개념

``` Go
// 예제 5.34 의 23번 줄과 30번 줄 코드
r, err := http.Get(os.Args[1])  
io.Copy(os.Stdout, r.Body)
``` 
```

p.126 설명
http.Get() 함수는 서버와의 통신이 성공하면 http.Response 타입의 포인터를 리턴한다. 그리고 http.Response 타입은 io.ReaderCloser 인터페이스 타입의 Body 라는 필드를 제공한다.

// r : *(http.Response) {Body : io.ReadCloser} 

다행히 Body 필드는 io.Reader 인터페이스를 구현하고 있기 때문에 Body 필드를 io.Copy 함수에 전달하여 웹 서버를 데이터 원본으로 활용하는 것이 가능하다.

//  ?? 도저히 이해가 안감
 ```
---
### 5.4.2 iface와 eface 고찰

 ``` Go
 // iface와 eface 고찰

 //The following source code is in runtime/runtime2.go
type iface struct {
	tab  *itab              // 동적 타입
	data unsafe.Pointer     // 동적 값(dynamic value)
}

// layout of Itab known to compilers
// allocated in non-garbage-collected memory
// Needs to be in sync with
// ../cmd/compile/internal/gc/reflect.go:/^func.dumptabs.
type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

//The following source code is in runtime/Type.go
type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
```
* iface의 다이어그램

    ![](https://www.fatalerrors.org/images/blog/947245e5c5a801424a816cd4a80102da.jpg)

---

``` Go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```

* eface의 다이어그램

    ![](https://www.fatalerrors.org/images/blog/2e29abd7a4c8b13c062f82b89fac324a.jpg)