# REST API Webservice

>Heads Up!

This project enable CORS by default, if you intend to disable it, comment this line of code on controller/endpoint.go:
```
enableCORS(&w)
```

## Setup

Clone this repository:
```bash
git clone https://github.com/ekharisma/REST-webservice.git
```

Install dependency of the project:
```bash
go get -u github.com/eclipse/paho.mqtt.golang
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
```

## Run The Web Service
Run the web service with the following command:
```bash
go run main.go
```
Or, if you want to build it into binary, you can use the following command:
```bash
go build -o <path>/<fileName>
```
## End points
<table>
  <tr>
    <th>No</th>
    <th>End Point</th>
    <th>Description</th>
  <tr>
    <td>1</td>
    <td>/sensor/latest</td>
    <td>Get sensor's latest temperature</td>
  </tr>
  <tr>
    <td>2</td>
    <td>/sensor/min</td>
    <td>Get sensor's minimal temperature</td>
  </tr>
    <tr>
    <td>3</td>
    <td>/sensor/max</td>
    <td>Get sensor's maximal temperature</td>
  </tr>
</table>

## JSON Structure
<table>
  <tr>
    <th>No</th>
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
  <tr>
    <td>1</td>
    <td>Id</td>
    <td>integer</td>
    <td>Respond's id</td>
  </tr>
  <tr>
    <td>2</td>
    <td>Temperature</td>
    <td>integer</td>
    <td>Temperature value</td>
  </tr>
    <tr>
    <td>3</td>
    <td>Timestamp</td>
    <td>string</td>
    <td>Time registered to a responds that records when a responds is transmited</td>
  </tr>
</table>