package luck_draw

import (
	"fmt"
	"github.com/gin-gonic/gin"
	xuperaccount "github.com/xuperchain/xuper-sdk-go/account"
	"github.com/xuperchain/xuper-sdk-go/contract"
	"net/http"
)

const (
	GET_LUCK_ID="getLuckid"
	START_LUCK_DRAW="startLuckDraw"
	GET_RESULT = "getResult"
)
var (
	account *xuperaccount.Account
	node string
	bcname string
)

func init(){
	var err error
	account,err = xuperaccount.GetAccountFromPlainFile("/Users/chenfengjin/baidu/xuperchain/output/data/keys/");

	if err!=nil{
		panic(err)
	}
	node="127.0.0.1:37101"
	bcname = "xuper"
}




func Deploy(c*gin.Context){
	args:= &struct {
		ContractId string `json:"contract_id"`
	}{}

	if err:=c.BindJSON(args);err!=nil{
		c.Error(err)
	}
	if args.ContractId==""{
		fmt.Println("args",args)
		c.JSON(http.StatusBadRequest,"missing contract_id")
		return
	}
	client := contract.InitWasmContract(account,"127.0.0.1:37101",bcname,args.ContractId,"XC1111111111111111@xuper")

	TxId,err:=client.DeployWasmContract(map[string]string{"admin":account.Address},"contract/luck_draw.wasm","c")
	if err!=nil{
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"txid":TxId,
	})

}
func GetLuckId(c*gin.Context){

	args:= &struct {
		ContractId string `json:"contract_id"`
	}{}

	if err:=c.BindJSON(args);err!=nil{
		c.Error(err)
	}
	if args.ContractId==""{
		c.JSON(http.StatusBadRequest,"missing contract_id")
	}

	client := contract.InitWasmContract(account,node,"xuper",args.ContractId,"XC1111111111111111@xuper")
	preInvokeResp, err := client.PreInvokeWasmContract(GET_LUCK_ID, map[string]string{})


	_, err=client.PostWasmContract(preInvokeResp)
	if err!=nil{
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"data":map[string]string{
			"id":string(preInvokeResp.GetResponse().GetResponse()[0]),
		},
	})
}
func StartLuckDraw(c*gin.Context){
	args:= &struct {
		ContractId string `json:"contract_id"`
		Seed string `json:"seed"`
	}{}
	c.BindJSON(args)
	client := contract.InitWasmContract(account,node,bcname,args.ContractId,"XC1111111111111111@xuper")
	preInvokeResp,err:= client.PreInvokeWasmContract(START_LUCK_DRAW,map[string]string{
		"seed":args.Seed,
	})
	if err!=nil{
		c.Error(err)
		return
	}
	luckId :=string(preInvokeResp.GetResponse().GetResponse()[0])
	_,err =client.PostWasmContract(preInvokeResp)
	if err!=nil{
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"data":map[string]string{
			"luck_id":luckId,
		},
	})
}

func GetResult(c*gin.Context){


	args:= &struct {
		ContractId string `json:"contract_id"`
	}{}

	if err:=c.BindJSON(args);err!=nil{
		c.Error(err)
	}
	if args.ContractId==""{
		c.JSON(http.StatusBadRequest,"missing contract_id")
	}

	client := contract.InitWasmContract(account,node,"xuper",args.ContractId,"XC1111111111111111@xuper")
	preInvokeResp, err := client.PreInvokeWasmContract(GET_RESULT, map[string]string{})


	_, err=client.PostWasmContract(preInvokeResp)
	if err!=nil{
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"data":map[string]string{
			"luck_user":string(preInvokeResp.GetResponse().GetResponse()[0]),
		},
	})

}

