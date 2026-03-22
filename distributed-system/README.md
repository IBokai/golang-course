# Distributed-system
Distributed system for obtaining information about Github repository. Allows to get the repository name, description, number of stargazers, number of forks and creation date.

## Usage

1) Clone the repository
```bash
git clone https://github.com/IBokai/golang-course
```
2) Change the directory
```bash
cd golang-course
cd distributed-system
```
3) Run system
```bash
# Run docker container
make docker-up
```
4) Open Swagger UI in browser:
`http://localhost:8080/swagger/index.html`
5) Stop the system
```bash
# Stop and remove docker container
make docker-down
```