package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"
)

var (
	currentFileName, rootDir, dockerDir, webCompressionPath, tmpDir, binaryDir, helm_path string // images.go 路径

	branch = "metadb-tj" //#"3.9.39.x"

	//validDir  = map[string]int{"cmdb_adminserver": 60004, "cmdb_webserver": 8090,
	//	"cmdb_apiserver": 8080, "cmdb_coreservice": 50009, "cmdb_toposerver": 60002, "cmdb_hostserver": 60001,
	//}

	// dir: src/bin/build/
	validDir = map[string]struct{}{"cmdb_adminserver": struct{}{}, "cmdb_webserver": struct{}{},
		"cmdb_apiserver": struct{}{}, "cmdb_coreservice": struct{}{}, "cmdb_toposerver": struct{}{},
		"cmdb_hostserver": struct{}{}, "cmdb_operationserver": struct{}{}, "cmdb_cacheservice": struct{}{},
		"cmdb_cloudserver": struct{}{}, "cmdb_eventserver": struct{}{}, "cmdb_procserver": struct{}{},
		"cmdb_taskserver": struct{}{},
	}

	//dockerTmp = "Dockerfile_tmp"

	harbor             = "harbor.dev.21vianet.com/cmdb/"
	github             = "meta42/"
	ver                = "latest"
	kubeconfig         = "/smb/50.25kubeconfig"
	helm_ns            = "cmdbv3"
	helm_uninstall_cmd = fmt.Sprintf("helm --kubeconfig=%s uninstall -n %s cmdb", kubeconfig, helm_ns)
	helm_install_cmd   = fmt.Sprintf("helm --kubeconfig=%s install -n %s cmdb -f values.yaml .", kubeconfig, helm_ns)
	helm_upgrade_cmd   = fmt.Sprintf("helm --kubeconfig=%s upgrade -n %s cmdb --history-max 3 -f values.yaml .", kubeconfig, helm_ns)
	t1                 = time.Now()
	version            = fmt.Sprintf("%d-%d-%d_%d%d%d", t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second())
)

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func init() {
	_, currentFileName, _, _ = runtime.Caller(0)
	//fmt.Println(CurrentFileName)
	rootDir = filepath.Dir(filepath.Dir(currentFileName))

	dockerDir = path.Join(rootDir, "DockerFile")
	webCompressionPath = path.Join(dockerDir, "web.tar.gz")
	tmpDir = path.Join(dockerDir, "tmp")
	binaryDir = path.Join(rootDir, "src", "bin", "build", branch)

	log.Printf("rootDir:%v\n", rootDir)
	log.Printf("binaryDir:%v\n", binaryDir)
	helm_path = filepath.Join(rootDir, "deploy", "cmdb-helm")
	if !IsDir(tmpDir) {
		os.MkdirAll(tmpDir, os.ModePerm)
	}
	err := RunCommand("docker pull harbor.dev.21vianet.com/taojun/python2.7-debug-tz:latest")
	if err != nil {
		log.Fatalln(err)
	}

}

func RunCommand(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	var e bytes.Buffer
	cmd.Stderr = &e
	//cmd.Start()
	//buf, _ := cmd.CombinedOutput()
	cmd.Run()
	if e.Len() != 0 || strings.Contains(e.String(), "fail") || strings.Contains(e.String(), "err") {
		//return e.String(), out.String()
		return errors.New(fmt.Sprintf("cmd: %s err:%s", command, e.String()))
	}

	return nil
}

