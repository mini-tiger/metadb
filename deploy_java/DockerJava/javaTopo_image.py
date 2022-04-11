# coding:utf-8
import os
import subprocess


class JavaImage():
    def __init__(self, *args, **kwargs):

        self.codeDir = kwargs["codeDir"]
        self.mvnScript = kwargs["shell_file"]
        self.dockerImage = kwargs["imageTag"]

    def mvn_command(self, capture_output=True):
        self.out = subprocess.run([os.path.join(self.codeDir, self.mvnScript), self.codeDir],
                                  capture_output=capture_output)

    def DockerBuild(self, capture_output=True):
        # os.chdir(self.codeDir)
        command = "/usr/bin/docker build -t %s -f Dockerfile . " % self.dockerImage
        print(command)
        # self.out = subprocess.run(["docker","version"], capture_output=capture_output)
        self.out = subprocess.run(command, shell=True, encoding="utf-8", capture_output=capture_output, timeout=30)

    def DockerPush(self, capture_output=True):
        command = "/usr/bin/docker push %s " % self.dockerImage
        print(command)
        self.out = subprocess.run(command, shell=True, encoding="utf-8", capture_output=capture_output)

    def begin(self):
        self.mvn_command(False)
        if self.out.returncode == 0:
            print("jar mvn success")
        else:
            print("jar mvn fail")
            print(self.out.stderr.__str__())

        self.DockerBuild(False)
        if self.out.returncode == 0:
            print("docker build success")
        else:
            print("docker build fail")
            print(self.out.stderr.__str__())

        self.DockerPush(False)
        if self.out.returncode == 0:
            print("docker push success")
        else:
            print("docker push fail")
            print(self.out.stderr.__str__())


if __name__ == "__main__":
    j = JavaImage(codeDir=os.path.abspath(os.curdir), shell_file="javapack.sh",
                  imageTag="harbor.dev.21vianet.com/cmdb/cmdb_javatopo:latest")
    j.begin()
    # print()
