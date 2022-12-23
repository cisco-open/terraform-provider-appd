package provider

var resourceConnectionAzureSecretMap = map[string]interface{}{
	"connection_details": map[string]interface{}{
		"client_id": map[string]interface{}{
			"valid":           []interface{}{"adb90c29-204d-43d9-987c-ab406d9199cc", "adb90c29-204d-43d9-987c-ab406d9199cc"},
			"invalid":         []interface{}{"ACTIVE", "INACTIVE"},
			"multiple_valids": []interface{}{"adb90c29-204d-43d9-987c-ab406d9199cc", "adb90c29-204d-43d9-987c-ab406d9199cc"},
		},

		"client_secret": map[string]interface{}{
			"valid":           []interface{}{"ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx", "ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx"},
			"invalid":         []interface{}{"ACTIVE", "INACTIVE"},
			"multiple_valids": []interface{}{"ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx", "ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx"},
		},

		"tenant_id": map[string]interface{}{
			"valid":           []interface{}{"f3db65e4-ce5a-4d7e-a140-255bf017d87f", "f3db65e4-ce5a-4d7e-a140-255bf017d87f"},
			"invalid":         []interface{}{"ACTIVE", "INACTIVE"},
			"multiple_valids": []interface{}{"f3db65e4-ce5a-4d7e-a140-255bf017d87f", "f3db65e4-ce5a-4d7e-a140-255bf017d87f"},
		},

		"subscription_id": map[string]interface{}{
			"valid":           []interface{}{"fca41da2-4908-49e2-b0cb-d3d2080fc5be", "fca41da2-4908-49e2-b0cb-d3d2080fcabc"},
			"invalid":         []interface{}{"ACTIVE", "INACTIVE"},
			"multiple_valids": []interface{}{"fca41da2-4908-49e2-b0cb-d3d2080fc5be", "fca41da2-4908-49e2-b0cb-d3d2080fcabc"},
		},
	},
}