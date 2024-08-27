package pkg

import (
	"errors"
	"os/exec"
	"strings"
)

func ListEnvPath() ([]string, error) {
	res := make([]string, 0)
	cmd := exec.Command("powershell", "-Command", "[System.Environment]::GetEnvironmentVariable('PATH', 'User')")
	output, err := cmd.Output()
	if err != nil {
		return res, err
	}

	res = strings.Split(string(output), ";")
	return res, nil
}

func SetPhpEnv(path string) error {
	if path == "" {
		return errors.New("path error")
	}

	phpPath, err := ReadConfig()
	if err != nil {
		return err
	}

	// 移除原有的php环境变量
	envPath, err := ListEnvPath()
	if err != nil {
		return err
	}

	for k, v := range envPath {
		if strings.Contains(v, phpPath) || v == "" {
			envPath = append(envPath[:k], envPath[k+1:]...)
		}
	}

	// 设置新的PATH环境变量
	envPath = append(envPath, path)
	newPath := strings.Join(envPath, ";")
	// 使用 PowerShell 设置用户级 PATH 环境变量
	return exec.Command("powershell", "-Command",
		"[System.Environment]::SetEnvironmentVariable('PATH', \""+newPath+"\", 'User')").Run()
}
