// Copyright 2017. All rights reserved.
// This file is part of cleaner project
// Created by duguying on 2017/10/1.

package cleaner

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	fc := New("/Users/duguying/com.yfcloud.opsapp.debug 2017-10-01 14:52.30.964.xcappdata/AppData", time.Second*24, time.Minute, -1)
	fc.StartCleanTask()
	time.Sleep(time.Minute * 5)
}
