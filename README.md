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

#responses
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

`executing command`

```bash
$ gpmarker

# responses
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

`executing command with -r option`

```bash
$ gpmarker -r

# responses
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

`Specifying the path and executing`

```bash
$ gpmarker -p . # specifying dir or file path

# show responses
...
```


## comment pattern
message is not required
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