package elf

import "debug/elf"

type File struct {
	*elf.File
}

func NewElfFile(f *elf.File) *File {
	return &File{f}
}

func (f *File) Close() error {
	return f.Close()
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
