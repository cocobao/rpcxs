package share

import (
	"context"
	"fmt"
	"reflect"
)

// Context is a rpcx customized Context that can contains multiple values.
type Context struct {
	tags map[interface{}]interface{}
	context.Context
}

func NewContext(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
		tags:    make(map[interface{}]interface{}),
	}
}

func (c *Context) Value(key interface{}) interface{} {
	if v, ok := c.tags[key]; ok {
		return v
	}
	return c.Context.Value(key)
}

func (c *Context) SetValue(key, val interface{}) {
	c.tags[key] = val
}

func (c *Context) String() string {
	return fmt.Sprintf("%v.WithValue(%v)", c.Context, c.tags)
}

func WithValue(parent context.Context, key, val interface{}) *Context {
	if key == nil {
		panic("nil key")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}

	tags := make(map[interface{}]interface{})
	tags[key] = val
	return &Context{Context: parent, tags: tags}
}

func WithLocalValue(ctx *Context, key, val interface{}) *Context {
	if key == nil {
		panic("nil key")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}

	ctx.tags[key] = val
	return ctx
}
