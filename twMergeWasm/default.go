package twMergeWasm

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"strings"

	"github.com/will-lol/merge"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)


type twMerge struct {
	Runtime wazero.Runtime
	Context context.Context
	Reader io.Reader
	Writer io.Writer
	WasmBytes *[]byte
}

//go:embed lib/index.wasm
var twMergeWasm []byte

func NewTwMerge() (merge.TwMerge, error) {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	return &twMerge{
		Runtime: r,
		Context: ctx,
		WasmBytes: &twMergeWasm,
	}, nil
}

func (m twMerge) Merge(existing string, incoming string) (*string, error) {
	reader := strings.NewReader(strings.Join([]string{incoming, existing}, " "))
	var writer bytes.Buffer
	_, err := m.Runtime.InstantiateWithConfig(m.Context, *m.WasmBytes, wazero.NewModuleConfig().WithStdin(reader).WithStdout(&writer))
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(&writer)
	string := string(bytes)
	return &string, nil
}

func (m twMerge) Close() {
	m.Runtime.Close(m.Context)
}
