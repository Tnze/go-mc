package main

import (
	"encoding/base64"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/matryer/is"
)

type mapFSWithMkdir struct {
	files fstest.MapFS
}

func (m mapFSWithMkdir) Open(name string) (fs.File, error) {
	return m.files.Open(name)
}

func (m mapFSWithMkdir) OpenFile(name string, flag int, perm fs.FileMode) (fs.File, error) {
	createMode := flag & os.O_CREATE
	truncMode := flag & os.O_TRUNC
	if createMode != 0 {
		m.files[name] = &fstest.MapFile{
			Mode: perm,
		}
	}

	if _, ok := m.files[name]; !ok {
		return nil, os.ErrNotExist
	}

	if truncMode != 0 {
		m.files[name].Data = nil
	}

	return m.files.Open(name)
}

func (m mapFSWithMkdir) Mkdir(name string, perm fs.FileMode) error {
	m.files[name] = &fstest.MapFile{Mode: fs.ModeDir | perm}
	return nil
}

func (m mapFSWithMkdir) ReadDir(name string) ([]fs.DirEntry, error) {
	return m.files.ReadDir(name)
}

func (m mapFSWithMkdir) Stat(name string) (fs.FileInfo, error) {
	return m.files.Stat(name)
}
func (m mapFSWithMkdir) WriteFile(name string, data []byte, perm fs.FileMode) error {
	m.files[name] = &fstest.MapFile{Data: data}
	return nil
}

func buildMockFS() mapFSWithMkdir {
	return mapFSWithMkdir{
		files: fstest.MapFS{},
	}
}

