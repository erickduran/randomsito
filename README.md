# randomsito
A simple CLI for interative classes. 

## Usage
### Prerequisistes
- Golang
- Docker (or a MongoDB)

### Installation
```bash
go get github.com/erickduran/randomsito
```

### Run
To run, first start your MongoDB locally using Docker:
```bash
docker run -p 27017:27017 -v ~/data:/data/db mongo:3.6-xenial
```
You may replace `~/data` with some place where you want your data to be persisted. You can run without the volume if you just just to test without saving any data.

Finally, run:
```bash
randomsito
```
NOTE: this CLI is only intended for local usage at the moment, as MongoDB authentication has not been implemented.

### Configuration
If you want to define your configuration using a file, you can define `~/.randomsito.yaml` as:
```yaml
language: en
mongodb_host: localhost
mongodb_port: 27017
mongodb_name: randomsito
```

## Project status
Work in progress. Pending features:
- Add auth to MongoDB
- Grading mode
- Quiz mode
- Disable students
- Remove students
- Import/export functionalities

## Author
Copyright © 2020, Erick Durán.
