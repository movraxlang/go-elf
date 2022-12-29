package elf

import (
	"debug/elf"
	"io"
	"os"
)

type File struct {
	*elf.File
	io.ReaderAt
}

func NewFile(name string) (*File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	f, err := elf.NewFile(file)
	if err != nil {
		return nil, err
	}
	return &File{f, file}, nil
}

func (f *File) FindSymbolAddress(symbol string) (uint64, error) {
	syms, err := f.Symbols()
	if err != nil {
		return 0, err
	}

	for _, sym := range syms {
		if sym.Name == symbol {
			return sym.Value, nil
		}
	}
	return 0, nil
}
