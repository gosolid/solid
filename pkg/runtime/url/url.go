//js:package url

package url

import (
	"context"
	"net/url"

	"github.com/grexie/isolates"
)

type Import interface{}

//go:generate go run github.com/grexie/isolates/codegen

type URL struct {
	Hash     string `v8:"hash"`
	Host     string `v8:"host"`
	Pathname string `v8:"pathname"`
	Password string `v8:"password"`
	Protocol string `v8:"protocol"`
	Search   string `v8:"search"`
	Username string `v8:"username"`
}

//js:export URL
//js:constructor URL
func NewURL(in isolates.FunctionArgs) (*URL, error) {
	urlString := in.Arg(in.Context.GetIsolate().GetExecutionContext(), 0).String()

	if urlParsed, err := url.Parse(urlString); err != nil {
		return nil, err
	} else {
		u := &URL{
			Hash:     urlParsed.Fragment,
			Host:     urlParsed.Host,
			Search:   urlParsed.RawQuery,
			Pathname: urlParsed.Path,
			Protocol: urlParsed.Scheme + ":",
		}
		return u, nil
	}
}

//js:method toString
func (u *URL) String() string {
	_u := &url.URL{
		RawFragment: u.Hash,
		Host:        u.Host,
		RawQuery:    u.Search,
		Path:        u.Pathname,
		Scheme:      u.Protocol[0 : len(u.Protocol)-1],
	}
	return _u.String()
}

//js:export pathToFileURL
//js:method
func PathToFileURL(ctx context.Context, path string) (*URL, error) {
	if urlv, err := isolates.For(ctx).New(NewURL, "file://"+path); err != nil {
		return nil, err
	} else {
		return urlv.(*URL), nil
	}
}
