package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

const (
	streamingJar   = "src/hadoop-streaming-2.7.3.jar"
	pythonMapper   = "src/word_count_mapper.py"
	pythonReducer  = "src/word_count_reducer.py"
	secondMapper   = "src/term_mapper.py"
	secondReducer  = "src/term_reducer.py"
	sparkJob       = "src/word_count_spark.py"
	outputPath     = "./output"
	sparkSubmitJar = "target/scala-2.11/simple-project_2.11-1.0.jar"

	hadoopCmd    = "./src/hadoop_command.sh %s"
	sparkCommand = "./src/spark_command.sh %s"
)

func RunSpark(filepath string) string {
	cmd := MakeSparkSubprocess(filepath)

	errStream, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
	}
	outStream, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Start()

	bytes, err := ioutil.ReadAll(errStream)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
	bytes, err = ioutil.ReadAll(outStream)
	if err != nil {
		fmt.Println(err)
	}
	out := string(bytes)
	cmd.Wait()

	return string(out)
}

func RunHadoop(filepath string) string {
	cmd := MakeHadoopSubprocess(filepath)

	errStream, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
	}
	outStream, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Start()

	bytes, err := ioutil.ReadAll(errStream)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
	bytes, err = ioutil.ReadAll(outStream)
	if err != nil {
		fmt.Println(err)
	}
	out := string(bytes)
	cmd.Wait()

	return string(out)
}
func HadoopCommand(filepath string) []string {
	return strings.Split(fmt.Sprintf(hadoopCmd, filepath), " ")
}

func SparkCommand(filepath string) []string {
	return strings.Split(fmt.Sprintf(sparkCommand, filepath), " ")
}

func MakeSparkSubprocess(filepath string) *exec.Cmd {
	cmdArgs := SparkCommand(filepath)
	fmt.Println(cmdArgs)
	return exec.Command(cmdArgs[0], cmdArgs[1:]...)
}

func MakeHadoopSubprocess(filepath string) *exec.Cmd {
	cmdArgs := HadoopCommand(filepath)
	fmt.Println(cmdArgs)
	fmt.Println(len(cmdArgs))
	return exec.Command(cmdArgs[0], cmdArgs...)
}
