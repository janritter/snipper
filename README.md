# Snipper

[![CircleCI](https://circleci.com/gh/janritter/snipper/tree/main.svg?style=svg)](https://circleci.com/gh/janritter/snipper/tree/main)

> Tool to get code / config snippets from various collections

> **Warning**
> This project is still in the very early alpha phase and currently more like a proof of concept of the idea. Please expect breaking changes and changes in functionality for versions below 1.0.0.

## Prerequisites for development
- Golang 1.18

## Installation via go

### Clone git repo
```bash
git clone git@github.com:janritter/snipper.git
```

### Open project directory
```bash
cd snipper
```

### Install via go
```bash
go install
```

## Installation via Homebrew (For Mac / Linux)

### Get the formula
```bash
brew tap janritter/snipper https://github.com/janritter/snipper
```

### Install formula
```bash
brew install snipper
```

## Additional steps for private repositories

To fix cloning errors when trying to use private repos as snipper collections, execute the following commands:

```bash
ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts
ssh-keyscan -t ecdsa github.com >> ~/.ssh/known_hosts
```

## Usage

### Run help

```bash
snipper
```

### Get snippet - Terraform S3 state backend

```bash
snipper get gh:janritter/snipper-collection terraform state s3
```

## License and Author

Author: Jan Ritter

License: MIT
