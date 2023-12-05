package api

import (
	"errors"
	"io"
	"log"
	"net/http"
)

const FuncNameEtherInfo = "homepage"

type EtherInfoData struct {
	SnapshootWan []WAN `json:"snapshoot_wan"`
}
type WAN struct {
	Id        int    `json:"id"`
	Errmsg    string `json:"errmsg"`
	Gateway   string `json:"gateway"`
	Interface string `json:"interface"`
	IpAddr    string `json:"ip_addr"`
}

func (i *IKuai) ShowEtherInfoByComment(url string, hostname string, token string) error {
	param := struct {
		Type string `json:"TYPE"`
	}{
		Type: "ether_info,snapshoot",
	}
	req := CallReq{
		FuncName: FuncNameEtherInfo,
		Action:   "show",
		Param:    &param,
	}
	result := EtherInfoData{}
	resp := CallResp{Data: &result}
	err := postJson(i.client, i.baseurl+"/Action/call", &req, &resp)
	if err != nil {
		return err
	}
	if resp.Result != 30000 {
		return errors.New(resp.ErrMsg)

	}
	IpAddr := ""
	for _, wan := range result.SnapshootWan {
		if wan.IpAddr != "" {
			IpAddr = wan.IpAddr
			break
		}
	}
	log.Println("外网ip：" + IpAddr)
	url += "?hostname=" + hostname + "&token=" + token + "&ipv4=" + IpAddr
	respd, errs := http.Get(url)
	if errs != nil {
		return err
	}
	if respd.StatusCode != 200 {
		err = errors.New(respd.Status)
		return err
	}
	defer respd.Body.Close()
	body, err := io.ReadAll(respd.Body)
	if err != nil {
		return err
	}
	log.Println(hostname + "绑定到" + IpAddr + ":" + string(body))
	return nil
}
