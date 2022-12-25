package elf

import "debug/elf"

type File struct {
	f *elf.File
}

func NewElfFile(f *elf.File) *File {
	return &File{f: f}
}

func (f *File) FindSymbolAddress(symbol string) (uint64, error) {
	syms, err := f.f.Symbols()
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
