// Chapter 18: Memory Management & GC Tuning
package main

import (
    "log"
    "runtime"
    "runtime/debug"
    "time"
)

func init() {
    debug.SetGCPercent(150) // tune GC target
}

func main() {
    var m runtime.MemStats
    for {
        runtime.ReadMemStats(&m)
        log.Printf("HeapAlloc=%d MiB, NumGC=%d
", m.HeapAlloc/1<<20, m.NumGC)
        time.Sleep(10 * time.Second)
    }
}
