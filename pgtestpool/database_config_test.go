package pgtestpool

import "testing"

// Using table driven tests as described here: https://github.com/golang/go/wiki/TableDrivenTests#parallel-testing
func TestDatabaseConfigConnectionString(t *testing.T) {
	t.Parallel() // marks table driven test execution function as capable of running in parallel with other tests

	tests := []struct {
		name   string
		config DatabaseConfig
		want   string
	}{
		{
			name: "Simple",
			config: DatabaseConfig{
				Host:     "localhost",
				Port:     5432,
				Username: "simple",
				Password: "database_config",
				Database: "simple_database_config",
			},
			want: "host=localhost port=5432 user=simple password=database_config dbname=simple_database_config sslmode=disable",
		},
		{
			name: "SSLMode",
			config: DatabaseConfig{
				Host:     "localhost",
				Port:     5432,
				Username: "simple",
				Password: "database_config",
				Database: "simple_database_config",
				AdditionalParams: map[string]string{
					"sslmode": "prefer",
				},
			},
			want: "host=localhost port=5432 user=simple password=database_config dbname=simple_database_config sslmode=prefer",
		},
		{
			name: "Complex",
			config: DatabaseConfig{
				Host:     "localhost",
				Port:     5432,
				Username: "simple",
				Password: "database_config",
				Database: "simple_database_config",
				AdditionalParams: map[string]string{
					"connect_timeout": "10",
					"sslmode":         "verify-full",
					"sslcert":         "/app/certs/pg.pem",
					"sslkey":          "/app/certs/pg.key",
					"sslrootcert":     "/app/certs/pg_root.pem",
				},
			},
			want: "host=localhost port=5432 user=simple password=database_config dbname=simple_database_config connect_timeout=10 sslcert=/app/certs/pg.pem sslkey=/app/certs/pg.key sslmode=verify-full sslrootcert=/app/certs/pg_root.pem",
		},
	}

	for _, tt := range tests {
		tt := tt // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // marks each test case as capable of running in parallel with each other

			if got := tt.config.ConnectionString(); got != tt.want {
				t.Errorf("invalid connection string, got %q, want %q", got, tt.want)
			}
		})
	}
}
