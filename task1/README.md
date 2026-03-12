# CLI-Tool
CLI-tool for obtaining information about Github repository. Allows to get the repository name, description, number of stargazers, number of forks and creation date.

## Usage

1) Clone the repository
```bash
git clone https://github.com/IBokai/golang-course
```
2) Change the directory
```bash
cd golang-course
cd task1
```
3) Run
```bash
# You can simply run it with "go run"
go run main.go <repository-url or owner/repo-name>

# Or build an executable with "go build"
go build
./clitool <repository-url or owner/repo-name>
```

## Usage examples

```bash
# URL with https
./clitool https://github.com/suvorovrain/golang-course

# URL with https
./clitool http://github.com/suvorovrain/golang-course

# URL with .git
./clitool https://github.com/suvorovrain/golang-course.git

# owner/repo-name
./clitool suvorovrain/golang-course
```