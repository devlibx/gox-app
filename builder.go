package gox_app

import "github.com/harishb2k/gox-base/lock"

type RouterConfigBuilder struct {
	routerConfig RouteConfig
}

func NewRouterConfigBuilder(name string) *RouterConfigBuilder {
	return &RouterConfigBuilder{RouteConfig{Name: name}}
}

func (r *RouterConfigBuilder) WithPath(path string) *RouterConfigBuilder {
	r.routerConfig.Path = path
	return r
}

func (r *RouterConfigBuilder) WithConsumeType(consumeType ContentType) *RouterConfigBuilder {
	r.routerConfig.Consumes = append(r.routerConfig.Consumes, consumeType)
	return r
}

func (r *RouterConfigBuilder) WithProduceType(produceType ContentType) *RouterConfigBuilder {
	r.routerConfig.Produces = append(r.routerConfig.Produces, produceType)
	return r
}

func (r *RouterConfigBuilder) WithHandlerFunc(handlerFunc HandlerFunc) *RouterConfigBuilder {
	r.routerConfig.HandlerFunc = handlerFunc
	return r
}

func (r *RouterConfigBuilder) WithRequestParser(parser RequestParser) *RouterConfigBuilder {
	r.routerConfig.RequestParser = parser
	return r
}

func (r *RouterConfigBuilder) WithResponseBuilder(builder ResponseBuilder) *RouterConfigBuilder {
	r.routerConfig.ResponseBuilder = builder
	return r
}

func (r *RouterConfigBuilder) WithLock(lock lock.DistributedLock) *RouterConfigBuilder {
	r.routerConfig.Lock = lock
	return r
}

func (r *RouterConfigBuilder) Build() RouteConfig {
	if len(r.routerConfig.Consumes) == 0 {
		r.routerConfig.Consumes = append(r.routerConfig.Consumes, ContentTypeJson)
	}
	if len(r.routerConfig.Produces) == 0 {
		r.routerConfig.Produces = append(r.routerConfig.Produces, ContentTypeJson)
	}
	return r.routerConfig
}
