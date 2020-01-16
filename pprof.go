/*
 * PPROF WARPER
 * 		Switch Environment PPROF_ENABLED = Y
 * 			PPROF_CPU_PROFILE
 *			PPROF_MEM_PROFILE
 *			PPROF_BLK_PROFILE
 */

package pprof

import (
	"runtime"
	"runtime/pprof"
	"os"
	"fmt"
	"path/filepath"
)

var pprof_enabled = false
var cpuProfile string = ""
var memProfile string = ""
var memProfileRate = 10
var blkProfile string = ""
var blkProfileRate = 10


func init() {
	var env_pprof_enabled = os.Getenv("PPROF_ENABLED")
	if env_pprof_enabled == "Y" {
		pprof_enabled = true

		curdir, _ := os.Getwd()
		// CPU Profile Related
		var env_pprof_cpu_profile = os.Getenv("PPROF_CPU_PROFILE")
		if env_pprof_cpu_profile != nil && env_pprof_cpu_profile != "" {
			cpuProfile = env_pprof_cpu_profile
		} else {
			cpuProfile = filepath.Join(curdir, "cpu.prof")
		}

		// MEMORY Profile Related
		var env_pprof_mem_profile = os.Getenv("PPROF_MEM_PROFILE")
		if env_pprof_mem_profile != nil && env_pprof_mem_profile != "" {
			memProfile = env_pprof_mem_profile
		} else {
			memProfile = filepath.Join(curdir, "mem.prof")
		}

		// BLOCK Profile Related
		var env_pprof_blk_profile = os.Getenv("PPROF_BLK_PROFILE")
		if env_pprof_blk_profile != nil && env_pprof_blk_profile != "" {
			blkProfile = env_pprof_blk_profile
		} else {
			blkProfile = filepath.Join(curdir, "blk.prof")
		}
	}
}

func StartCPUProfile() {
	if pprof_enabled {
		f, err := os.Create(cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",
				err)
			return
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
			f.Close()
			return
		}
	}
}

func StopCPUProfile() {
	if pprof_enabled {
		pprof.StopCPUProfile()
	}
}

func StartMemProfile() {
	if pprof_enabled {
		runtime.MemProfileRate = memProfileRate
	}
}

func StopMemProfile() {
	if pprof_enabled {
		f, err := os.Create(memProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create mem profile output file: %s", err)
			return
		}
		if err = pprof.WriteHeapProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not write %s: %s", memProfile, err)
		}
		f.Close()
	}
}

func startBlkProfile() {
	if pprof_enabled {
		runtime.SetBlockProfileRate(blkProfileRate)
	}
}

func stopBlkProfile() {
	if pprof_enabled {
		f, err := os.Create(blkProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create block profile output file: %s", err)
			return
		}
		if err = pprof.Lookup("block").WriteTo(f, 0); err != nil {
			fmt.Fprintf(os.Stderr, "Can not write %s: %s", blkProfile, err)
		}
		f.Close()
	}
}
