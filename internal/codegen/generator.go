package codegen

import (
	"context"
	"sync/atomic"
)

type SequentialGenerator struct {
	next atomic.Uint64
}

func NewSequentialGenerator() *SequentialGenerator {
	g := &SequentialGenerator{}
	g.next.Store(MinimumGeneratedID())
	return g
}

func (g *SequentialGenerator) Generate(context.Context) (string, error) {
	id := g.next.Add(1) - 1
	return EncodeBase62(id), nil
}
