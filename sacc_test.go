
/*
Copyright Hitachi America Ltd. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
        "fmt"
        "testing"

        "github.com/hyperledger/fabric-chaincode-go/shim"
        "github.com/hyperledger/fabric-chaincode-go/shimtest"
)

func checkInit(t *testing.T, stub *shimtest.MockStub, args [][]byte) {
        res := stub.MockInit("1", args)
        if res.Status != shim.OK {
                fmt.Println("Init failed", string(res.Message))
                t.FailNow()
        }
}

func checkState(t *testing.T, stub *shimtest.MockStub, name string, value string) {
        bytes := stub.State[name]
        if bytes == nil {
                fmt.Println("State", name, "failed to get value")
                t.FailNow()
        }
        if string(bytes) != value {
          Last login: Thu Sep 17 02:52:51 on ttys000
yejinrozdynamic:~ user$ cd Downloads
yejinrozdynamic:Downloads user$ cd hyper
yejinrozdynamic:hyper user$ vagrant up
Bringing machine 'node1' up with 'virtualbox' provider...
==> node1: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> node1: flag to force provisioning. Provisioners marked to run always will still run.
yejinrozdynamic:hyper user$ vagrant ssh node1
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.4.0-159-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

121 packages can be updated.
89 updates are security updates.

New release '18.04.5 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


Last login: Wed Sep 16 17:57:32 2020 from 10.0.2.2
vagrant@node1:~$ cd hyper-samples
-bash: cd: hyper-samples: No such file or directory
vagrant@node1:~$ ls
fabric-samples  README.md
vagrant@node1:~$ cd fabric-samples
vagrant@node1:~/fabric-samples$ ls
asset-transfer-basic              bin                 CODEOWNERS        high-throughput      README.md
asset-transfer-ledger-queries     chaincode           commercial-paper  interest_rate_swaps  scripts
asset-transfer-private-data       CHANGELOG.md        config            LICENSE              SECURITY.md
asset-transfer-sbe                ci                  CONTRIBUTING.md   MAINTAINERS.md       test-application
asset-transfer-secured-agreement  CODE_OF_CONDUCT.md  fabcar            off_chain_data       test-network
vagrant@node1:~/fabric-samples$ cd chaincode
vagrant@node1:~/fabric-samples/chaincode$ ls
abac  abstore  fabcar  marbles02  marbles02_private  README.md  sacc
vagrant@node1:~/fabric-samples/chaincode$ cd fabcar
vagrant@node1:~/fabric-samples/chaincode/fabcar$ ls
external  go  java  javascript  typescript
vagrant@node1:~/fabric-samples/chaincode/fabcar$ cd go
vagrant@node1:~/fabric-samples/chaincode/fabcar/go$ ls
fabcar.go  go.mod  go.sum
vagrant@node1:~/fabric-samples/chaincode/fabcar/go$ cd ..
vagrant@node1:~/fabric-samples/chaincode/fabcar$ ls
external  go  java  javascript  typescript
vagrant@node1:~/fabric-samples/chaincode/fabcar$ cd javascript
vagrant@node1:~/fabric-samples/chaincode/fabcar/javascript$ ls
index.js  lib  package.json
vagrant@node1:~/fabric-samples/chaincode/fabcar/javascript$ cd ..
vagrant@node1:~/fabric-samples/chaincode/fabcar$ cd ..
vagrant@node1:~/fabric-samples/chaincode$ cd ..
vagrant@node1:~/fabric-samples$ ls
asset-transfer-basic              bin                 CODEOWNERS        high-throughput      README.md
asset-transfer-ledger-queries     chaincode           commercial-paper  interest_rate_swaps  scripts
asset-transfer-private-data       CHANGELOG.md        config            LICENSE              SECURITY.md
asset-transfer-sbe                ci                  CONTRIBUTING.md   MAINTAINERS.md       test-application
asset-transfer-secured-agreement  CODE_OF_CONDUCT.md  fabcar            off_chain_data       test-network
vagrant@node1:~/fabric-samples$ cd fabcar
vagrant@node1:~/fabric-samples/fabcar$ ls
go  java  javascript  networkDown.sh  startFabric.sh  typescript
vagrant@node1:~/fabric-samples/fabcar$ cd javascript
vagrant@node1:~/fabric-samples/fabcar/javascript$ ls
enrollAdmin.js  invoke.js  package.json  package-lock.json  query.js  registerUser.js  wallet
vagrant@node1:~/fabric-samples/fabcar/javascript$ vi query.js
vagrant@node1:~/fabric-samples/fabcar/javascript$ vi enrollAdmin.js
vagrant@node1:~/fabric-samples/fabcar/javascript$ cd ..
vagrant@node1:~/fabric-samples/fabcar$ cd go
vagrant@node1:~/fabric-samples/fabcar/go$ ls
fabcar.go  go.mod  go.sum  runfabcar.sh
vagrant@node1:~/fabric-samples/fabcar/go$ cd ..
vagrant@node1:~/fabric-samples/fabcar$ ls
go  java  javascript  networkDown.sh  startFabric.sh  typescript
vagrant@node1:~/fabric-samples/fabcar$ cd ..
vagrant@node1:~/fabric-samples$ cd chaincode
vagrant@node1:~/fabric-samples/chaincode$ ls
abac  abstore  fabcar  marbles02  marbles02_private  README.md  sacc
vagrant@node1:~/fabric-samples/chaincode$ cd sacc
vagrant@node1:~/fabric-samples/chaincode/sacc$ ls
go.mod  go.sum  sacc.go  sacc_test.go
vagrant@node1:~/fabric-samples/chaincode/sacc$ vi sacc_test.go

        if string(bytes) != value {
                fmt.Println("State value", name, "was not", value, "as expected")
                t.FailNow()
        }
}

func checkQuery(t *testing.T, stub *shimtest.MockStub, name string, value string) {
        res := stub.MockInvoke("1", [][]byte{[]byte("query"), []byte(name)})
        if res.Status != shim.OK {
                fmt.Println("Query", name, "failed", string(res.Message))
                t.FailNow()
        }
        if res.Payload == nil {
                fmt.Println("Query", name, "failed to get value")
                t.FailNow()
        }
        if string(res.Payload) != value {
                fmt.Println("Query value", name, "was not", value, "as expected")
                t.FailNow()
        }
}

func checkInvoke(t *testing.T, stub *shimtest.MockStub, args [][]byte) {
        res := stub.MockInvoke("1", args)
        if res.Status != shim.OK {
                fmt.Println("Invoke", args, "failed", string(res.Message))
                t.FailNow()
        }
}

func TestSacc_Init(t *testing.T) {
        cc := new(SimpleAsset)
        stub := shimtest.NewMockStub("sacc", cc)
        // Init a=10
        checkInit(t, stub, [][]byte{[]byte("a"), []byte("10")})

        checkState(t, stub, "a", "10")
}

func TestSacc_Query(t *testing.T) {
        cc := new(SimpleAsset)
        stub := shimtest.NewMockStub("sacc", cc)

        // Init a=10
        checkInit(t, stub, [][]byte{[]byte("a"), []byte("10")})

        // Query a
        checkQuery(t, stub, "a", "10")
}

func TestSacc_Invoke(t *testing.T) {
        cc := new(SimpleAsset)
        stub := shimtest.NewMockStub("sacc", cc)

        // Init a=10
        checkInit(t, stub, [][]byte{[]byte("a"), []byte("10")})

        // Invoke: Set a=20
        checkInvoke(t, stub, [][]byte{[]byte("set"), []byte("a"), []byte("20")})

        // Query a
        checkQuery(t, stub, "a", "20")
}

func TestSacc_InitWithIncorrectArguments(t *testing.T) {
  cc := new(SimpleAsset)
       stub := shimtest.NewMockStub("sacc", cc)

       // Init with incorrect arguments
       res := stub.MockInit("1", [][]byte{[]byte("a"), []byte("10"), []byte("10")})

       if res.Status != shim.ERROR {
               fmt.Println("Invalid Init accepted")
               t.FailNow()
       }

       if res.Message != "Incorrect arguments. Expecting a key and a value" {
               fmt.Println("Unexpected Error message:", string(res.Message))
               t.FailNow()
       }
}

func TestSacc_QueryWithIncorrectArguments(t *testing.T) {
       cc := new(SimpleAsset)
       stub := shimtest.NewMockStub("sacc", cc)

       // Init a=10
       checkInit(t, stub, [][]byte{[]byte("a"), []byte("10")})

       // Query with incorrect arguments
       res := stub.MockInvoke("1", [][]byte{[]byte("query"), []byte("a"), []byte("b")})

       if res.Status != shim.ERROR {
               fmt.Println("Invalid query accepted")
               t.FailNow()
       }

       if res.Message != "Incorrect arguments. Expecting a key" {
         fmt.Println("Unexpected Error message:", string(res.Message))
               t.FailNow()
       }
}

func TestSacc_QueryForAssetNotFound(t *testing.T) {
       cc := new(SimpleAsset)
       stub := shimtest.NewMockStub("sacc", cc)

       // Init a=10
       checkInit(t, stub, [][]byte{[]byte("a"), []byte("10")})

       // Query for b (as a asset not found)
       res := stub.MockInvoke("1", [][]byte{[]byte("query"), []byte("b")})

       if res.Status != shim.ERROR {
               fmt.Println("Invalid query accepted")
               t.FailNow()
       }

       if res.Message != "Asset not found: b" {
               fmt.Println("Unexpected Error message:", string(res.Message))
               t.FailNow()
       }
}

func TestSacc_InvokeWithIncorrectArguments(t *testing.T) {
       cc := new(SimpleAsset)
       stub := shimtest.NewMockStub("sacc", cc)

       // Init a=10
       checkInit(t, stub, [][]byte{[]byte("a"), []byte("10")})
       // Invoke with incorrect arguments
      res := stub.MockInvoke("1", [][]byte{[]byte("set"), []byte("a")})
      if res.Status != shim.ERROR {
              fmt.Println("Invalid Invoke accepted")
              t.FailNow()
      }

      if res.Message != "Incorrect arguments. Expecting a key and a value" {
              fmt.Println("Unexpected Error message:", string(res.Message))
              t.FailNow()
      }
}
