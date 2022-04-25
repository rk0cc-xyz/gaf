package storage

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/rk0cc-xyz/gaf/structure"
	"github.com/ulikunitz/xz"
)

// Compress content by XZ.
func compressContent(content []structure.GitHubRepositoryStructure) ([]byte, error) {
	marshaled, merr := json.Marshal(content)
	if merr != nil {
		return nil, merr
	}

	var buf bytes.Buffer

	xzw, xzwerr := xz.NewWriter(&buf)
	if xzwerr != nil {
		return nil, xzwerr
	}
	if _, xzwwerr := xzw.Write(marshaled); xzwwerr != nil {
		return nil, xzwwerr
	}
	if xzwcerr := xzw.Close(); xzwcerr != nil {
		return nil, xzwcerr
	}

	return buf.Bytes(), nil
}

// Decompress XZ to content.
func decompressContent(compressed []byte) ([]structure.GitHubRepositoryStructure, error) {
	xzr, xzrerr := xz.NewReader(bytes.NewBuffer(compressed))
	if xzrerr != nil {
		return nil, xzrerr
	}

	dcxzr, dcxzrerr := io.ReadAll(xzr)
	if dcxzrerr != nil {
		return nil, dcxzrerr
	}

	var grsa []structure.GitHubRepositoryStructure

	if uerr := json.Unmarshal(dcxzr, &grsa); uerr != nil {
		return nil, uerr
	}

	return grsa, nil
}
