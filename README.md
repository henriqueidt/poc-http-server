# POC HTTP SERVER

The intent of this project is to create a simple HTTP Server written in GO

- HTTP is a TCP/IP-based communication protocol
- HTTP Requests are sent through a TCP socket
- HTTP is stateless, meaning each request is independent of the previous one

## HTTP/0.9 (1991)

- First version of HTTP
- Only supported GET method
- Response was always an html

## HTTP/1.0 (1996)

- New content types supported
- More methods supported (POST, HEAD, etc.)
- HTTP headers introduced
- Status codes introduced
- Can't make multiple requests on the same connection, needs to open a new TCP connection for each request

### Three-way Handshake

- Client and server share a bunch of information before starting to share the application data
- SYN - Client sends a SYN packet to the server
- SYN-ACK - Server responds with a SYN-ACK packet (based on the SYN packet)
- ACK - Client responds with an ACK packet (based on the SYN-ACK packet)

## HTTP/1.1 (1997)

- Additional HTTP methods (PUT, OPTIONS, PATCH, DELETE) introduced
- Hostname identification (Host header) introduced
- Persistent connections introduced. To close the connection, the clients must send a `Connection: close` header, telling the server to close the connection after the response
- Introduced pipelining, allowing multiple requests to be sent on the same connection without waiting for the response of the previous request. For that, the `Content-Length` header must be set, then the client can know when the response ends
- Chunked transfer encoding introduced, allowing the server to send the response in chunks without knowing the total size of the response beforehand.
  - The last chunk is sent with a size of 0 to indicate the end of the response.
  - The `Transfer-Encoding: chunked` header is used to indicate that the response is chunked
- Few other headers and features were introduced:
  - `Cache-Control` header to control caching behavior
  - Byte range requests, allowing clients to request only a part of the resource
  - Character sets, allowing clients to specify the character set used in the response
  - Language negociation, allowing clients to specify the language of the response
  - Client cookies, allowing clients to store state on the client side
  - New status codes
  - Etc
