package entry

import (
	"MineArcade-backend/clients/accounts"
	"MineArcade-backend/clients/player_store"
	"MineArcade-backend/protocol"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
)

func Main() {
	account_db := accounts.OpenAccountDB()
	store_db := player_store.OpenPlayerStoreDB()
	defer account_db.Close()
	defer store_db.Close()
	for {
		fmt.Println("----------- MineArcade Manager -----------")
		fmt.Println("  1. 添加账号  2. 删除账号 3. 账号列表")
		fmt.Println("  4. 修改账号  5. 退出")
		fmt.Println("------------------------------------------")
		fmt.Print("请选择选项> ")
		var section string
		_, err := fmt.Scanf("%s ", &section)
		if err != nil {
			fmt.Println("输入错误")
			continue
		}
		switch section {
		case "1":
			AddAccount()
		case "2":
			RemoveAccount()
		case "3":
			ListAccounts()
		case "4":
			ModifyAccount()
		default:
			return
		}
	}
}

func AddAccount() {
	fmt.Print("请输入账号名: ")
	var account_name string
	_, err := fmt.Scanf("%s ", &account_name)
	if err != nil || strings.HasPrefix(account_name, "__") {
		fmt.Println("输入错误")
		return
	}
	if _, ok := accounts.GetUserAuthInfo(account_name); ok {
		fmt.Println("账号已存在")
		return
	}
	for {
		fmt.Print("请输入密码: ")
		var password string
		_, err = fmt.Scanf("%s ", &password)
		if err != nil {
			fmt.Println("输入错误")
			return
		}
		fmt.Print("请输入昵称: ")
		var nickname string
		_, err = fmt.Scanf("%s ", &nickname)
		if err != nil {
			fmt.Println("输入错误")
			return
		}
		pwdMD5 := md5.Sum([]byte(password))
		account := accounts.NewUserAuthInfo(account_name, base64.StdEncoding.EncodeToString(pwdMD5[:]), nickname)
		store := player_store.NewPlayerStore()
		player_store.SavePlayerStore(account.UIDStr, store)
		fmt.Println("账号创建成功")
	}
}

func RemoveAccount() {
	account_db := accounts.OpenAccountDB()
	store_db := player_store.OpenPlayerStoreDB()
	var account_name string
	fmt.Print("请输入账号名: ")
	_, err := fmt.Scanf("%s ", &account_name)
	if err != nil || strings.HasPrefix(account_name, "__") {
		fmt.Println("输入错误")
		return
	}
	if _, ok := accounts.GetUserAuthInfo(account_name); !ok {
		fmt.Println("账号不存在")
		return
	}
	account_db.Delete([]byte(account_name), nil)
	store_db.Delete([]byte(account_name), nil)
	fmt.Println("账号删除成功")
}

func ListAccounts() {
	account_db := accounts.OpenAccountDB()
	iter := account_db.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		acc := &accounts.UserAuthInfo{}
		account_name := string(iter.Key())
		if strings.HasPrefix(account_name, "__") {
			continue
		}
		r := protocol.NewReader(iter.Value())
		acc.Unmarshal(&r)
		fmt.Printf("账号名: %s  昵称: %s\n", account_name, acc.Nickname)
	}
}

func ModifyAccount() {
	var username string
	fmt.Print("请输入账号名: ")
	_, err := fmt.Scanf("%s ", &username)
	if err != nil || strings.HasPrefix(username, "__") {
		fmt.Println("输入错误")
		return
	}
	account, ok := accounts.GetUserAuthInfo(username)
	if !ok {
		fmt.Println("账号不存在")
		return
	}
	for {
		fmt.Println("-------- 请选择一项进行设置 --------")
		fmt.Println("  1 - 修改密码  2 - 修改昵称")
		fmt.Println("  3 - 修改金币  4 - 修改体力")
		fmt.Println("  5 - 修改点数  6 - 修改等级")
		fmt.Println("  7 - 修改经验  其他退出")
		fmt.Println("-----------------------------------")
		fmt.Print(" 请选择选项> ")
		var choice string
		_, err = fmt.Scanf("%s ", &choice)
		if err != nil {
			return
		}
		store := player_store.ReadPlayerStore(account.UIDStr)
		switch choice {
		case "1":
			fmt.Print("请输入新密码: ")
			var newPassword string
			_, err = fmt.Scanf("%s ", &newPassword)
			if err != nil {
				fmt.Println("密码有误")
				return
			}
			password_md5 := md5.Sum([]byte(newPassword))
			account.PasswordMD5 = base64.StdEncoding.EncodeToString(password_md5[:])
			accounts.SaveUserAuthInfo(account)
			fmt.Println("密码设置完成")
		case "2":
			fmt.Print("请输入新昵称: ")
			var newNickname string
			_, err = fmt.Scanf("%s ", &newNickname)
			if err != nil {
				fmt.Println("昵称有误")
				return
			}
			store.Nickname = newNickname
			account.Nickname = newNickname
			accounts.SaveUserAuthInfo(account)
			player_store.SavePlayerStore(account.UIDStr, store)
			fmt.Println("昵称设置完成")
		case "3":
			fmt.Print("请输入金币数: ")
			var newMoney int64
			_, err = fmt.Scanf("%d ", &newMoney)
			if err != nil {
				fmt.Println("金币数有误")
				return
			}
			store.Money = float64(newMoney)
			player_store.SavePlayerStore(account.UIDStr, store)
			fmt.Println("金币数设置完成")
		case "4":
			fmt.Print("请输入体力值: ")
			var newPower int32
			_, err = fmt.Scanf("%d ", &newPower)
			if err != nil {
				fmt.Println("体力值有误")
				return
			}
			store.Power = newPower
			player_store.SavePlayerStore(account.UIDStr, store)
			fmt.Println("体力值设置完成")
		case "5":
			fmt.Print("请输入点数: ")
			var newPoints int32
			_, err = fmt.Scanf("%d ", &newPoints)
			if err != nil {
				fmt.Println("点数有误")
				return
			}
			store.Points = newPoints
			player_store.SavePlayerStore(account.UIDStr, store)
			fmt.Println("点数设置完成")
		case "6":
			fmt.Print("请输入等级: ")
			var newLevel int32
			_, err = fmt.Scanf("%d ", &newLevel)
			if err != nil {
				fmt.Println("等级有误")
				return
			}
			store.Level = newLevel
			player_store.SavePlayerStore(account.UIDStr, store)
			fmt.Println("等级设置完成")
		case "7":
			fmt.Print("请输入经验值: ")
			var newExp int32
			_, err = fmt.Scanf("%d ", &newExp)
			if err != nil {
				fmt.Println("经验值有误")
				return
			}
			store.Exp = newExp
			player_store.SavePlayerStore(account.UIDStr, store)
			fmt.Println("经验值设置完成")
		default:
			return
		}
	}

}
