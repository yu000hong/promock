package main

//Config 配置信息
type Config struct {
	HTTPServerAddress string
	RedisAddress      string
	RedisPassword     string
	MysqlDSN          string
}

//ReadConfig 从配置文件中读取配置信息
func ReadConfig(path string) *Config {
	config := new(Config)
	k2v := ReadConfigFile(path)

	config.HTTPServerAddress = GetString(k2v, "http_server_address")
	config.RedisAddress = GetString(k2v, "redis_address")
	config.RedisPassword = GetOptString(k2v, "redis_password", "")
	config.MysqlDSN = GetString(k2v, "mysql_datasource")
	return config
}


// GetInt 从配置MAP中获取int类型配置
func GetInt(configs map[string]string, key string) int64 {
	value, ok := configs[key]
	if !ok {
		log.Fatalf("key:%s non exist", key)
	}
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatalf("key:%s is't integer", key)
	}
	return n
}

// GetOptInt 从配置MAP中获取int类型配置
func GetOptInt(configs map[string]string, key string, defaultValue int64) int64 {
	value, ok := configs[key]
	if !ok {
		return defaultValue
	}
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatalf("key:%s is't integer", key)
	}
	return n
}

// GetString 从配置MAP中获取string类型配置
func GetString(configs map[string]string, key string) string {
	value, ok := configs[key]
	if !ok {
		log.Fatalf("key:%s non exist", key)
	}
	return value
}

// GetOptString 从配置MAP中获取string类型配置
func GetOptString(configs map[string]string, key string, defaultValue string) string {
	value, ok := configs[key]
	if !ok {
		return defaultValue
	}
	return value
}

// GetStringArray 从配置MAP中获取string slice类型配置
func GetStringArray(configs map[string]string, key string) []string {
	value, ok := configs[key]
	if !ok {
		log.Fatalf("key:%s non exist", key)
	}
	if value == "" {
		return []string{}
	}
	fields := strings.Split(value, ",")
	var array = make([]string, 0, len(fields))
	for _, ele := range fields {
		if ele != "" {
			array = append(array, ele)
		}
	}
	return array
}

// GetOptStringArray 从配置MAP中获取string slice类型配置
func GetOptStringArray(configs map[string]string, key string, defaultValue []string) []string {
	value, ok := configs[key]
	if !ok {
		return defaultValue
	}
	if value == "" {
		return []string{}
	}
	fields := strings.Split(value, ",")
	var array = make([]string, 0, len(fields))
	for _, ele := range fields {
		if ele != "" {
			array = append(array, ele)
		}
	}
	return array
}

// ReadConfigFile 从文件中读取配置MAP
func ReadConfigFile(path string) map[string]string {
	configs := make(map[string]string)
	err := cfg.Load(path, configs)
	if err != nil {
		log.Fatal(err)
	}
	return configs
}
