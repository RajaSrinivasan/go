package password

import (
	"os"
	"fmt"
	"testing"
)

func TestSaltLength(arg *testing.T) {
	fmt.Printf("Salt Length %d\n", saltLength)
}

func TestStore(arg *testing.T) {
	wd,_ := os.Getwd()
	fmt.Printf("Working dir %s\n",wd)
	store := New("password.conf")
	store.Set("app" , "creator", "rsrinivasan","designer")
	store.Verify("app" , "creator" , "rsrinivasan","designer")
	store.Set("app2" , "admin" , "srini" , "lowlife")
	store.Verify("app2" , "admin" , "srini" , "lowlife")
}
