package elf

import "debug/elf"

type File struct {
	*elf.File
}

func NewElfFile(name string) (*File, error) {
	f, err := elf.Open(name)
	if err != nil {
		return nil, err
	}
	return &File{f}, nil
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
