// 代码生成时间: 2025-09-18 09:31:24
package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "os/exec"
    "strings"
)

// NetworkStatusChecker 结构体用于检查网络连接状态
type NetworkStatusChecker struct {
    host string
    port int
}

// NewNetworkStatusChecker 创建一个新的NetworkStatusChecker实例
func NewNetworkStatusChecker(host string, port int) *NetworkStatusChecker {
    return &NetworkStatusChecker{
        host: host,
        port: port,
    }
}

// CheckConnection 检查指定的主机和端口是否可达
func (nsc *NetworkStatusChecker) CheckConnection() (bool, error) {
    // 构建网络地址
    address := fmt.Sprintf("%s:%d", nsc.host, nsc.host)

    // 尝试建立TCP连接
    conn, err := net.DialTimeout("tcp", address, 5e9)
    if err != nil {
        return false, err
    }
    defer conn.Close()

    // 如果连接成功，则返回true
    return true, nil
}

// IsReachable 检查网络地址是否可达
func IsReachable(address string) (bool, error) {
    // 尝试解析地址
    _, err := net.ResolveIPAddr("ip", address)
    if err != nil {
        return false, err
    }

    // 尝试建立TCP连接
    conn, err := net.DialTimeout("tcp", address, 5e9)
    if err != nil {
        return false, err
    }
    defer conn.Close()

    // 如果连接成功，则返回true
    return true, nil
}

// RunPing 发送ping命令检查主机是否可达
func RunPing(host string) (bool, error) {
    // 构建ping命令
    cmd := exec.Command("ping", "-c", "1", host)

    // 执行ping命令
    if err := cmd.Run(); err != nil {
        return false, err
    }

    // 如果ping命令成功执行，则返回true
    return true, nil
}

func main() {
    // 示例：检查Google的端口80是否可达
    nsc := NewNetworkStatusChecker("www.google.com", 80)
    reachable, err := nsc.CheckConnection()
    if err != nil {
        fmt.Printf("Error checking connection: %v
", err)
    } else {
        fmt.Printf("Host %s is %sreachable.
", nsc.host, func() string {
            if reachable {
                return ""
            }
            return "not "
        }())
    }

    // 示例：检查Google是否可达
    reachable, err := IsReachable("www.google.com:80")
    if err != nil {
        fmt.Printf("Error checking reachability: %v
", err)
    } else {
        fmt.Printf("Host www.google.com is %sreachable.
", func() string {
            if reachable {
                return ""
            }
            return "not "
        }())
    }

    // 示例：使用ping命令检查Google是否可达
    reachable, err := RunPing("www.google.com")
    if err != nil {
        fmt.Printf("Error pinging host: %v
", err)
    } else {
        fmt.Printf("Host www.google.com is %sreachable.
", func() string {
            if reachable {
                return ""
            }
            return "not "
        }())
    }
}