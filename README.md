# LabSec-Challenge


<p> This is my LabSEC solution for challenge proposed</p>


## Part 1: Generating certificates
<p> Challenge 1, 2, 3 and 5 require to generate certificates. To solve that run: </p>


```bash
go run main.go 
```

## Part 2: Server and client connection

<p>
<p> Challenge 4 and 5 talk about server-side (4) and mutual TLS connection (5) </p>

### Challenge 4

<p>1. Go to server folder</p>

```bash
cd challenge4/server/
```
<p>2. Run server.go</p>

```bash
go run server.go
```
#### Open a new terminal and run the following commands


<p>3. Go to client folder</p>

```bash
cd challenge4/client/
```

<p>4. Run client.go</p>

```bash
go run client.go
```

### Challenge 5
Since server and client tls have to change some especific filds, I created a new file for both.
So before you start running the below commands you have to finish the server created in challenge 4.

1. Go to server folder
```bash
cd challenge5/serverMTLS/
```

2. Run server.go

```bash
go run server.go
```

#### Open a new terminal and run the following commands

3. Go to client folder

```bash
cd challenge5/clientMTLS/
```

4. Run client.go

```bash
go run client.go
```

## Other possibilities to test the connection

### Using curl
After running server.go. Use the curl command:

##### Challenge 4

```bash
curl -Lv --cacert 3.servCert.pem  https://localhost:8443/hello
```
If you don't give a server certificate for, the curl command will not trust in server

##### Challenge 5

```bash 
curl -Lv --cacert 3.servCert.pem --cert 5.clientCert.pem --key 6.clientKey.pem  https://localhost:8443/hello
```

<p>Note that challenge 5 requires client certificate and key because it's mutual TLS.</p>


### Web