func RunCommandStd(command string) (bytes.Buffer, bytes.Buffer) {
	cmd := exec.Command("/bin/bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	var e bytes.Buffer
	cmd.Stderr = &e
	//cmd.Start()
	//buf, _ := cmd.CombinedOutput()
	cmd.Run()
	//if e.Len() != 0 || strings.Contains(e.String(), "fail") || strings.Contains(e.String(), "err") {
	//	//return e.String(), out.String()
	//	return errors.New(fmt.Sprintf("cmd: %s err:%s", command, e.String()))
	//}

	return out, e
}

func main() {

	var err error
	listDir, _ := ioutil.ReadDir(binaryDir)
	var srcDir, destDir string

	for _, subDir := range listDir {
		//fmt.Println(subDir)
		if subDir.IsDir() {
			_, ok := validDir[subDir.Name()]
			if ok {
				srcDir = path.Join(binaryDir, subDir.Name())
				destDir = path.Join(tmpDir, "cmdb")
				log.Printf("===============  %s  =================\n", srcDir)
				log.Printf("copy %s %s\n", srcDir, destDir)
				err = binaryFileDirCopy(path.Join(binaryDir, subDir.Name()), destDir, true)
				if err != nil {
					log.Fatalln(err)
				}
				if strings.Contains(subDir.Name(), "adminserver") {
					log.Printf("copy %s %s\n", "init.py", destDir)
					err = binaryFileDirCopy(path.Join(dockerDir, "init.py"), path.Join(destDir, "init.py"), false)
					if err != nil {
						log.Fatalln(err)
					}
				} else {
					log.Printf("copy %s %s\n", "init_skipadminserver.py", destDir)
					err = binaryFileDirCopy(path.Join(dockerDir, "init_skipadminserver.py"), path.Join(destDir, "init.py"), false)
					if err != nil {
						log.Fatalln(err)
					}
				}
				if strings.Contains(subDir.Name(), "webserver") {
					log.Printf("copy %s %s\n", "web.tar.gz", destDir)
					err = binaryFileDirCopy(webCompressionPath, path.Join(tmpDir, "web.tar.gz"), false)
					if err != nil {
						log.Fatalln(err)
					}
				}

				sv := TplVariables{
					//Port:        port,
					AdminServer: false,
					WebServer:   false,
					AppName:     subDir.Name(),
					HarborUri:   "",
					ImageTag:    version,
				}
				_ = sv.Entry()
			}
		}
	}

	// xxx Deploy k8s dev
	//log.Println(helm_uninstall_cmd)

	//std, stderr := RunCommandStd(helm_uninstall_cmd)
	//log.Printf("std:%v stderr:%v\n", std.String(), stderr.String())

	// helm values.yaml.tpl
	err = createFile(path.Join(helm_path, "values.yaml"), path.Join(helm_path, "values.yaml.tpl"),
		map[string]interface{}{"version": version})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("chdir %s\n", helm_path)
	err = os.Chdir(helm_path)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(helm_upgrade_cmd)
	std, stderr := RunCommandStd(helm_upgrade_cmd)
	log.Printf("std:%v stderr:%v\n", std.String(), stderr.String())
	//log.Println("sleep 10s")
	//time.Sleep(10 * time.Second)
	//log.Println(helm_install_cmd)

	//std, stderr = RunCommandStd(helm_install_cmd)
	//log.Printf("std:%v stderr:%v\n", std.String(), stderr.String())

}
func binaryFileDirCopy(src, dest string, dir bool) error {
	var cmdstr string
	if dir {
		if err := cleanDest(dest); err != nil {
			return err
		}
		cmdstr = fmt.Sprintf("cp -rf %s %s", src, dest)
	} else {
		if err := cleanDest(dest); err != nil {
			return err
		}
		cmdstr = fmt.Sprintf("cp -f %s %s", src, dest)
	}

	return RunCommand(cmdstr)

}
func cleanDest(dest string) error {
	var cmdstr string
	if IsDir(dest) {
		cmdstr = fmt.Sprintf("rm -rf %s/*", dest)
	} else {
		cmdstr = fmt.Sprintf("rm -f %s", dest)
	}

	return RunCommand(cmdstr)
}

type TplVariables struct {
	//Port        int
	AdminServer bool
	WebServer   bool
	//ExtraCommand1 string
	//ExtraCommand2 string
	AppName   string
	HarborUri string
	ImageTag  string
}

func (sv *TplVariables) Entry() (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Dir: %s fail\n", sv.AppName)
		}
	}()
	log.Println("generate run.sh")
	err = sv.generateShell()
	if err != nil {
		log.Println(err)
	}
	log.Println("generate dockerfile")
	err = sv.generateDockerFile()
	if err != nil {
		log.Println(err)
	}
	log.Println("generate dockerImage")
	err = sv.generateDockerImage()
	if err != nil {
		log.Println(err)
	}
	log.Println("push dockerImage To harbor")
	err = sv.pushDockerImage()
	if err != nil {
		log.Println(err)
	}
	log.Println("Local Clean dockerImage")
	err = sv.cleanDockerImage()
	if err != nil {
		log.Println(err)
	}
	//log.Println("push dockerImage To dockerhub")
	//err = sv.pushDockerHubImage()
	//if err != nil {
	//	log.Println(err)
	//}
	return nil
}

func (t *TplVariables) generateDockerImage() error {
	var cmdstr string
	t.HarborUri = fmt.Sprintf("%s%s:%s", harbor, t.AppName, t.ImageTag)
	//log.Printf("%s\n", t.HarborUri)
	cmdstr = fmt.Sprintf("pushd %s && docker build --no-cache -f Dockerfile -t %s . && popd", tmpDir, t.HarborUri)

	return RunCommand(cmdstr)
}
func (t *TplVariables) pushDockerImage() error {
	var cmdstr string
	//log.Printf("%s\n", t.HarborUri)
	cmdstr = fmt.Sprintf("docker push %s ", t.HarborUri)

	return RunCommand(cmdstr)
}

func (t *TplVariables) cleanDockerImage() error {
	var cmdstr string
	//log.Printf("%s\n", t.HarborUri)
	cmdstr = fmt.Sprintf("docker rmi %s ", t.HarborUri)

	return RunCommand(cmdstr)
}

func (t *TplVariables) pushDockerHubImage() error {
	var cmdstr string
	githubUri := fmt.Sprintf("%s%s:%s", github, t.AppName, ver)

	//log.Printf("docker tag %s %s\n", t.HarborUri, githubUri)

	_ = RunCommand(fmt.Sprintf("docker tag %s %s", t.HarborUri, githubUri))

	log.Printf("%s \n", githubUri)

	cmdstr = fmt.Sprintf("docker push %s ", githubUri)

	return RunCommand(cmdstr)
}

func (t *TplVariables) generateShell() error {

	switch true {
	case strings.Contains(t.AppName, "adminserver"):
		t.AdminServer = true
	}
	return createFile(path.Join(tmpDir, "run.sh"), path.Join(dockerDir, "run.sh.tpl"), t)
}

func (t *TplVariables) generateDockerFile() error {

	switch true {
	case strings.Contains(t.AppName, "webserver"):

		t.WebServer = true

	}
	return createFile(path.Join(tmpDir, "Dockerfile"), path.Join(dockerDir, "Dockerfile.tpl"), t)
}

func createFile(file, tpl string, p interface{}) error {
	t, err := template.ParseFiles(tpl) //parse the html file homepage.html
	if err != nil {                    // if there is an error
		//log.Print("template parsing error: ", err) // log it
		return err
	}
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755) //读写模式
	if err != nil {
		//log.Println(err)
		return err
		//os.Exit(1)
	}

	if err != nil {
		//log.Println(err)
		return err
	}
	//p := DockerFileVariables{extraCommand: ""}
	//if err := t.Execute(os.Stdout, p); err != nil {
	//	fmt.Println("There was an error:", err.Error())
	//}
	if err := t.Execute(f, p); err != nil {
		return err
	}
	return nil

}
