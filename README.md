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

## HTTP/2 (2015)

- Binary protocol, instead of text
  - Major building blocks are Frames and Sreams:
    - HTTP messages are split into frames (ie. headers, data, etc.)
    - Frames are binary pieces of data.
    - Streams are a sequence of frames that form a complete HTTP message
    - Each request and response has a unique identifier (stream ID) - divided into frames
    - Each frame has a stream ID, that identifies the stream it belongs to
    - Requests from client uses odd numbers, responses from server uses even numbers
    - RST_STREAM is a special frame type that is used to abort a stream. I.e. the client letting the server know it doesn't need the stream anymore (notice that this won't stop the connection, just the stream. The client can still receive other streams on the same connection)
- Designed for low latency of content transfer
- Multiplexing - Allows multiple requests to be sent on the same connection without waiting for the response of the previous request, similar to HTTP/1.1 pipelining, but more efficient
- Header compression - Reduces the size of the headers sent in the requests and responses
  - Huffman code is used to compress the headers
  - Server and client maintain a table of headers, allowing them to omit repetitive headers in subsequent requests and responses
- Server push - Multiple responses can be sent for a single request, allowing the server to push resources to the client before the client requests them (For exemple, the server can push the CSS and JS files needed for a page before the client requests them)
  - The server sends a PUSH_PROMISE frame to the client, indicating that it will send a response for a specific request and that the client doesn't need to request it
- Request prioritization - Allows the client to specify the priority of the requests, allowing the server to prioritize the responses
- Security

## Why HTTP uses TCP (Transmission Control Protocol)

- TCP is reliable that data is going to be sent in order
- TCP compared to UDP (User Datagram Protocol):
  | | TCP | UDP |
  |----------------|-----|-----|
  | Connection | Yes | No |
  | Handshake | Yes | No |
  | In Order | Yes | No |
  | Fast | No | Yes |

- TCP makes use of sliding windows, that will tell how many packages can be in flight at one time
  - Everytime one package is completely sent, the receiver will send an `ack` to the sender. The sender then can use that information to move the sliding window further
- TCP makes use of handshakes, to acknowledge that the receiver has estabilished the connection before sending data, while UDP just sends the data expecting receiver to receive it, that's why it is much faster

## HTTP?

- TCP doesn't tell what type of data exactly is being transfered, that's why we need HTTP

### HTTP Message Structure

```bash
METHOD /path PROTOCOL-VERSION\r\n
field: value\r\n (header)
field: value\r\n (header)
\r\n (end of headers)
{
  body content
}
```

- **METHOD:** The HTTP method of the request (POST, GET, DELETE, PATCH)
- **/path:** The path of the resource
- **PROTOCOL-VERSION:** The version of the protocol (HTTP/1.1, HTTP/2, etc)
- **\r\n:** End of line indicator (last one is end of headers indicator)
- **field: value:** HTTP Headers

Example:

```bash
GET /posts HTTP/1.1\r\n
Host: posts-example.com\r\n
User-Agent: Mozilla/5.0\r\n
Accept: application/json\r\n
Content-Length: 126\r\n
\r\n

```
