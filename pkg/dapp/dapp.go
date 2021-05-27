package dapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuperchain/dapp-demo/pkg/config"
	xuperaccount "github.com/xuperchain/xuper-sdk-go/account"
	"github.com/xuperchain/xuper-sdk-go/contract"
)

const (
	GET_LUCK_ID     = "getLuckid"
	START_LUCK_DRAW = "startLuckDraw"
	GET_RESULT      = "getResult"
)

var (
	account *xuperaccount.Account
)

func init() {
	var err error
	account, err = xuperaccount.GetAccountFromPlainFile(config.KeyPath)

	if err != nil {
		panic(err)
	}
}

func Deploy(c *gin.Context) {
	args := &struct {
		ContractId string `json:"contract_id"`
	}{}

	if err := c.BindJSON(args); err != nil {
		c.Error(err)
	}
	if args.ContractId == "" {
		c.JSON(http.StatusBadRequest, "missing contract_id")
		return
	}
	client := contract.InitWasmContract(
		account,
		config.Host,
		config.BCName,
		args.ContractId,
		config.ContractAccount,
	)

	TxId, err := client.DeployWasmContract(
		map[string]string{"admin": account.Address},
		config.CodePath,
		"c",
	)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"txid": TxId,
	})

}
func GetLuckId(c *gin.Context) {

	args := &struct {
		ContractId string `json:"contract_id"`
	}{}

	if err := c.BindJSON(args); err != nil {
		c.Error(err)
	}
	if args.ContractId == "" {
		c.JSON(http.StatusBadRequest, "missing contract_id")
	}

	client := contract.InitWasmContract(
		account,
		config.Host,
		config.BCName,
		args.ContractId,
		config.ContractAccount,
	)

	preInvokeResp, err := client.PreInvokeWasmContract(GET_LUCK_ID, map[string]string{})
	if err != nil {
		c.Error(err)
		return
	}
	_, err = client.PostWasmContract(preInvokeResp)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{
			"id": string(preInvokeResp.GetResponse().GetResponse()[0]),
		},
	})
}
func StartLuckDraw(c *gin.Context) {
	args := &struct {
		ContractId string `json:"contract_id"`
		Seed       string `json:"seed"`
	}{}
	c.BindJSON(args)
	client := contract.InitWasmContract(account, config.Host, config.BCName, args.ContractId, "XC1111111111111111@xuper")
	preInvokeResp, err := client.PreInvokeWasmContract(START_LUCK_DRAW, map[string]string{
		"seed": args.Seed,
	})
	if err != nil {
		c.Error(err)
		return
	}
	luckUser := string(preInvokeResp.GetResponse().GetResponse()[0])
	_, err = client.PostWasmContract(preInvokeResp)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{
			"luck_user": luckUser,
		},
	})
}

func GetResult(c *gin.Context) {

	args := &struct {
		ContractId string `json:"contract_id"`
	}{}

	if err := c.BindJSON(args); err != nil {
		c.Error(err)
	}
	if args.ContractId == "" {
		c.JSON(http.StatusBadRequest, "missing contract_id")
	}

	client := contract.InitWasmContract(account, config.Host, "xuper", args.ContractId, "XC1111111111111111@xuper")
	preInvokeResp, err := client.PreInvokeWasmContract(GET_RESULT, map[string]string{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{
			"luck_user": string(preInvokeResp.GetResponse().GetResponse()[0]),
		},
	})

}
