/* Cows and Bulls logical layer
- Balamurugan Alagumalai
- copyright -http://balaweblog.com
*/

package CowsandBullstest

import(
"testing"
"CowsandBulls" 
"fmt")

var (
systemguess []string
)

/* Test System guess function */
func TestSystemguess(t *testing.T){
    systemguess= CowsandBulls.Systemguess()
    
    if len(systemguess)!=4 {
        t.Fail()
        t.Log("Invalid system guess")
    }
    fmt.Println("system guess",systemguess)
}

/* Test get output first run */
func TestGetOutputRun1(t *testing.T){
    userguess:="TEST"
    cows,bulls := CowsandBulls.GetOutput(userguess,systemguess)
    if cows<0 && bulls<0{
        t.Fail()
        t.Log("Invalid output in first run")
    }
    fmt.Println("user first guess" ,userguess)
    fmt.Println("cows and bulls are" ,cows,bulls)
}
/* Test get output second run */
func TestGetOutputRun2(t *testing.T){
    userguess:="VALS"
    cows,bulls := CowsandBulls.GetOutput(userguess,systemguess)
    if cows<0 && bulls<0{
        t.Fail()
        t.Log("Invalid output in second run")
    }
    fmt.Println("user second guess" ,userguess)
    fmt.Println("cows and bulls are" ,cows,bulls)
}
/* Test get output third run */
func TestGetOutputRun3(t *testing.T){
    userguess:="SAND"
    cows,bulls := CowsandBulls.GetOutput(userguess,systemguess)
    if cows<0 && bulls<0{
        t.Fail()
        t.Log("Invalid output in third run")
    }
    fmt.Println("user third guess" ,userguess)
    fmt.Println("cows and bulls are" ,cows,bulls)
}

/* Test get output Fourth run */
func TestGetOutputRun4(t *testing.T){
    userguess:="NAMS"
    cows,bulls := CowsandBulls.GetOutput(userguess,systemguess)
    if cows<0 && bulls<0{
        t.Fail()
        t.Log("Invalid output in fourth run")
    }
    fmt.Println("user fourth guess" ,userguess)
    fmt.Println("cows and bulls are" ,cows,bulls)
}


