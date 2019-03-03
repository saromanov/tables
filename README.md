# tables
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/tables)](https://goreportcard.com/report/github.com/saromanov/tables)
[![Build Status](https://travis-ci.org/saromanov/tables.svg?branch=master)](https://travis-ci.org/saromanov/tables)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/1351fe8963a24b93a273d0c0f91fadfb)](https://www.codacy.com/app/saromanov/tables?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=saromanov/tables&amp;utm_campaign=Badge_Grade)
[![Coverage Status](https://coveralls.io/repos/github/saromanov/tables/badge.svg?branch=master)](https://coveralls.io/github/saromanov/tables?branch=master)

Simple representation of tables on console

### Example
```go

import "github.com/saromanov/tables"

func main() {
	t := tables.New()
	t.AddHeader("foo", "bar")
	t.AddLine("a", "b")
	t.AddLine("c", "d")
	t.Build()
}
```
Output:

```
foo  bar
---  ---
a    b
c    d

```