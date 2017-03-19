/* Cows and Bulls logical layer
- Balamurugan Alagumalai
- copyright -http://balaweblog.com
*/

package CowsandBulls

import (
"math/rand"
"strings"
"time")

var(
alphabets []string
)

/* Init function */
func init(){
    alphabets=[]string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","1"}
}

/* return the system guess */ 
func Systemguess()([]string){
    return randomletters()
}

/* get the output values for each run */
func GetOutput(userguess string, systemguess []string)(bulls,cows int){
    return cowsandbulls(userguess,systemguess)
}

/* Generate 4 digit random letters */
func randomletters()([]string){
    var numbers []int
    rand.Seed(time.Now().UnixNano())
    for i:=0;i<4;i++{
        numbers = append(numbers,int(rand.Int31n(26)))
    }
    
    return []string{alphabets[numbers[0]],alphabets[numbers[1]],alphabets[numbers[2]],alphabets[numbers[3]]}
}

/* get the cows & bulls */ 
func cowsandbulls(userguess string, systemguess []string)(bulls int,cows int){
    for index,value:= range systemguess{
        // bulls
        if strings.Contains(userguess,value){
        bulls++    
        }
        // cows
        if value == string([]rune(userguess)[index]){
        cows++ 
        }
    }
    return bulls,cows
}