package main

import (
	"fmt"
	"time"

	"github.com/dhermes/first-friday-feb-2019/pkg/verify"
)

func main() {
	var err error

	bearerTokenJWT := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjMyMjkwNzIyLTUzYWQtNDAwNi1iYmIxLTk0YWNiNTRlZDRhMiJ9.eyJzb21lIjoicGF5bG9hZCIsImlhdCI6MTU1MDI2NzE1OCwiZXhwIjoxNTUwNTY3MTU4LCJhdWQiOiJ1cm46Zmlyc3QtZnJpZGF5LWZlYi0yMDE5IiwiaXNzIjoidXJuOmZpcnN0LWZyaWRheS1mZWItMjAxOSIsInN1YiI6IjRjMDVkMmNlLTJiMWEtNDMyYS05MjVlLTMyNDJjNWUwMGE3OCJ9.df49RHvCREAcmrPfDihpx8v_c9rLh9h7ScdaLyDE5wlKz8q8rfW07UTJa68VLaeOL-GE1Ug7PtAYklWP93_oYblOvq06dKCXul-WaTYmFAzxRGaUoRvLuXB7Pmv5WUk_SMdh789nvXvJVvWNKADvGdTh97jQ3Zg9v7cr6gWgiQ821i9GkZPejnAx0j81lil-9B2OxvDCUm7QD2Mn0yH32y2THOacvGpY6JsO5C6nIzXkj5A_kBOaSArqYl117h6XkcuvQWpA2dKOZY-EK_gAj-s9IW-pij0zK4-KZhbAijKHM0bGOncZOemaXjs_nSfRRlz5l-UkCVxq_l1XzmrRHg"
	publicKeyPEMBytes := []byte("-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEApbRv8NhJ8jZ2fHK8FlomklkYCb1jxbLTSjNQ8IUdCJ1TBaP0u7sh\n1rvyKhK0TPwx5tZkm4ZtgACKmw8Wfok8Lf8OkYPOYdZ1Lj9ftxIS+B8/S/tld73x\nqwRj3S+iKUH8UNKtVovgdUsojlBvuMe5fwRw/QL0/cO3iwl73vMFjn6MPQWUeZzO\n9S8peDf/HogVhHcO2k2wsunfepgX0cbZPfwwtOQ9ZJq1+RcUNVQaR5EU6CoTnc6l\ntaesywuZi4a4OaN/eMD+ZXX6JEldq3t/PZP2tA3tLtWtkVHJhN6pb8vhzeEEftqV\n5FtrGzJM+H9lrFNJmfNAL05Rc14GZZrqiQIDAQAB\n-----END RSA PUBLIC KEY-----\n") // TODO: Get these from somewhere.
	var valid bool
	valid, err = verify.Verify(bearerTokenJWT, publicKeyPEMBytes, time.Now())
	fmt.Println(valid)
	fmt.Println(err)
}
