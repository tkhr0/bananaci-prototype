package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// FilePath is リポジトリ内の BananaCI の設定ファイルパス
const FilePath string = ".bananaci/config.yml"

// Yaml を読み込む
func UnmarshalFromFile(yamlPath string) (*Config, error) {
	content, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read the file: %s\n", yamlPath)
	}

	return Unmarshal(content)
}

func Unmarshal(in []byte) (*Config, error) {
	config := &Config{}

	err := yaml.Unmarshal(in, config)
	if err != nil {
		return nil, fmt.Errorf("%s cannot unmarshal data: %s\n", err, string(in))
	}

	return config, nil
}
