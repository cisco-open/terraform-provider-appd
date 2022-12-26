package provider

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var expectErrorMap = map[string]string{
	"IsCIDR":               "expected (.)+ to be a valid IPv4 Value, got (.)+: (.)+",
	"IsIPAddress":          "expected (.)+ to contain a valid IP, got: (.)+",
	"IsIPv4Address":        "expected (.)+ to contain a valid IPv4 address, got: (.)+",
	"IsIPv6Address":        "expected (.)+ to contain a valid IPv6 address, got: (.)+",
	"IsMACAddress":         "expected (.)+ to be a valid MAC address, got (.)+: (.)+",
	"IsRFC3339Time":        "expected (.)+ to be a valid RFC3339 date, got (.)+: (.)+",
	"IsURLWithHTTPS":       "",
	"IsURLWithHTTPorHTTPS": "",
	"IsUUID":               "expected (.)+ to be a valid UUID, got (.)+",
	"StringIsBase64":       "expected (.)+ to be a base64 string, got (.)+",
	"StringIsJSON":         "(.)+ contains an invalid JSON: (.)+",
	"StringIsValidRegExp":  "(.)+: (.)+",
	"StringInSlice":        "expected (.)+ to be one of (.)+, got (.)+",
	"StringNotInSlice":     "expected (.)+ to not be any of (.)+, got (.)+",
	"IsCIDRNetwork":        "expected (.)+ to contain a network Value with between (.)+ and (.)+ significant bits, got: (.)+",
	"IntBetween":           "expected (.)+ to be in the range ((.)+ - (.)+), got (.)+",
	"IsPortNumber":         "expected (.)+ to be a valid port number, got: (.)+",
	"IsPortNumberOrZero":   "expected (.)+ to be a valid port number or 0, got: (.)+",
	"FloatBetween":         "expected (.)+ to be in the range ((.)+ - (.)+), got (.)+",
}

