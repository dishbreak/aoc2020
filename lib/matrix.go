package lib

import (
	"image"
	"strings"
)

type Matrix struct {
	d [][]byte
	n int
}

func NewMatrixWithoutFrame(s []string) *Matrix {
	m := &Matrix{
		d: make([][]byte, len(s)-2),
		n: len(s) - 2,
	}

	for i, line := range s[1 : len(s)-1] {
		m.d[i] = []byte(line[1 : len(s)-1])
	}
	return m
}

func NewMatrix(s []string) *Matrix {
	m := &Matrix{
		d: make([][]byte, len(s)),
		n: len(s),
	}

	for i, line := range s {
		m.d[i] = []byte(line)
	}

	return m
}

func NewMatrixFromBytes(b [][]byte) *Matrix {
	m := &Matrix{
		d: make([][]byte, len(b)),
		n: len(b),
	}

	for i, row := range b {
		m.d[i] = make([]byte, len(row))
		copy(m.d[i], row)
	}

	return m
}

func (m *Matrix) GetBytes() [][]byte {
	d := make([][]byte, len(m.d))
	for i, r := range m.d {
		d[i] = make([]byte, len(m.d[i]))
		copy(d[i], r)
	}
	return d
}

func (m *Matrix) GetDim() int {
	return m.n
}

func (m *Matrix) swap(one, other image.Point) {
	m.d[one.Y][one.X], m.d[other.Y][other.X] = m.d[other.Y][other.X], m.d[one.Y][one.X]
}

func (m *Matrix) Rotate() {
	for l := 0; l < len(m.d[0]); l++ {
		nw := image.Point{l, l}
		ne := image.Point{m.n - l - 1, l}
		se := image.Point{m.n - l - 1, m.n - l - 1}
		sw := image.Point{l, m.n - l - 1}
		for i := l; i < m.n-l-1; i++ {
			m.swap(nw, sw)
			m.swap(sw, se)
			m.swap(se, ne)
			nw = nw.Add(image.Point{1, 0})
			sw = sw.Add(image.Point{0, -1})
			se = se.Add(image.Point{-1, 0})
			ne = ne.Add(image.Point{0, 1})
		}
	}
}

func (m *Matrix) FlipHorizontal() {
	var swap []byte
	for i := 0; i < len(m.d); i++ {
		swap = m.d[i]
		m.d[i] = m.d[m.n-i-1]
		m.d[m.n-i-1] = swap
	}
}

func (m *Matrix) FlipVertical() {
	for _, row := range m.d {
		for i := 0; i < len(row)/2; i++ {
			swap := row[i]
			row[i] = row[m.n-i-1]
			row[m.n-i-1] = swap
		}
	}
}

func (m *Matrix) String() string {
	b := strings.Builder{}
	for _, row := range m.d {
		for _, col := range row {
			b.WriteByte(col)
		}
		b.WriteString("\n")
	}
	return b.String()
}
