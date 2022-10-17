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

package service

import (
	"github.com/armory-io/potato-facts/internal/potatofacts/config"
	"go.uber.org/zap"
	"math/rand"
)

type PotatoFactsSvc struct {
	logger *zap.SugaredLogger
	facts  *[]string
}

func PotatoFactsSvcProvider(
	logger *zap.SugaredLogger,
	potatoFactsConf config.PotatoFactsConfiguration,
) *PotatoFactsSvc {
	logger.Infof("Loaded %d facts", len(potatoFactsConf.Facts))

	return &PotatoFactsSvc{
		logger: logger,
		facts:  &potatoFactsConf.Facts,
	}
}

func (p *PotatoFactsSvc) GetFact() string {
	return (*p.facts)[rand.Intn(len(*p.facts))]
}
