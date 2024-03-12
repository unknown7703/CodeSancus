//mod number
const mod int64 =33554393
//base number 
const base int64 =747287

//token hash function
func tokenHash(token string) int64 {
	var hash int =0
	for i:=0;i<len(token);i++{
		hash=((hash+token[i])*base)%mod
	}
	return hash
}