const manifestData = `{"latest":{"release":"1.20.1","snapshot":"1.20.1"},"versions":[{"id":"1.20.1","type":"release","url":"https://piston-meta.mojang.com/v1/packages/715ccf3330885e75b205124f09f8712542cbe7e0/1.20.1.json","time":"2023-06-12T13:32:21+00:00","releaseTime":"2023-06-12T13:25:51+00:00","sha1":"715ccf3330885e75b205124f09f8712542cbe7e0","complianceLevel":1},{"id":"1.20","type":"release","url":"https://piston-meta.mojang.com/v1/packages/52f6c28f40ee907d167a1f217d7a48cbec4936c5/1.20.json","time":"2023-06-12T10:37:38+00:00","releaseTime":"2023-06-02T08:36:17+00:00","sha1":"52f6c28f40ee907d167a1f217d7a48cbec4936c5","complianceLevel":1},{"id":"1.19.4","type":"release","url":"https://piston-meta.mojang.com/v1/packages/a4118bd311bc49c9ca298284c0055f25a007e4f8/1.19.4.json","time":"2023-06-12T10:18:09+00:00","releaseTime":"2023-03-14T12:56:18+00:00","sha1":"a4118bd311bc49c9ca298284c0055f25a007e4f8","complianceLevel":1}]}`
const versionData = `base64:eyJhc3NldEluZGV4Ijp7ImlkIjoiNSIsInNoYTEiOiI5ZDU4ZmRkMjUzOGM2ODc3ZmI1YzVjNTU4ZWJjNjBlZTBiNmQwZTg0Iiwic2l6ZSI6NDExNTgxLCJ0b3RhbFNpemUiOjYxNzcxODc5OSwidXJsIjoiaHR0cHM6Ly9waXN0b24tbWV0YS5tb2phbmcuY29tL3YxL3BhY2thZ2VzLzlkNThmZGQyNTM4YzY4NzdmYjVjNWM1NThlYmM2MGVlMGI2ZDBlODQvNS5qc29uIn0sIm1haW5DbGFzcyI6Im5ldC5taW5lY3JhZnQuY2xpZW50Lm1haW4uTWFpbiIsIm1pbmltdW1MYXVuY2hlclZlcnNpb24iOjIxLCJyZWxlYXNlVGltZSI6IjIwMjMtMDYtMTJUMTM6MjU6NTErMDA6MDAiLCJ0aW1lIjoiMjAyMy0wNi0xMlQxMzoyNTo1MSswMDowMCIsInR5cGUiOiJyZWxlYXNlIn0K`
const versionHashesData = `base64:eyJvYmplY3RzIjp7Im1pbmVjcmFmdC9sYW5nL2ZpbF9waC5qc29uIjp7Imhhc2giOiI3MTJhMjM2Nzk0MzEzZTUxMmU0N2UxYzAwMGE5ZmNkMmNjMjQ0ZWRjIiwic2l6ZSI6NDMxODg3fX19Cg==`
const languageData = `base64:eyJibG9jay5taW5lY3JhZnQuYWNhY2lhX2J1dHRvbiI6IkFrYXN5YW5nIFBpbmR1dGFuIiwiYmxvY2subWluZWNyYWZ0LmFjYWNpYV9kb29yIjoiQWthc3lhbmcgUGludG8iLCJibG9jay5taW5lY3JhZnQuYWNhY2lhX2ZlbmNlIjoiQWthc3lhbmcgQmFrb2QiLCJibG9jay5taW5lY3JhZnQuYWNhY2lhX2ZlbmNlX2dhdGUiOiJBa2FzeWFuZyBUYXJhbmdrYWhhbiIsImJsb2NrLm1pbmVjcmFmdC5hY2FjaWFfaGFuZ2luZ19zaWduIjoiTmFrYXNhYml0IG5hIEthcmF0dWxhbmcgQWthc3lhIiwiYmxvY2subWluZWNyYWZ0LmFjYWNpYV9sZWF2ZXMiOiJEYWhvbmcgQWthc3lhIiwiYmxvY2subWluZWNyYWZ0LmFjYWNpYV9sb2ciOiJBa2FzeWFuZyBUcm9zbyIsImJsb2NrLm1pbmVjcmFmdC5hY2FjaWFfcGxhbmtzIjoiQWthc3lhbmcgVGFibGEiLCJibG9jay5taW5lY3JhZnQuYWNhY2lhX3ByZXNzdXJlX3BsYXRlIjoiQWthc3lhbmcgQXBha2FuIiwiYmxvY2subWluZWNyYWZ0LmFjYWNpYV9zYXBsaW5nIjoiSGFsYW1hbmcgQWthc3lhIiwiYmxvY2subWluZWNyYWZ0LmFjYWNpYV9zaWduIjoiQWthc3lhbmcgS2FyYXR1bGEiLCJibG9jay5taW5lY3JhZnQuYWNhY2lhX3NsYWIiOiJBa2FzeWFuZyBUaWxhZCIsImJsb2NrLm1pbmVjcmFmdC5hY2FjaWFfc3RhaXJzIjoiQWthc3lhbmcgSGFnZGFuYW4iLCJibG9jay5taW5lY3JhZnQuYWNhY2lhX3RyYXBkb29yIjoiTWFsaWl0IG5hIEFrYXN5YW5nIFBpbnRvIiwiYmxvY2subWluZWNyYWZ0LmFjYWNpYV93YWxsX2hhbmdpbmdfc2lnbiI6Ik5ha2FzYWJpdCBuYSBLYXJhdHVsYW5nIEFrYXN5YSIsImJsb2NrLm1pbmVjcmFmdC5hY2FjaWFfd2FsbF9zaWduIjoiQWthc3lhbmcgS2FyYXR1bGEgc2EgUGFkZXIiLCJibG9jay5taW5lY3JhZnQuYWNhY2lhX3dvb2QiOiJBa2FzeWFuZyBLYWhveSIsImJsb2NrLm1pbmVjcmFmdC5hY3RpdmF0b3JfcmFpbCI6IlRhZ2EtYnVrYXMgbmEgUmlsZXMiLCJibG9jay5taW5lY3JhZnQuYWlyIjoiSGltcGFwYXdpZCIsImJsb2NrLm1pbmVjcmFmdC5hbGxpdW0iOiJBbGxpdW0iLCJibG9jay5taW5lY3JhZnQuYW1ldGh5c3RfYmxvY2siOiJBbWV0aXN0YW5nIEJsb2tlIiwiYmxvY2subWluZWNyYWZ0LmFtZXRoeXN0X2NsdXN0ZXIiOiJLdW1wb2wgbmcgQW1ldGlzdGEiLCJibG9jay5taW5lY3JhZnQuYW5jaWVudF9kZWJyaXMiOiJTaW5hdW5hbmcgWWFnaXQiLCJibG9jay5taW5lY3JhZnQuYW5kZXNpdGUiOiJBbmRlc2F5dCIsImJsb2NrLm1pbmVjcmFmdC5hbmRlc2l0ZV9zbGFiIjoiQW5kZXNheXQgbmEgVGlsYWQiLCJibG9jay5taW5lY3JhZnQuYW5kZXNpdGVfc3RhaXJzIjoiQW5kZXNheXQgbmEgSGFnZGFuYW4iLCJibG9jay5taW5lY3JhZnQuYW5kZXNpdGVfd2FsbCI6IlBhZGVyIG5hIEFuZGVzYXl0IiwiYmxvY2subWluZWNyYWZ0LmFudmlsIjoiUGFsaWhhbiIsImJsb2NrLm1pbmVjcmFmdC5hdHRhY2hlZF9tZWxvbl9zdGVtIjoiTmFrYWthYml0IG5hIFRhbmdrYXkgbmcgTWVsb24iLCJibG9jay5taW5lY3JhZnQuYXR0YWNoZWRfcHVtcGtpbl9zdGVtIjoiTmFrYWthYml0IG5hIFRhbmdrYXkgbmcgS2FsYWJhc2EiLCJibG9jay5taW5lY3JhZnQuYXphbGVhIjoiQXphbGV5YSIsImJsb2NrLm1pbmVjcmFmdC5hemFsZWFfbGVhdmVzIjoiRGFob25nIEF6YWxleWEiLCJibG9jay5taW5lY3JhZnQuYXp1cmVfYmx1ZXQiOiJBc3VsIG5hIExpZ2F3IG5hIEJ1bGFrbGFrIn0K`

