# caddy-deflate

A caddy plugin that adds a [deflate](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding)
compression encoder for caddy.

## Building with xcaddy

```shell
xcaddy build \
  --with github.com/rushiiMachine/caddy-deflate
```

## Samples

The [encode](https://caddyserver.com/docs/caddyfile/directives/encode) directive will have a new format
named `deflate`.\
Valid compression levels are listed
[here](https://github.com/klauspost/compress/blob/0836a1cac5461da096074c0125c507f1b3fc0fdb/flate/deflate.go#L17-L31)
otherwise any value in the range `[0-9]`.

```caddyfile
:80 {
    encode deflate
    respond "guhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguh"
}
```

or

```caddyfile
:80 {
    # Use multiple compressors
    encode deflate gzip zstd
    respond "guhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguh"
}
```

or

```caddyfile
:80 {
    encode {
        # Configure the compression level
        deflate 4
        
        # Configure another backup compressor
        gzip
    }
    respond "guhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguhguh"
}
```
