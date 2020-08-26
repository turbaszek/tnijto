![CI](https://github.com/turbaszek/tnijto/workflows/CI/badge.svg?branch=master)

# TnijTo

Easy to deploy link shortener.


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Development](#development)
- [Contributing](#contributing)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Development

To run the app just do:
```
export GCP_PROJECT="your-project-id"
go run -v ./pkg/tnijto.go
```

## Contributing

We welcome all contributions! Please submit an issue or PR no matter if it's bug or a typo.

This project is using [pre-commits](https://pre-commit.com) to ensure the
quality of the code. To install pre-commits just do:
```bash
pip install pre-commit
# or
brew install pre-commit
```
And then from project directory run `pre-commit install`.
