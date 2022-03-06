# Notes

```go
config := vault.DefaultConfig()
config.Address = "http://127.0.0.1:8200"
hello := base64.StdEncoding.EncodeToString([]byte("Hello World"))
client, err := vault.NewClient(config)
if err != nil {
	log.Fatal(err)
}
client.SetToken(VAULT_TOKEN)
encrypt, err := client.Logical().Write("transit/encrypt/hello", map[string]interface{}{
	"plaintext": hello,
})
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Ecrypted : %v\n", encrypt.Data["ciphertext"].(string))
```
