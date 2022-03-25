# LabSec-Challenge


<p> This is my LabSEC solution for challenge proposed</p>


# Part 1: Generating certificates
<p> Challenge 1, 2, 3 and 5 require to generate certificates. To solve that run: </p>


```bash
go run main.go 
```

# Part 2: Server and client connection

<p>
<p> Challenge 4 and 5 talk about connection: server-side (4) and mutual TLS connection (5) </p>
<p><a href ="#4.1 Run client source"> 1. Run client source</a></p>
<p><a href ="#4.2 Using curl"> 2. Using curl </a></p>
</p>

# 2.1 Run client source

<p>1. Go to server folder</p>

```bash
# Challenge 4
cd challenge4/server/
# Challenge 5
cd challenge5/serverMTLS/
```

<p>2. Run server.go</p>

```bash
go run server.go
```

<p>3. Go to client folder</p>

```bash
# Challenge 4:
cd challenge4/client/
```

```bash
# Challenge 5:
cd challenge5/clientMTLS/
```

<p>4. Run client.go</p>

```bash
go run client.go
```

# 2.2 Using curl

<p>1. Go to server folder</p>

```bash
# Challenge 4:
cd challenge4/server/
```

```bash
# Challenge 5:
cd challenge5/serverMTLS/
```

<p>2. Run server.go</p>

```bash
go run server.go
```

<p>3. Use curl command</p>

```bash
# Challenge 4:
curl -Lv --cacert 3.servCert.pem  https://localhost:8443/hello
```

```bash
# Challenge 5: 
curl -Lv --cacert 3.servCert.pem --cert 5.clientCert.pem --key 6.clientKey.pem  https://localhost:8443/hello
```

<p>Note that challenge 5 requires client certificate and key because it's mutual TLS.</p>
