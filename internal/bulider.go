package internal

// @Author: Derek
// @Description:
// @Date: 2023/4/15 22:47
// @Version 1.0

type Dog interface {
	Do()
}

type Builder interface {
	Build() Dog
}

var gl = &wrapSelector{}

var _ Builder = (*wrapSelector)(nil)

type wrapSelector struct {
	Builder
}

func Get() Builder {
	if gl.Builder == nil {
		return nil
	}
	return gl
}

func Set(builder Builder) {
	gl.Builder = builder
	return
}
