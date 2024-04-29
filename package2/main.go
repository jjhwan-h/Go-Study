package main

/*
Go가 import된 패키지를 찾는방법
1. go설치시 같이 설치된 기본패키지들은 Go설치경로에서 찾는다. GOROOT 확인=> go env | grep GOROOT
2. 깃허브와 같은 외부저장소에 저장된 패키지의 경우 외부 저장소에서 다운받아 GOPATH\pkg폴더에 설치된다.
3. 현재 모듈아래 위치한 패키지 인지 검사
*/
import (
	"github.com/jjhwan-h/myfirstpkg" //빨간불이 들어오지만 잘가져와진다.
)

func main() {
	myfirstpkg.MyfirstFunc()
}
