package muid

import (
  "io/ioutil"
  "fmt"
)

// TODO: This will only work for linux but Multiverse OS is a linux operating
// system and so our motivation to support other operating systems is not really
// existant; although if other developers find this project useful and
// appreciate our upgrades over the popular xid library we will gladly add in
// other os methods of pulling machine-id -- but it will work regardless of OS
// beacuse we fallback to using random bytes which for most usecases is just as
// good, in the very rare use case it needs to be ran on hundreds of thousands
// of machines, is when the machine id is really desired. 
func machineID() ([]byte, error) {
	idBytes, err := ioutil.ReadFile("/etc/machine-id")
  if err != nil || len(idBytes) == 0 {
		idBytes, err = ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
    if err != nil || len(idBytes) == 0 {
      fmt.Println("failed to get machine id bytes, should generate random bytes which is fine for almost all use cases")
      // TODO: Pull the correct number of bytes to replace 
      //       machine id with random sequence which is 
      //       perfectly fine in almost every usecase. 
      // TODO: Determine the length of the bytes and then generate that number
      //       of random bytes and return it

    }
	}
  
	return idBytes, err
}
