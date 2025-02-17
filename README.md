# config-mgtm
config-mgtm is a wrapper on top of viper

## How to use

`go get github.com/csturiale/config-mgtm`

Have a file yaml under assets/application/configuration/ folder named **configuration.yaml**
Then can be used as follow
```golang
configuration.GetBool("my.path")
```

# Per env use

Define the env environment variable like:

`export env=prod` 

And create a file named **configuration-prod.yaml**

This will be merged with the base configuration.yaml

