package logger

type ZapConfig struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`                // 软链接名称
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}

var LogConfig *ZapConfig
var Level string

// NewDefaultZapConfig 返回一个默认的配置
func NewDefaultZapConfig() *ZapConfig {
	return &ZapConfig{
		Director:      "logs",
		Level:         Level,
		ShowLine:      true,
		StacktraceKey: "stacktrace",
		EncodeLevel:   "LowercaseColorLevelEncoder",
		LinkName:      "latest_log",
		LogInConsole:  true,
		Prefix:        "[ logger ]",
		Format:        "console",
	}
}

/*
	# zap logger configuration
	zap:
	  level: 'info'
	  format: 'console'
	  prefix: '[log]'
	  director: 'logs'
	  link-name: 'latest_log'
	  show-line: true
	  encode-level: 'LowercaseColorLevelEncoder'
	  stacktrace-key: 'stacktrace'
	  log-in-console: true
*/
