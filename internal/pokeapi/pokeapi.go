package pokeapi

import (
    "net/http"
)

// 별도의 패키지를 생성한 이유?
// 사용자의 입력을 처리하는 REPL과 HTTP 별도 로직이기 때문이다.
// 그래서 이 패키지는?
// http 요청을 보내고, 받아온 응답을 JSON 파싱하여 원하는 결과를 출력한다.

// TODO: 구조체를 정의한다. 이 구조체는 JSON 파싱을 위한 구조체이다.

// TODO: 해당 함수는 location-area 엔드포인트로부터 위치 영역 목록을 반환한다.
// JSON 데이터를 받아서, 구조체의 정의에 따라 파싱한다.
func getLocationAreas() []string {
	result := make([]string)
	location_area_url := "https://pokeapi.co/api/v2/location-area/"

	return result
}


http.Get("")

func 
