package main

import (
	"net/http"
	"strings"
)

type router struct {
	// 키: http 메서드
	// 값 : URL 패턴별로 실행할 HandlerFunc
	handlers map[string]map[string]HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h HandlerFunc) {
	// http 메서드로 등록된 맵이 있는지 확인
	m, ok := r.handlers[method]
	if !ok {
		// 등록된 맵이 없는 경우 새로운 맵 생성
		m = make(map[string]HandlerFunc)
		r.handlers[method] = m
	}
	// method로 등록된 맵에 URL 패턴과 핸들러 함수 등록
	m[pattern] = h
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func match(pattern, path string) (bool, map[string]string) {
	// 패턴과 패스가 정확히 일치하는 경우 바로 true 를 리턴
	if pattern == path {
		return true, nil
	}

	// 패턴과 패스를 '/' 단위로 구분
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	// 패턴과 패스를 '/' 로 구분한 후 부분 문자열 집합의 개수가 다르면 false 로 리턴
	if len(patterns) != len(paths) {
		return false, nil
	}

	// 패턴에 일치하는 url 파라미터를 담기 위한 params 맵 생성
	params := make(map[string]string)

	// '/' 로 구분된 패턴/패스의 각 문자열을 하나씩 비교
	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
		// 패턴과 패스의 부분 문자열이 일치하는 경우
		// => 바로 다음 루프 실행
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			// 패턴이 ":" 문자열로 시작하는 경우
			// => params 에 url params 를 담은 후 다음 루푸 수행
			params[patterns[i][1:]] = paths[i]
		default:
			// 일치하는 경우가 없으면 false 를 리턴
			return false, nil
		}

	}
	// true 와 params 를 리턴
	return true, params
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// http 메서드에 맞는 모든 handlers 를 반복하여 요청 URL에 해당하는 handler 를 찾음
	for pattern, handler := range r.handlers[req.Method] {
		if ok, params := match(pattern, req.URL.Path); ok {
			// Context 생성
			c := Context{
				Params:         make(map[string]interface{}),
				ResponseWriter: w,
				Request:        req,
			}
			for k, v := range params {
				c.Params[k] = v
			}
			// 요청 URL 에 해당하는 handler 수행
			handler(&c)
			return
		}
	}
	http.NotFound(w, req)
	return
}
