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
	systems := g.Cfg().GetString("gfcli.build.system")

	systemSlice := strings.Split(systems, ",")

	for _, item := range systemSlice {
		out := fmt.Sprintf("%v/%v/%v_%v/%v", output, version, item, arch, name)

		osarch := fmt.Sprintf("%v/%v", item, arch)

		// gox -osarch="linux/amd64" -output=bin/v1.0.0/linux_amd64/new-upload-gf
		Exec("gox", "-osarch", osarch, "-output", out)
	}

	fmt.Println("==========> done!")
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
