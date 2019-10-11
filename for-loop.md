---
layout: default
---

# For Loop

The example demonstrates usage of for loop in go.

```
func main() {
	// Variation #1
	var indexA int = 1
	for indexA <= 50 {
		fmt.Println(indexA)
		indexA++
	}

	// Variation #2
	for indexB := 51; indexB <= 100; indexB++ {
		fmt.Println(indexB)
	}

	// Variation #3 Iterating over array
	var numbers [5]int

	numbers[0] = 0
	numbers[1] = 1
	numbers[2] = 2
	numbers[3] = 3
	numbers[4] = 4

	for index, number := range numbers {
		fmt.Printf("Number at index [%d] is %d", index, number)
		fmt.Println()
	}

	// In case if we don't require index then we can use '_' as
	// placeholder for it

	for _, number := range numbers {
		fmt.Println(number)
	}

	//Special case iterating over Map
	//Special case when accessing a key from Map
	//which is not present in Map
	alphabets := make(map[string]string)
	alphabets["A"] = "Apple"
	alphabets["B"] = "Ball"

	for key, value := range alphabets {
		fmt.Println(key, value)
	}
}
```

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/for-loop.go)

### Output

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64
65
66
67
68
69
70
71
72
73
74
75
76
77
78
79
80
81
82
83
84
85
86
87
88
89
90
91
92
93
94
95
96
97
98
99
100
Number at index [0] is 0
Number at index [1] is 1
Number at index [2] is 2
Number at index [3] is 3
Number at index [4] is 4
0
1
2
3
4
A Apple
B Ball
```

[Back](./)
