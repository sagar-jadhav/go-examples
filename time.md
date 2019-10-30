---
layout: default
---

# Time

The example demonstrates the various things you can do with time in Go

When formatting time, you can use one of the [format constants](https://golang.org/pkg/time/#pkg-constants), or write your own using the following options:

| Type     | Options                                          |
| -------- | ------------------------------------------------ |
| Year     | 06 2006                                          |
| Month    | 01 1 Jan January                                 |
| Day      | 02 2 \_2 (width two, right justified)            |
| Weekday  | Mon Monday                                       |
| Hours    | 03 3 15                                          |
| Minutes  | 04 4                                             |
| Seconds  | 05 5                                             |
| ms μs ns | .000 .000000 .000000000                          |
| ms μs ns | .999 .999999 .999999999 (trailing zeros removed) |
| am/pm    | PM pm                                            |
| Timezone | MST                                              |
| Offset   | -0700 -07 -07:00 Z0700 Z07:00                    |

[_table source_](https://yourbasic.org/golang/format-parse-string-time-date-example/#layout-options)

```Go
func main() {
    /* Now() returns the current local time
    in the format: 2019-10-30 17:29:02.6455037 -0500 CDT m=+0.004001001 */
    now := time.Now()
    fmt.Println(now)

    // parse time
    date := "November 28, 2019 5:00 PM CST"
    dateFormat := "January 2, 2006 3:04 PM MST"
    t, _ := time.Parse(dateFormat, date)
    fmt.Println(t)

    // format time using 'RFC822' predefined format
    formattedTime := t.Format(time.RFC822)
    fmt.Println(formattedTime)

    // Add 1 year to the current time
    later := now.AddDate(1, 0, 0)
    fmt.Println(later)

    // Subtract 1 month from current date
    past := now.AddDate(0, -1, 0)
    fmt.Println(past)

    // Add 1 day to current time
    later = now.AddDate(0, 0, 1)
    fmt.Println(later)

    // Add 5 hours to the current time
    later = now.Add(5 * time.Hour)
    fmt.Println(later)

    // Subtract 30 minutes to the current time
    past = now.Add(-30 * time.Minute)
    fmt.Println(past)

    // Add 30 seconds to the current time
    later = now.Add(30 * time.Second)
    fmt.Println(later)

    // Print different parts of the date
    fmt.Println(now.Year())
    fmt.Println(now.Month())
    fmt.Println(now.Day())
    fmt.Println(now.Hour())
    fmt.Println(now.Minute())
    fmt.Println(now.Second())
    fmt.Println(now.Zone())

    // wait 2 seconds
    time.Sleep(2 * time.Second)

    /* calculate the duration
    Since() returns the time elapsed*/
    duration := time.Since(now)
    fmt.Println(duration)
}
```

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/time.go)

### Output

```plain
2019-10-30 18:12:25.6052925 -0500 CDT m=+0.005999201
2019-11-28 17:00:00 -0600 CST
28 Nov 19 17:00 CST
2020-10-30 18:12:25.6052925 -0500 CDT
2019-09-30 18:12:25.6052925 -0500 CDT
2019-10-31 18:12:25.6052925 -0500 CDT
2019-10-30 23:12:25.6052925 -0500 CDT m=+18000.005999201
2019-10-30 17:42:25.6052925 -0500 CDT m=-1799.994000799
2019-10-30 18:12:55.6052925 -0500 CDT m=+30.005999201
2019
October
30
18
12
25
CDT -18000
2.0030032s
```

[Back](./)
