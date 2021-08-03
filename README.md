# GPMarker

## How to install

```bash
$ go install github.com/gari8/gpmarker/cmd/gpmarker@latest
# or
$ go get -u github.com/gari8/gpmarker/cmd/gpmarker
```

## How to use

```bash
$ tree
.
└── a.go
```

***↓a.go*** `putting mark as comment-out`
```go
package a

import "fmt"

func f() {
	print("ok") // mark:info
	var tiny int // mark
	fmt.Println(tiny) // mark:info it's show time
	println("debug") // mark:warn this print is not require
	fmt.Println("last line!") // mark typo
}
```

`execution command`
`search and show line which putting mark`

```bash
$ gpmarker

=== Warnings
[a.go:9] => this print is not require

=== Information
[a.go:6]
[a.go:8] => it's show time

=== Others
[a.go:7]
[a.go:10] => typo

How to use CLI?:
        gpmarker CLI MODE
        -h: show help message
        -p: you can instruct file path
        -r: show source code near comment
```

`execution command with -r option`

```bash
$ gpmarker -r

=== Warnings
[a.go:9] => this print is not require
9:      println("debug") // mark:warn this print is not require
10:     fmt.Println("last line!") // mark typo
11: }

=== Information
[a.go:6]
6:      print("ok") // mark:info
7:      var tiny int // mark
8:      fmt.Println(tiny) // mark:info it's show time
[a.go:8] => it's show time
8:      fmt.Println(tiny) // mark:info it's show time
9:      println("debug") // mark:warn this print is not require
10:     fmt.Println("last line!") // mark typo

=== Others
[a.go:7]
7:      var tiny int // mark
8:      fmt.Println(tiny) // mark:info it's show time
9:      println("debug") // mark:warn this print is not require
[a.go:10] => typo
10:     fmt.Println("last line!") // mark typo
11: }


How to use CLI?:
        gpmarker CLI MODE
        -h: show help message
        -p: you can instruct file path
        -r: show source code near comment
```


## comment pattern
message is not require
- `// mark [*message]` As Others 
- `// mark:info [*message]` As Information 
- `// mark:warn [*message]` As Warnings

```bash
# responses

=== Warnings
[xxx.go: n] => message
...

=== Information
[xxx.go: n] => message
...

=== Others
[xxx.go: n] => message
...
```