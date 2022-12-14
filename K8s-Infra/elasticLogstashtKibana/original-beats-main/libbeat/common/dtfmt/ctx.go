// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package dtfmt

import (
	"time"
)

// ctx stores pre-computed time fields used by the formatter.
type ctx struct {
	year             int
	month            time.Month
	day              int
	weekday          time.Weekday
	yearday          int
	isoWeek, isoYear int

	hour, min, sec int
	nano           int

	tzOffset int

	buf []byte
}

type ctxConfig struct {
	date     bool
	clock    bool
	weekday  bool
	yearday  bool
	nano     bool
	iso      bool
	tzOffset bool
}

func (c *ctx) initTime(config *ctxConfig, t time.Time) {
	if config.date {
		c.year, c.month, c.day = t.Date()
	}
	if config.clock {
		c.hour, c.min, c.sec = t.Clock()
	}
	if config.iso {
		c.isoYear, c.isoWeek = t.ISOWeek()
	}

	if config.nano {
		c.nano = t.Nanosecond()
	}

	if config.yearday {
		c.yearday = t.YearDay()
	}

	if config.weekday {
		c.weekday = t.Weekday()
	}

	if config.tzOffset {
		_, c.tzOffset = t.Zone()
	}
}

func (c *ctxConfig) enableDate() {
	c.date = true
}

func (c *ctxConfig) enableClock() {
	c.clock = true
}

func (c *ctxConfig) enableNano() {
	c.nano = true
}

func (c *ctxConfig) enableWeekday() {
	c.weekday = true
}

func (c *ctxConfig) enableYearday() {
	c.yearday = true
}

func (c *ctxConfig) enableISO() {
	c.iso = true
}

func (c *ctxConfig) enableTimeZoneOffset() {
	c.tzOffset = true
}
