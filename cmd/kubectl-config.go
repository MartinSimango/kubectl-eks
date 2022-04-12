package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

var homeDir, _ = os.UserHomeDir()
var configFolder = homeDir + "/.kube/eks"
var configFile = homeDir + "/.kube/eks/config"

type KubectlEKSConfig interface {
	AddContext(name, clusterName, profile, region string, currentContext bool) error
	DeleteContext(name string) error
	UseContext(name string) error
	Sync() error
	GetConfig() *KubectlEKSConfigImpl
	DoesContextExist(name string) (bool, int)
	GetCurrentContext() *Context
}

type Context struct {
	Name        string `yaml:"name"`
	ClusterName string `yaml:"cluster-name"`
	Region      string `yaml:"region,omitempty"`
	Profile     string `yaml:"profile"`
}

func NewContext(name, clusterName, region, profile string) *Context {
	return &Context{
		Name:        name,
		ClusterName: clusterName,
		Region:      region,
		Profile:     profile,
	}
}

type KubectlEKSConfigImpl struct {
	Contexts       []Context `yaml:"contexts"`
	CurrentContext string    `yaml:"current-context"`
}

func NewKubectlEKSConfigImpl() *KubectlEKSConfigImpl {
	config := &KubectlEKSConfigImpl{}
	return config.GetConfig()
}

// Check we implement the interface
var _ KubectlEKSConfig = &KubectlEKSConfigImpl{}

func (ke *KubectlEKSConfigImpl) AddContext(name, clusterName, profile, region string, currentContext bool) error {
	contextExists, _ := ke.DoesContextExist(name)
	if contextExists {
		log.Fatal("Context with name already exists")
	}

	context := NewContext(name, clusterName, region, profile)

	config := ke.GetConfig()
	config.Contexts = append(config.Contexts, *context)

	if len(config.Contexts) == 1 {
		config.CurrentContext = name
	} else if currentContext {
		config.CurrentContext = name
	}

	writeToConfigFile(config)
	if currentContext || len(config.Contexts) == 1 {
		fmt.Printf("Context '%s' successfully added and set as current context\n", name)
	} else {
		fmt.Printf("Context '%s' successfully added\n", name)
	}
	return nil
}

func (ke *KubectlEKSConfigImpl) DeleteContext(name string) error {
	contextExists, index := ke.DoesContextExist(name)
	if !contextExists {
		log.Fatal("Context with name doesn't exist")
	}

	config := ke.GetConfig()

	copy(config.Contexts[index:], config.Contexts[index+1:])
	config.Contexts = config.Contexts[:len(config.Contexts)-1]

	if config.CurrentContext == name {
		config.CurrentContext = ""
	}

	writeToConfigFile(config)
	fmt.Printf("Context %s removed from config file\n", name)

	return nil
}

func (ke *KubectlEKSConfigImpl) UseContext(name string) error {
	contextExists, _ := ke.DoesContextExist(name)
	if !contextExists {
		log.Fatal("Context with name doesnt exist")
	}

	config := ke.GetConfig()
	config.CurrentContext = name

	writeToConfigFile(config)
	return nil

}

func (ke *KubectlEKSConfigImpl) Sync() error {
	context := ke.GetCurrentContext()
	argArray := []string{"eks", "update-kubeconfig", "--name", context.ClusterName, "--profile", context.Profile}

	if context.Region != "" {
		argArray = append(argArray, "--region")
		argArray = append(argArray, context.Region)
	}
	output, err := exec.Command("aws", argArray...).CombinedOutput()
	if err != nil {
		log.Fatal(string(output[:]))
	}

	fmt.Println(string(output[:]))
	return nil
}

func (ke *KubectlEKSConfigImpl) GetConfig() *KubectlEKSConfigImpl {
	checkConfigFile()

	yfile, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Fatal(err)
	}

	data := &KubectlEKSConfigImpl{}

	err = yaml.Unmarshal(yfile, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (ke *KubectlEKSConfigImpl) DoesContextExist(name string) (bool, int) {
	kubectlConfig := ke.GetConfig()
	for index, context := range kubectlConfig.Contexts {
		if context.Name == name {
			return true, index
		}
	}
	return false, -1

}

func (ke *KubectlEKSConfigImpl) GetCurrentContext() *Context {
	if ke.CurrentContext == "" {
		log.Fatal("Current context is currenty not set")
	}
	contextExists, index := ke.DoesContextExist(ke.CurrentContext)
	if !contextExists {
		log.Fatal("No context found with name ", ke.CurrentContext)
	}
	return &ke.Contexts[index]
}

func checkConfigFile() {
	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		createConfigFolder()

	} else if _, err := os.Stat(configFile); os.IsNotExist(err) {
		writeToConfigFile(NewKubectlEKSConfigImpl())
	}
}

func createConfigFolder() {
	err := os.MkdirAll(configFolder, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
	writeToConfigFile(NewKubectlEKSConfigImpl())

}

func writeToConfigFile(config *KubectlEKSConfigImpl) {
	yamlData, err := yaml.Marshal(config)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	err = ioutil.WriteFile(configFile, yamlData, 0666)

	if err != nil {

		log.Fatal(err)
	}
}
