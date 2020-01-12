# Use Environment Variable in your next Golang Project

In this tutorial, we will access environment variables in 3 different ways.

You can use according to your requirement.

  - `os` package
  - `godotenv` package
  - `viper` package

The complete tutorial is available on [blog](https://schadokar.dev/posts/go-env-ways/) :confetti_ball:.

## Test the code :construction:

### Clone the repo outside the $GOPATH

```git
git clone github.com/schadokar/blog-projects

// Check out to go-env-ways
git checkout go-env-ways
```

### Install all the dependencies

```go
go install
```

### Setup the environment

Create a `.env` file and paste the below code.

```toml
STRONGEST_AVENGER=Thor
```

Create a `config.yaml` file and paste the below code.

```yaml
I_AM_INEVITABLE: "I am Iron Man"
```

### Run Test Cases :wrench:

```go
go run test