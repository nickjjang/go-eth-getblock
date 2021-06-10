package controllers

import (
	"ethcache/cache"
	"ethcache/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func getBlockNumber(param string) int64 {
	var num int64
	var err error
	if param == "latest" {
		num = services.EthBlockNumber()
	} else {
		num, err = strconv.ParseInt(param, 10, 64)
		if err != nil {
			num = -1
		}
	}
	return num
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, ETHCACHE!")
}

func Block(c echo.Context) error {
	blockStr := c.ParamValues()[0]
	blockNum := getBlockNumber(blockStr)

	if blockNum == -1 {
		return c.JSON(http.StatusNotAcceptable, map[string]string{"message": "Block number is invalid."})
	}

	block := cache.Cache().Get(blockNum)
	if block == false {
		block = services.EthGetBlockByNumber(blockNum)
		if block == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Not found block."})
		}
		cache.Cache().Put(blockNum, block)
	}
	return c.JSON(http.StatusOK, block)
}

func Transaction(c echo.Context) error {
	blockStr := c.ParamValues()[0]
	blockNum := getBlockNumber(blockStr)

	if blockNum == -1 {
		return c.JSON(http.StatusNotAcceptable, map[string]string{"message": "Block number is invalid."})
	}

	block := cache.Cache().Get(blockNum)
	if block == false {
		block = services.EthGetBlockByNumber(blockNum)
		if block == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Not found block."})
		}
		cache.Cache().Put(blockNum, block)
	}

	txsNode := block.(map[string]interface{})["transactions"]
	txs := txsNode.([]interface{})
	txsStr := c.ParamValues()[1]
	for _, tx := range txs {
		if strings.ToLower(tx.(map[string]interface{})["hash"].(string)) == strings.ToLower(txsStr) {
			return c.JSON(http.StatusOK, tx)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Transaction is not exist in the block."})
}
