// 代码生成时间: 2025-10-11 23:02:46
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/net"
)

// SystemMonitor 结构体用于存储系统监控的相关数据
type SystemMonitor struct {
    Context context.Context
}

// NewSystemMonitor 创建一个新的系统监控器
func NewSystemMonitor(ctx context.Context) *SystemMonitor {
    return &SystemMonitor{
        Context: ctx,
    }
}

// MonitorCpu 监控CPU使用情况
func (sm *SystemMonitor) MonitorCpu() (*cpu.TimesStat, error) {
    cpuTimes, err := cpu.Times(false)
    if err != nil {
        log.Printf("Error monitoring CPU: %v", err)
        return nil, err
    }
    return cpuTimes[0], nil
}

// MonitorMemory 监控内存使用情况
func (sm *SystemMonitor) MonitorMemory() (*mem.VirtualMemoryStat, error) {
    memStat, err := mem.VirtualMemory()
    if err != nil {
        log.Printf("Error monitoring memory: %v", err)
        return nil, err
    }
    return memStat, nil
}

// MonitorDisks 监控磁盘使用情况
func (sm *SystemMonitor) MonitorDisks() ([]*disk.PartitionStat, error) {
    partitions, err := disk.Partitions(false)
    if err != nil {
        log.Printf("Error monitoring disks: %v", err)
        return nil, err
    }
    return partitions, nil
}

// MonitorNetwork 监控网络使用情况
func (sm *SystemMonitor) MonitorNetwork() ([]net.IOCountersStat, error) {
    netStats, err := net.IOCounters(true)
    if err != nil {
        log.Printf("Error monitoring network: %v", err)
        return nil, err
    }
    return netStats, nil
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    monitor := NewSystemMonitor(ctx)
    for {
        // 监控CPU
        cpuStat, err := monitor.MonitorCpu()
        if err != nil {
            fmt.Println("Failed to monitor CPU: ", err)
            continue
        }
        fmt.Printf("CPU Usage: %v
", cpuStat.User)

        // 监控内存
        memStat, err := monitor.MonitorMemory()
        if err != nil {
            fmt.Println("Failed to monitor memory: ", err)
            continue
        }
        fmt.Printf("Memory Used: %.2f%%
", memStat.UsedPercent)

        // 监控磁盘
        partitions, err := monitor.MonitorDisks()
        if err != nil {
            fmt.Println("Failed to monitor disks: ", err)
            continue
        }
        for _, partition := range partitions {
            fmt.Printf("Disk: %s Used: %.2f%%
", partition.Device, partition.UsedPercent)
        }

        // 监控网络
        netStats, err := monitor.MonitorNetwork()
        if err != nil {
            fmt.Println("Failed to monitor network: ", err)
            continue
        }
        for _, stat := range netStats {
            fmt.Printf("Network Interface: %s - Sent: %d bytes, Received: %d bytes
", stat.Name, stat.BytesSent, stat.BytesRecv)
        }

        // 每秒检查一次
        time.Sleep(1 * time.Second)
    }
}
