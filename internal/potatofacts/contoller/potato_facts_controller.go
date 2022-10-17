/*
 * Copyright 2022 Armory, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package contoller

import (
	"context"
	"github.com/armory-io/go-commons/server"
	"github.com/armory-io/go-commons/server/serr"
	"github.com/armory-io/potato-facts/internal/potatofacts/service"
	"go.uber.org/zap"
	"net/http"
)

func NewPotatoFactController(
	log *zap.SugaredLogger,
	p4o *service.PotatoFactsSvc,
) server.Controller {
	return server.Controller{
		Controller: &PotatoFactsController{
			log: log,
			p4o: p4o,
		},
	}
}

type PotatoFactsController struct {
	log *zap.SugaredLogger
	p4o *service.PotatoFactsSvc
}

func (p *PotatoFactsController) Handlers() []server.Handler {
	return []server.Handler{
		server.NewHandler(p.factsHandler, server.HandlerConfig{
			Method:     http.MethodGet,
			Path:       "/fact",
			AuthOptOut: true,
		}),
	}
}

type PotatoFact struct {
	Fact string `json:"fact"`
}

func (p *PotatoFactsController) factsHandler(ctx context.Context, _ server.Void) (*server.Response[PotatoFact], serr.Error) {
	return server.SimpleResponse(PotatoFact{Fact: p.p4o.GetFact()}), nil
}
