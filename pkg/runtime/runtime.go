package runtime

import (
	"github.com/gosolid/solid/packages/runtime"
	"github.com/gosolid/solid/pkg/runtime/assert"
	"github.com/gosolid/solid/pkg/runtime/buffer"
	"github.com/gosolid/solid/pkg/runtime/console"
	"github.com/gosolid/solid/pkg/runtime/crypto"
	crypto_web "github.com/gosolid/solid/pkg/runtime/crypto/web"
	"github.com/gosolid/solid/pkg/runtime/events"
	"github.com/gosolid/solid/pkg/runtime/fs"
	"github.com/gosolid/solid/pkg/runtime/http"
	"github.com/gosolid/solid/pkg/runtime/net"
	"github.com/gosolid/solid/pkg/runtime/os"
	"github.com/gosolid/solid/pkg/runtime/repl"
	"github.com/gosolid/solid/pkg/runtime/stream"
	"github.com/gosolid/solid/pkg/runtime/tty"
	"github.com/gosolid/solid/pkg/runtime/url"
	"github.com/gosolid/solid/pkg/runtime/util"
	"github.com/gosolid/solid/pkg/runtime/webrtc"
	"github.com/gosolid/solid/pkg/runtime/worker_threads"
	"github.com/grexie/isolates"
)

type Import interface{}

var _ crypto_web.Import
var _ assert.Import
var _ buffer.Buffer
var _ crypto.Hash
var _ console.Console
var _ fs.FS
var _ http.Http
var _ events.EventEmitter
var _ stream.Stream
var _ net.Net
var _ fs.FS
var _ tty.ReadStreamBase
var _ util.Import
var _ events.EventHandler
var _ repl.Import
var _ os.Import
var _ worker_threads.Import
var _ webrtc.PeerConnection
var _ url.Import

func AddRuntime() error {
	f := fs.NewFS()

	f.Mount("/", fs.NewEmbedFS(&runtime.FS), "/lib")

	// isolates.RegisterRuntimeLibrary("@grexie/workers/console", fs, "/console.js")
	isolates.RegisterRuntimeLibrary("crypto", f, "/crypto.js")
	// isolates.RegisterRuntimeLibrary("os", fs, "/os.js")
	isolates.RegisterRuntimeLibrary("path", f, "/path.js")
	isolates.RegisterRuntimeLibrary("@grexie/workers/stream", f, "/stream/index.js")
	isolates.RegisterRuntimeLibrary("@grexie/workers/text", f, "/text.js")
	// isolates.RegisterRuntimeLibrary("tty", fs, "/tty.js")
	isolates.RegisterRuntimeLibrary("zlib", f, "/zlib.js")
	isolates.RegisterRuntimeLibrary("native:@grexie/workers/global", f, "/global.js")

	return nil
}

var _ = AddRuntime()
