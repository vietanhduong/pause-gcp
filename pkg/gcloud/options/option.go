package options

type Options struct {
	CredentialsFilePath string
	BackupStatePath     string
}

type Option func(*Options)

func WithCredentialsPath(path string) Option {
	return func(options *Options) { options.CredentialsFilePath = path }
}

func WithBackupStatePath(path string) Option {
	return func(options *Options) { options.BackupStatePath = path }
}
