# Korean.go
![KING-Sejong](/image/sejong.jpg)  
> 한글을 간편하게 다룰 수 있는 Go언어 라이브러리.
---
## 설치
아래 명령어를 프로젝트 최상단에 입력하여 라이브러리를 설치합니다.
```bash
go get github.com/Neoration/korean
```
## 사용법
### korean.Text()
```go
str := korean.Text("텍스트")
```
Text 타입을 통해 `Number`를 제외한 함수를 객체처럼 사용할 수 있습니다.  
### korean.Josa
```go
Josa(str, target string) string
```
str에는 대상이 되는 글자를, target에는 조사를 넣어줍니다.  
조사는 아래와 같이 넣어줄 수 있으며, 유효하지 않은 조사일 경우에는 매개변수를 그대로 반환합니다.  
조사 반환 기준 대상은 str의 마지막 글자입니다.  
| 지원하는 조사 | 유효한 조사 매개변수 예시 |
| - | - |
| `을(를)`, `이(가)`, `와(과)`, `은(는)`, `(으)로` `(이)나`, `(이)란`, `(이)든가`, `(이)던가`, `(이)든지`, `(이)나마`, `(이)네`|`을/를`, `이(가)`, `와-과`, `(이)든지`, `은는` |

### korean.IsHangul
```go
IsHangul(str string, onlyCombined ...bool) bool
```
str의 마지막 글자가 한글인지 아닌지 판별하며, onlyCombined 옵션을 true로 줄 경우 해당 글자가 조합형 한글`(가-힣)`인지 판별합니다.  

### korean.IsHangulArray
```go
IsHangulArray(str string, onlyCombined ...bool) []bool
```
str의 모든 글자를 한 자씩 나누어 각각의 한글 여부를 배열에 담아서 반환합니다.  
IsHangul과 똑같이 onlyCombined 옵션을 true로 줄 경우 해당 글자가 조합형 한글`(가-힣)`인지 판별합니다.  

### korean.HasJongSeong
```go
HasJongSeong(str string) bool
```
str의 마지막 글자에 종성이 있는지 없는지를 판별합니다.  

### korean.HasJongSeongArray
```go
HasJongSeongArray(str string) []bool
```
str의 모든 글자를 한 자씩 나누어 각각의 종성 여부를 배열에 담아서 반환합니다.  

### korean.GetSyllableArray
```go
GetSyllableArray(str string, opt ...SyllableOption) [][]string
```
str을 한 글자씩 각각 분리해서 배열에 담아 반환합니다.  
아래와 같이 사용할 수 있으며, 옵션 설명은 패키지 내 주석을 참고하시기 바랍니다.  
```go
a1 := korean.GetSyllableArray("밥값") // [ [ㅂ ㅏ ㅂ] [ㄱ ㅏ ㅄ] ]
a2 := korean.GetSyllableArray("밥값", korean.SyllableOption{
  SeparateHangul: true
}) // [ [ㅂ ㅏ ㅂ] [ㄱ ㅏ ㅂ ㅅ] ]
```

### korean.GeySyllables
```go
GetSyllables(str string, opt ...SyllableOption) []string
```
`GeySyllableArray`의 결과 값에 나온 배열을 한 번 더 전개하여 반환합니다.
```go
a1 := korean.GetSyllableArray("안녕") // [ [ㅇ ㅏ ㄴ] [ㄴ ㅕ ㅇ] ]
a2 := korean.GetSyllables("안녕") // [ ㅇ ㅏ ㄴ ㄴ ㅕ ㅇ ]
```

### korean.Number
```go
Number(num int, opt ...NumberOption) []string
```
한국어로 숫자를 읽는 법을 반환합니다.  
아래와 같이 사용할 수 있으며, 옵션 설명은 패키지 내 주석을 참고하시기 바랍니다.
```go
n1 := korean.Number(12345) // [ 만 이천삼백사십오 ]
n2 := korean.Number(12345, korean.NumberOption{true}) // [ 일만 이천삼백사십오 ]
```