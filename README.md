# openai-golang

Simple OpenAI SDK for Golang to intereract with ChatGPT.

## Contents

Content table:

- [Installation](#installation)
- [Quickstart](#quickstart)
- [Documentation]()

## Installation

Install the package:

```shell
$ go get -u "github.com/leozz37/openai-golang"
```

Import it in your code:

```golang
import "github.com/leozz37/openai-golang"
```

## Quickstart

You need an `API_KEY`. If you don't have one, you can get [here](https://platform.openai.com/account/api-keys).

Copy the [.env.example](.env.example) file as `.env` and set your `API_KEY`:

```shell
INSTANCE_URL=https://api.openai.com/v1/completions
OPENAI_API_KEY=sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```