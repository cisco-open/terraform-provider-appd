package provider

var resourceConnectionAwsAccessSecretMap = map[string]interface{}{
	"connection_details": map[string]interface{}{
		"access_type": map[string]interface{}{
			"valid":           []interface{}{"access_key"},
			"invalid":         []interface{}{"k3ta9kzd5c"},
			"multiple_valids": []interface{}{"access_key"},
		},

		"access_key_id": map[string]interface{}{
			"valid":           []interface{}{"AKIA4Q3VJVOID2MGTYLU", "AKIA4Q3VJVOID2MGTYLU"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"AKIA4Q3VJVOID2MGTYLU", "AKIA4Q3VJVOID2MGTYLU"},
		},

		"secret_access_key": map[string]interface{}{
			"valid":           []interface{}{"fELEQROV7T7p4ajIJ5pO3d8xmz849R8nJgxazUdE", "fELEQROV7T7p4ajIJ5pO3d8xmz849R8nJgxazUdE"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"fELEQROV7T7p4ajIJ5pO3d8xmz849R8nJgxazUdE", "fELEQROV7T7p4ajIJ5pO3d8xmz849R8nJgxazUdE"},
		},
	},
}
