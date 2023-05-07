# stag

[![Build Status](https://github.com/SylvainDumas/stag/workflows/ci/badge.svg)](https://github.com/SylvainDumas/stag/actions?query=workflow%3Aci+branch%3Amain)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=github%3A632996387&metric=coverage)](https://sonarcloud.io/summary/new_code?id=github%3A632996387)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The goal of this module is to provide a simple way to associate a behavior with the tag of a Go struct.

## Tech Stack

- **[Go](https://go.dev/)**: any one of the **two latest major** [releases](https://golang.org/doc/devel/release#policy).

## How to use
```go
type Config struct {
	Foo struct {
		Bar string `oneActionTag:"dev"`
	}
	Age int `anotherActionTag:"42"`
}

func TestXxx(t *testing.T) {
	var config Config

	Browse(&config,
		WithTagFn("oneActionTag", oneActionTag),
		WithTagFn("anotherActionTag", anotherActionTag),
	)
}

func oneActionTag(tagContent string, field FieldIf) error {
	// do some stuffs
	return nil
}

func anotherActionTag(tagContent string, field FieldIf) error {
	// do some stuffs
	return nil
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is open source and available under the [MIT License](LICENSE).
