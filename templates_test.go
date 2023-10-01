// Copyright Nicolas Paul (2023)
//
// * Nicolas Paul
//
// This software is a computer program whose purpose is to allow the hosting
// and sharing of Go modules using a personal domain.
//
// This software is governed by the CeCILL license under French law and
// abiding by the rules of distribution of free software.  You can  use,
// modify and/ or redistribute the software under the terms of the CeCILL
// license as circulated by CEA, CNRS and INRIA at the following URL
// "http://www.cecill.info".
//
// As a counterpart to the access to the source code and  rights to copy,
// modify and redistribute granted by the license, users are provided only
// with a limited warranty  and the software's author,  the holder of the
// economic rights,  and the successive licensors  have only  limited
// liability.
//
// In this respect, the user's attention is drawn to the risks associated
// with loading,  using,  modifying and/or developing or reproducing the
// software by the user in light of its specific status of free software,
// that may mean  that it is complicated to manipulate,  and  that  also
// therefore means  that it is reserved for developers  and  experienced
// professionals having in-depth computer knowledge. Users are therefore
// encouraged to load and test the software's suitability as regards their
// requirements in conditions enabling the security of their systems and/or
// data to be ensured and,  more generally, to use and operate it in the
// same conditions as regards security.
//
// The fact that you are presently reading this means that you have had
// knowledge of the CeCILL license and that you accept its terms.

package main

import (
	"bytes"
	"io"
	"testing"
)

func TestExecIndex(t *testing.T) {
	type args struct {
		url   string
		count uint8
	}
	tests := []struct {
		name    string
		args    args
		want    *bytes.Buffer
		wantErr bool
	}{
		{
			name: "valid url waiting 5 seconds",
			args: args{
				url:   "https://example.com",
				count: 5,
			},
			want: bytes.NewBufferString(`<!DOCTYPE html>
<meta charset="UTF-8">
<meta http-equiv="refresh" content="5; url='https://example.com'">
<meta name="viewport" content="width=device-width,initial-scale=1">
<p>Nothing to see, redirecting <a href="https://example.com">here</a>.`),
			wantErr: false,
		},
		{
			name: "valid url waiting 0 second",
			args: args{
				url:   "https://example.com",
				count: 0,
			},
			want: bytes.NewBufferString(`<!DOCTYPE html>
<meta charset="UTF-8">
<meta http-equiv="refresh" content="0; url='https://example.com'">
<meta name="viewport" content="width=device-width,initial-scale=1">
<p>Nothing to see, redirecting <a href="https://example.com">here</a>.`),
			wantErr: false,
		},
		{
			name: "not the same url",
			args: args{
				url:   "https://example.com",
				count: 0,
			},
			want: bytes.NewBufferString(`<!DOCTYPE html>
<meta charset="UTF-8">
<meta http-equiv="refresh" content="0; url='https://example2.com'">
<meta name="viewport" content="width=device-width,initial-scale=1">
<p>Nothing to see, redirecting <a href="https://example.com">here</a>.`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := ExecIndex(&buf, tt.args.url, tt.args.count)

			if !tt.wantErr && err != nil {
				t.Errorf("ExecIndex() error = %v", err)
			}

			if !tt.wantErr && !bytes.Equal(buf.Bytes(), tt.want.Bytes()) {
				t.Errorf("ExecIndex() = %v, want %v", buf, tt.want)
			}
		})
	}
}

func TestExecModule(t *testing.T) {
	type args struct {
		w      *io.Writer
		prefix string
		vcs    string
		home   string
		dir    string
		file   string
	}
	tests := []struct {
		name    string
		args    args
		want    *bytes.Buffer
		wantErr bool
	}{
		{
			name: "module example.com/foo hosted on github",
			args: args{
				prefix: "example.com/foo",
				vcs:    "git",
				home:   "https://github.com/example/foo",
				dir:    "https://github.com/example/foo/tree/master{/dir}",
				file:   "https://github.com/example/foo/blob/master{/dir}/{file}#L{line}",
			},
			want: bytes.NewBufferString(`<!DOCTYPE html>
<meta charset="UTF-8">
<meta name="go-import" content="example.com/foo git https://github.com/example/foo">
<meta name="go-source" content="example.com/foo https://github.com/example/foo https://github.com/example/foo/tree/master{/dir} https://github.com/example/foo/blob/master{/dir}/{file}#L{line}">
<meta name="viewport" content="width=device-width,initial-scale=1">
<p>There is nothing to see, redirecting <a href="https://pkg.go.dev/example.com/foo">here</a>.`),
			wantErr: false,
		},
		{
			name: "module example.com/foo hosted on gitlab",
			args: args{
				prefix: "example.com/foo",
				vcs:    "git",
				home:   "https://gitlab.com/example/foo",
				dir:    "https://gitlab.com/example/foo/-/tree/master{/dir}",
				file:   "https://gitlab.com/example/foo/-/blob/master{/dir}/{file}#L{line}",
			},
			want: bytes.NewBufferString(`<!DOCTYPE html>
<meta charset="UTF-8">
<meta name="go-import" content="example.com/foo git https://gitlab.com/example/foo">
<meta name="go-source" content="example.com/foo https://gitlab.com/example/foo https://gitlab.com/example/foo/-/tree/master{/dir} https://gitlab.com/example/foo/-/blob/master{/dir}/{file}#L{line}">
<meta name="viewport" content="width=device-width,initial-scale=1">
<p>There is nothing to see, redirecting <a href="https://pkg.go.dev/example.com/foo">here</a>.`),
			wantErr: false,
		},
		{
			name: "module example.com/foo hosted on Source Hut with Mercurial",
			args: args{
				prefix: "example.com/foo",
				vcs:    "hg",
				home:   "https://hg.sr.ht/~example/foo",
				dir:    "https://hg.sr.ht/~example/foo/tree/master{/dir}",
				file:   "https://hg.sr.ht/~example/foo/browse/master{/dir}/{file}#{line}",
			},
			want: bytes.NewBufferString(`<!DOCTYPE html>
<meta charset="UTF-8">
<meta name="go-import" content="example.com/foo hg https://hg.sr.ht/~example/foo">
<meta name="go-source" content="example.com/foo https://hg.sr.ht/~example/foo https://hg.sr.ht/~example/foo/tree/master{/dir} https://hg.sr.ht/~example/foo/browse/master{/dir}/{file}#{line}">
<meta name="viewport" content="width=device-width,initial-scale=1">
<p>There is nothing to see, redirecting <a href="https://pkg.go.dev/example.com/foo">here</a>.`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := ExecModule(&buf, tt.args.prefix, tt.args.vcs, tt.args.home,
				tt.args.dir, tt.args.file)

			if !tt.wantErr && err != nil {
				t.Errorf("ExecIndex() error = %v", err)
			}

			if !tt.wantErr && !bytes.Equal(buf.Bytes(), tt.want.Bytes()) {
				t.Errorf("ExecIndex() = %v, want %v", buf, tt.want)
			}
		})
	}
}