var Test = map[string]interface{}{
	"base64": map[string]interface{}{
		"valid":           []interface{}{"ZDB1d3kyc2o1cw==", "a2VnbzF1NHZwaw==", "M3NrMHM3ZzM2YQ==", "ejlwZHlydmk5Ng=="},
		"invalid":         []interface{}{"a3+J1b%mFs//"},
		"multiple_valids": []interface{}{"dGUzbnF6bnBtNg==", "ZDN3aGpwbGZ3cg==", "empzY2hpaW1vZQ==", "M2dqNWNwaWJiag==", "Mmw5dXdkaDBvYg==", "OXdlNXUzZTBubQ==", "ZDV4aGwyN25mOA==", "eDU5d213cG40Ng==", "M2V6Z2djaWZwNw==", "ZGR3MGwwYXV1bQ==", "bjU2NWVudG11bg==", "Z3IyamVudzFtNQ==", "YzRnNTU5bWY2MA==", "dGY0OTB2bG5iNA==", "dWk3dmp3OTRiNg=="},
	},
	"cidr": map[string]interface{}{
		"valid":           []interface{}{"233.76.0.0/20", "233.76.128.0/20", "233.76.160.0/20", "233.76.144.0/20"},
		"invalid":         []interface{}{"258.269.265.297/10"},
		"multiple_valids": []interface{}{"233.76.0.0/20", "233.76.128.0/20", "233.76.160.0/20", "233.76.144.0/20", "233.76.240.0/20", "233.76.80.0/20", "233.76.16.0/20", "233.76.224.0/20", "233.76.208.0/20", "233.76.112.0/20", "233.76.32.0/20", "233.76.96.0/20", "233.76.192.0/20", "233.76.48.0/20", "233.76.176.0/20"},
	},
	"ipv4": map[string]interface{}{
		"valid":           []interface{}{"233.76.74.98", "233.76.94.167", "233.76.23.30", "233.76.202.254"},
		"invalid":         []interface{}{"279.284.273.298"},
		"multiple_valids": []interface{}{"233.76.74.98", "233.76.94.167", "233.76.23.30", "233.76.202.254", "233.76.195.135", "233.76.69.199", "233.76.101.41", "233.76.129.109", "233.76.199.174", "233.76.87.175", "233.76.185.77", "233.76.83.55", "233.76.106.36", "233.76.5.115", "233.76.117.22"},
	},
	"ipv6": map[string]interface{}{
		"valid":           []interface{}{"2001:db8::34f4:0:0:f35a", "2001:db8::34f4:0:0:f361", "2001:db8::34f4:0:0:f3e7", "2001:db8::34f4:0:0:f3cb"},
		"invalid":         []interface{}{"invalidIPv6"},
		"multiple_valids": []interface{}{"2001:db8::34f4:0:0:f35a", "2001:db8::34f4:0:0:f361", "2001:db8::34f4:0:0:f3e7", "2001:db8::34f4:0:0:f3cb", "2001:db8::34f4:0:0:f335", "2001:db8::34f4:0:0:f3e6", "2001:db8::34f4:0:0:f3cf", "2001:db8::34f4:0:0:f317", "2001:db8::34f4:0:0:f3f3", "2001:db8::34f4:0:0:f36f", "2001:db8::34f4:0:0:f313", "2001:db8::34f4:0:0:f347", "2001:db8::34f4:0:0:f372", "2001:db8::34f4:0:0:f301", "2001:db8::34f4:0:0:f3cd"},
	},
	"json": map[string]interface{}{
		"valid":           []interface{}{"json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })"},
		"invalid":         []interface{}{"json({ name : val)"},
		"multiple_valids": []interface{}{"json({ \"attribute\" : \"value0\" })", "json({ \"attribute\" : \"value1\" })", "json({ \"attribute\" : \"value2\" })", "json({ \"attribute\" : \"value3\" })", "json({ \"attribute\" : \"value4\" })", "json({ \"attribute\" : \"value5\" })", "json({ \"attribute\" : \"value6\" })", "json({ \"attribute\" : \"value7\" })", "json({ \"attribute\" : \"value8\" })", "json({ \"attribute\" : \"value9\" })", "json({ \"attribute\" : \"value10\" })", "json({ \"attribute\" : \"value11\" })", "json({ \"attribute\" : \"value12\" })", "json({ \"attribute\" : \"value13\" })", "json({ \"attribute\" : \"value14\" })"},
	},
	"mac": map[string]interface{}{
		"valid":           []interface{}{"0d:eb:72:89:2d:14", "89:72:f3:18:0f:75", "7f:e2:03:59:52:87", "f0:28:9e:bf:4e:87"},
		"invalid":         []interface{}{"invalidMAC"},
		"multiple_valids": []interface{}{"1b:09:8e:12:e2:36", "4f:ef:c3:9c:a2:e3", "c8:23:25:71:7b:ef", "8f:9f:4f:a8:79:83", "17:69:13:ff:a1:57", "c9:53:54:46:e2:0a", "20:17:97:77:c0:aa", "d9:a2:84:da:04:0d", "80:d9:94:3d:dd:bc", "43:a4:b5:d4:b9:8b", "1c:b1:14:4c:98:ad", "fe:07:61:9a:eb:71", "9f:db:5e:59:8d:5e", "00:27:17:b0:f1:65", "e8:ed:2c:1e:57:1a"},
	},
	"regex": map[string]interface{}{
		"valid":           []interface{}{"(?m)^[0-9]{2}$", "^(\\$)(\\d)+"},
		"invalid":         []interface{}{"[0-9)++"},
		"multiple_valids": []interface{}{"(?m)^[0-9]{2}$", "^(\\$)(\\d)+"},
	},
	"string": map[string]interface{}{
		"valid":           []interface{}{"aczlh43f2i", "dli2w7tdda", "eeio75z94b", "dlz4c74rd8"},
		"invalid":         []interface{}{12345},
		"multiple_valids": []interface{}{"5hkcf9qhl5", "fsejckwyzh", "0p0217eysy", "tsjzlj43fs", "vjtsyyu0ch", "ebl6i4rdxd", "uv0op6yfxu", "vvfm51qwjx", "n8qk1ditc1", "x1v71rojh4", "yfoa2r3pdl", "x97hc6657f", "s4ucwpef4e", "98wju6w845", "v91suo2xew"},
	},
	"time": map[string]interface{}{
		"valid":           []interface{}{"2022-12-08T06:50:46.677392+00:00", "2022-12-31T06:50:46.677392+00:00", "2023-01-23T06:50:46.677392+00:00", "2023-02-15T06:50:46.677392+00:00"},
		"invalid":         []interface{}{"2022-12-08 12:20:46.677392"},
		"multiple_valids": []interface{}{"2022-12-08T06:50:46.677392+00:00", "2022-12-31T06:50:46.677902+00:00", "2023-01-23T06:50:46.677902+00:00", "2023-02-15T06:50:46.677902+00:00", "2023-03-10T06:50:46.677902+00:00", "2023-04-02T06:50:46.677902+00:00", "2023-04-25T06:50:46.677902+00:00", "2023-05-18T06:50:46.677902+00:00", "2023-06-10T06:50:46.677902+00:00", "2023-07-03T06:50:46.677902+00:00", "2023-07-26T06:50:46.677902+00:00", "2023-08-18T06:50:46.677902+00:00", "2023-09-10T06:50:46.677902+00:00", "2023-10-03T06:50:46.677902+00:00", "2023-10-26T06:50:46.677902+00:00"},
	},
	"url-http": map[string]interface{}{
		"valid":           []interface{}{"http://nb0d4sbmf570jof.com", "http://4x66ud7ps5vxh8k.com", "http://rr789d07qlx5fl1.com", "http://umkobyu41w2emo7.com"},
		"invalid":         []interface{}{"ht:/2z5zlkwckdlz51i.com"},
		"multiple_valids": []interface{}{"http://vd3upfotvaq8253.com", "http://ljly8uwkubq4n32.com", "http://a3qa0ff77uqruxy.com", "http://1i7zn7p2bmc6wo8.com", "http://c8go50qoro6enuc.com", "http://soi5e8hiuaebvqq.com", "http://bsmoh9ylb8qq1i5.com", "http://yv5vmlrh11bchpi.com", "http://t10jjy7snl7tomc.com", "http://f83bhvc75a9a4az.com", "http://iik7njjr2doiwgm.com", "http://qa37owmhagbbeoe.com", "http://a4t3syk1uuv2frl.com", "http://pupghkwqb2dju2p.com", "http://atp9oemd6kmzpzr.com"},
	},
	"url-https": map[string]interface{}{
		"valid":           []interface{}{"https://oq7v5vom6nk80wp.com", "https://pqq1167bqfn260t.com", "https://jsu6y0y4renv3xr.com", "https://250k1agm2t4ogdi.com"},
		"invalid":         []interface{}{"hts:/356lphhx93n5l1h.com"},
		"multiple_valids": []interface{}{"https://z32k31zx75b3c82.com", "https://uinkyz3rvu6nn1u.com", "https://koju6afnc164h8d.com", "https://t3hb1nvufb0566w.com", "https://swz2ozc1ysdry9l.com", "https://febibl3yy8ra6vn.com", "https://8raouy6wjzt8l03.com", "https://w3ic162l91ym1mx.com", "https://2bsjco3cpux1gnf.com", "https://3d8h20xrma0paiw.com", "https://43lbgu4zpp2efey.com", "https://awlw7g17qygka52.com", "https://39gjseidcdt1v37.com", "https://csf9nq01739hv79.com", "https://rzy4nnbkns445lz.com"},
	},
	"uuid": map[string]interface{}{
		"valid":           []interface{}{"ea0bb97a-b34b-4995-a154-e7c06e0d4658", "eb964944-ef85-42f7-bf8e-91c0bc36be19", "0e5763d3-7fc8-4011-8d82-4180afca6ba6", "29369e95-ebd0-4ba0-93da-123601a42012"},
		"invalid":         []interface{}{"invalid323Uuid12"},
		"multiple_valids": []interface{}{"e83679d7-be8e-406d-b2eb-407a2b238897", "a4ac5ec7-76c4-11ed-bea1-7c8ae1979a87", "a4ac5ec8-76c4-11ed-96e2-7c8ae1979a87", "a4ac5ec9-76c4-11ed-832a-7c8ae1979a87", "a4ac5eca-76c4-11ed-a27e-7c8ae1979a87", "a4ac5ecb-76c4-11ed-9b0e-7c8ae1979a87", "a4ac5ecc-76c4-11ed-9998-7c8ae1979a87", "a4ac5ecd-76c4-11ed-b8f2-7c8ae1979a87", "a4ac5ece-76c4-11ed-b9ae-7c8ae1979a87", "a4ac5ecf-76c4-11ed-9f1a-7c8ae1979a87", "a4ac5ed0-76c4-11ed-a711-7c8ae1979a87", "a4ac5ed1-76c4-11ed-ad0a-7c8ae1979a87", "a4ac5ed2-76c4-11ed-b82d-7c8ae1979a87", "a4ac5ed3-76c4-11ed-aa01-7c8ae1979a87", "a4ac5ed4-76c4-11ed-8d80-7c8ae1979a87"},
	},
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"appdynamicscloud": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {

	if v := os.Getenv("APPDYNAMICS_TENANT_NAME"); v == "" {
		t.Fatal("APPDYNAMICS_TENANT_NAME env variable must be set for acceptance tests")
	}

	if v := os.Getenv("APPDYNAMICS_SAVE_TOKEN"); v == "" {
		t.Fatal("APPDYNAMICS_SAVE_TOKEN env variable must be set for acceptance tests")
	}

	loginMode := os.Getenv("APPDYNAMICS_LOGIN_MODE")
	if loginMode == "" {
		t.Fatal("APPDYNAMICS_LOGIN_MODE env variable must be set for acceptance tests")
	} else {
		if strings.Contains(loginMode, "service_principal") {
			if v := os.Getenv("APPDYNAMICS_CLIENT_ID"); v == "" {
				t.Fatal("APPDYNAMICS_CLIENT_ID env variable must be set for acceptance tests")
			}
			if v := os.Getenv("APPDYNAMICS_CLIENT_SECRET"); v == "" {
				t.Fatal("APPDYNAMICS_CLIENT_SECRET env variable must be set for acceptance tests")
			}
		} else if strings.Contains(loginMode, "headless") {
			if v := os.Getenv("APPDYNAMICS_USERNAME"); v == "" {
				t.Fatal("APPDYNAMICS_USERNAME env variable must be set for acceptance tests")
			}
			if v := os.Getenv("APPDYNAMICS_PASSWORD"); v == "" {
				t.Fatal("APPDYNAMICS_PASSWORD env variable must be set for acceptance tests")
			}
		}
	}
}

var providerFactories = map[string]func() (*schema.Provider, error){
	"appdynamicscloud": func() (*schema.Provider, error) {
		return testAccProvider, nil
	},
}

func makeTestVariable(s string) string {
	return fmt.Sprintf("TestAcc_%v", s)
}

func searchInObject(testMap map[string]interface{}, attr string) interface{} {
	jsonStr, _ := json.Marshal(testMap)
	jsonParsed, _ := gabs.ParseJSON([]byte(jsonStr))
	return jsonParsed.Path(attr).Data()
}
