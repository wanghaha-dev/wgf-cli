package main

import (
	"fmt"
	"github.com/wanghaha-dev/gf/frame/g"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if exists, err := checkLsExists("gox"); err != nil || !exists {
		fmt.Println("exec => go install github.com/mitchellh/gox@latest")
		Exec2("go", "install", "github.com/mitchellh/gox@latest")
	}

	if exists, err := checkLsExists("air"); err != nil || !exists {
		fmt.Println("exec => go install github.com/air-verse/air@latest")
		Exec2("go", "install", "github.com/air-verse/air@latest")
	}

	args := os.Args

	if len(args) == 1 {
		fmt.Println("Please choose build or run or init.")
		return
	}

	arg := args[1]
	switch arg {
	case "build":
		build()
	case "run":
		Exec2("air")
	case "init":
		initTpl()
	}
}

func build() {
	name := g.Cfg().GetString("gfcli.build.name")
	output := g.Cfg().GetString("gfcli.build.output")
	version := g.Cfg().GetString("gfcli.build.version")
	arch := g.Cfg().GetString("gfcli.build.arch")
	system := g.Cfg().GetString("gfcli.build.system")

	if !strings.Contains(system, ",") {
		if system == "all" {
			system = "linux,darwin,windows"
		}
	}
	systemSlice := strings.Split(system, ",")

	if !strings.Contains(arch, ",") {
		if arch == "all" {
			arch = "amd64,arm"
		}
	}
	archSlice := strings.Split(arch, ",")
	var outSlice []string

	for _, systemItem := range systemSlice {

		for _, archItem := range archSlice {
			if systemItem == "darwin" && archItem == "arm" {
				continue
			}
			if systemItem == "windows" && archItem == "arm" {
				continue
			}

			out := fmt.Sprintf("%v/%v/%v_%v/%v", output, version, systemItem, archItem, name)
			osarch := fmt.Sprintf("%v/%v", systemItem, archItem)

			outSlice = append(outSlice, out)

			// gox -osarch="linux/amd64" -output=bin/v1.0.0/linux_amd64/new-upload-gf
			fmt.Println("gox", "-osarch", osarch, "-output", out)
			Exec("gox", "-osarch", osarch, "-output", out)
		}

	}

	fmt.Println("==========> done!")
	for _, item := range outSlice {
		fmt.Println(item)
	}
	// go build -o ~/go/bin/wgf .
}

func Exec(name string, args ...string) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s\n", string(out))
}

// Exec2 直接终端输出
func Exec2(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func initTpl() {
	Exec2("git", "clone", "https://gitee.com/develop1024/gfproject-tpl.git")
}

func checkLsExists(cmd string) (bool, error) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		return false, err
	}

	fmt.Println(cmd, "Found =>", path)
	return true, nil
}
