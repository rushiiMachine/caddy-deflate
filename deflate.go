package caddy_deflate

import (
	"fmt"
	"strconv"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp/encode"

	"github.com/klauspost/compress/zlib"
)

func init() {
	caddy.RegisterModule(Deflate{})
}

type Deflate struct {
	Level int `json:"level,omitempty"`
}

// CaddyModule implements caddy.Module
func (Deflate) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.encoders.deflate",
		New: func() caddy.Module { return new(Deflate) },
	}
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler
func (h *Deflate) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	h.Level = 5

	for d.Next() {
		if !d.NextArg() {
			continue
		}

		level, err := strconv.Atoi(d.Val())
		if err != nil {
			return d.WrapErr(err)
		}

		h.Level = level
	}
	return nil
}

// Validate implements caddy.Validator
func (h Deflate) Validate() error {
	if h.Level < -2 || h.Level > 9 {
		return fmt.Errorf("invalid compression level %d, must be in range of [-2,9]", h.Level)
	}

	return nil
}

// AcceptEncoding implements encode.Encoding
func (h Deflate) AcceptEncoding() string { return "deflate" }

// NewEncoder implements encode.Encoding
func (h Deflate) NewEncoder() encode.Encoder {
	writer, err := zlib.NewWriterLevel(nil, h.Level)

	if err != nil {
		panic(err)
	}

	return writer
}

// Interface guards
var (
	_ encode.Encoding       = (*Deflate)(nil)
	_ caddy.Validator       = (*Deflate)(nil)
	_ caddyfile.Unmarshaler = (*Deflate)(nil)
)