func buildMockHTTPGet() func(url string) (*http.Response, error) {
	urlsToData := map[string]string{
		"https://piston-meta.mojang.com/mc/game/version_manifest_v2.json":                                 manifestData,
		"https://piston-meta.mojang.com/v1/packages/715ccf3330885e75b205124f09f8712542cbe7e0/1.20.1.json": versionData,
		"https://piston-meta.mojang.com/v1/packages/9d58fdd2538c6877fb5c5c558ebc60ee0b6d0e84/5.json":      versionHashesData,
		"https://resources.download.minecraft.net/71/712a236794313e512e47e1c000a9fcd2cc244edc":            languageData,
	}

	return func(url string) (*http.Response, error) {
		resp := http.Response{}
		v, ok := urlsToData[url]
		if !ok {
			resp.StatusCode = http.StatusNotFound
			resp.Body = io.NopCloser(strings.NewReader(""))
			return &resp, nil
		}
		content, err := decodeBase64(v)
		if err != nil {
			resp.StatusCode = http.StatusInternalServerError
			resp.Body = io.NopCloser(strings.NewReader(err.Error()))
			return &resp, nil
		}
		resp.Body = io.NopCloser(strings.NewReader(content))
		resp.StatusCode = http.StatusOK
		return &resp, nil
	}
}

func decodeBase64(str string) (string, error) {
	if strings.HasPrefix(str, "base64:") {
		decoded, err := base64.StdEncoding.DecodeString(str[7:])
		if err != nil {
			return "", err
		}
		return string(decoded), nil
	}
	return str, nil
}

func TestRunWithNoArgs(t *testing.T) {
	is := is.New(t)

	mockFS := buildMockFS()

	is.NoErr(run(mockFS, buildMockHTTPGet(), []string{}))

	langDir, ok := mockFS.files["fil-ph"]
	is.True(ok) // did not create language parent directory
	is.True(langDir.Mode.IsDir())

	_, ok = mockFS.files["fil-ph/fil_ph.go"]
	is.True(ok) // did not generate Go src from language data
}
