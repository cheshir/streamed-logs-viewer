package app

import (
	"sync/atomic"
)

const (
	ModeRegexp     = 1
	ModeMultiline  = 2
	ModeIgnoreCase = 4
	ModeJSON       = 8
	ModeWrap       = 16
)

type Mode uint32

func NewMode() *Mode {
	return new(Mode)
}

func (m *Mode) Regexp() bool {
	return m.load()&ModeRegexp > 0
}

func (m *Mode) EnableRegexp() {
	v := m.load()
	v |= ModeRegexp
	m.set(v)
}

func (m *Mode) DisableRegexp() {
	v := m.load()
	v ^= ModeRegexp | ModeMultiline | ModeIgnoreCase
	m.set(v)
}

func (m *Mode) Multiline() bool {
	return m.load()&ModeMultiline > 0
}

func (m *Mode) EnableMultiline() {
	v := m.load()
	v |= ModeRegexp | ModeMultiline
	m.set(v)
}

func (m *Mode) DisableMultiline() {
	v := m.load()
	v ^= ModeMultiline
	m.set(v)
}

func (m *Mode) IgnoreCase() bool {
	return m.load()&ModeIgnoreCase > 0
}

func (m *Mode) EnableIgnoreCase() {
	v := m.load()
	v |= ModeRegexp | ModeIgnoreCase
	m.set(v)
}

func (m *Mode) DisableIgnoreCase() {
	v := m.load()
	v ^= ModeIgnoreCase
	m.set(v)
}

func (m *Mode) JSON() bool {
	return m.load()&ModeJSON > 0
}

func (m *Mode) EnableJSON() {
	v := m.load()
	v |= ModeJSON
	m.set(v)
}

func (m *Mode) DisableJSON() {
	v := m.load()
	v ^= ModeJSON
	m.set(v)
}

func (m *Mode) Wrap() bool {
	return m.load()&ModeWrap > 0
}

func (m *Mode) EnableWrap() {
	v := m.load()
	v |= ModeWrap
	m.set(v)
}

func (m *Mode) DisableWrap() {
	v := m.load()
	v ^= ModeWrap
	m.set(v)
}

func (m *Mode) load() uint32 {
	v := uint32(*m)
	return atomic.LoadUint32(&v)
}

func (m *Mode) set(v uint32) {
	var i interface{} = m
	atomic.StoreUint32(i.(*uint32), v)
}
