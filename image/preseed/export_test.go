// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2019-2022 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package preseed

import (
	"github.com/snapcore/snapd/seed"
)

var (
	SystemSnapFromSeed       = systemSnapFromSeed
	ChooseTargetSnapdVersion = chooseTargetSnapdVersion
	CreatePreseedArtifact    = createPreseedArtifact
)

type PreseedOpts = preseedOpts

func MockSeedOpen(f func(rootDir, label string) (seed.Seed, error)) (restore func()) {
	oldSeedOpen := seedOpen
	seedOpen = f
	return func() {
		seedOpen = oldSeedOpen
	}
}

func SnapdPathAndVersion(targetSnapd *targetSnapdInfo) (string, string) {
	return targetSnapd.path, targetSnapd.version
}
