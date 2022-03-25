# LabSec-Challenge


<p> This is my LabSEC solution for challenge proposed</p>


## Part 1: Generating certificates
<p> Challenge 1, 2, 3 and 5 require to generate certificates. To solve that run: </p>


```bash
go run main.go 
```

## Part 2: Server and client connection

<p>
<p> Challenge 4 and 5 talk about connection: server-side (4) and mutual TLS connection (5) </p>
<p><a href ="#4.1 Run client source"> * 1. Run client source</a></p>
<p><a href ="#4.2 Using curl"> * 2. Using curl </a></p>
</p>

### 2.1 Run client source

<p>1. Go to server folder</p>
<p>Challenge 1</p>

```bash
cd challenge4/server/
```
<p>Challenge 5</p>

```bash
cd challenge5/serverMTLS/
```

<p>2. Run server.go</p>

```bash
go run server.go
```

<p>Now you can choose how to make a request: using a client source or curl</p>

<p>3. Go to client folder</p>
<p>Challenge 4</p>

```bash
cd challenge4/client/
```
<p>Challenge 5</p>

```bash
cd challenge5/clientMTLS/
```

<p>4. Run client.go</p>

```bash
go run client.go
```

### 2.2 Using curl

<p>3. Use curl command</p>

<p>Challenge 4</p>

```bash
curl -Lv --cacert 3.servCert.pem  https://localhost:8443/hello
```

<p>Challenge 5</p>

```bash 
curl -Lv --cacert 3.servCert.pem --cert 5.clientCert.pem --key 6.clientKey.pem  https://localhost:8443/hello
```

<p>Note that challenge 5 requires client certificate and key because it's mutual TLS.</p>
