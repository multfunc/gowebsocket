package models_x

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"gowebsocket/utils_x"
	"strings"
)

/**
* @Author junfenghe
* @Description create `MessageGroupLive`
* @Date 2021-03-23 10:40
* @Param
* @return
**/
func CreateMessageGroupLive(detail *MessageGroupLive) (i interface{}, err error) {
	uuid, err := utils_x.GenerateUUID()
	if err != nil {
		return nil, err
	}
	byteDetail, err := json.Marshal(&detail)

	if err != nil {
		return nil, err
	}
	req := esapi.IndexRequest{
		Index:      "message_group_live",
		DocumentID: uuid.(string),
		Body:       strings.NewReader(string(byteDetail)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), clientES)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, errors.New(res.Status())
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return r, nil
}
