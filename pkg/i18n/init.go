package i18n

// 全局i18n
var Def = NewI18n()

func init() {
	Def.SetLangValues("zh", map[string]string{})
	Def.SetLangValues("en", map[string]string{
		"请上传zip格式的文件":        "Please upload a file in zip format",
		"此小程序有正在审核的版本":       "This mini program has a version under review",
		"创建的小程序已达到上限":        "The maximum number of mini programs has been reached",
		"此版本已存在，请提升版本号":      "This version already exists, please increase the version number",
		"用户已被禁用":             "User has been disabled",
		"此用户ID不存在":           "This user ID does not exist",
		"无权限":                "No permission",
		"无法上架，因为没有审核的通过版本记录": "Unable to go online because there is no approved version record",
		"版本号格式不正确":           "Incorrect version number format",
		"注册新会员奖励":            "Rewards for Registering New Members",
		"邀请奖励":               "Invitation Rewards",
		"充值增加":               "Increase in Recharge Amount",
		"开发者分成增加":            "Increase in Developer's Share",
		"聊天消耗":               "Chatting Expenses",
		"使用小程序消耗":            "Expenses for Using Mini Programs",
		"未知类型":               "Unknown Type",
		"创作":                 "Creation",
		"图像":                 "Image",
	})
	Def.SetLangValues("jp", map[string]string{
		"请上传zip格式的文件":        "zip形式のファイルをアップロードしてください",
		"此小程序有正在审核的版本":       "この小プログラムには審査中のバージョンがあります",
		"创建的小程序已达到上限":        "作成した小プログラムの数が上限に達しました",
		"此版本已存在，请提升版本号":      "このバージョンは既に存在します。バージョン番号を上げてください",
		"用户已被禁用":             "ユーザーは無効になっています",
		"此用户ID不存在":           "このユーザーIDは存在しません",
		"无权限":                "権限がありません",
		"无法上架，因为没有审核的通过版本记录": "承認されたバージョン記録がないため、オンラインにすることができません",
		"版本号格式不正确":           "バージョン番号の形式が正しくありません",
		"注册新会员奖励":            "新規会員登録ボーナス",
		"邀请奖励":               "招待ボーナス",
		"充值增加":               "チャージ額増加",
		"开发者分成增加":            "開発者のシェア増加",
		"聊天消耗":               "チャット消費",
		"使用小程序消耗":            "ミニプログラム利用消費",
		"未知类型":               "未知の種類",
		"创作":                 "創造",
		"图像":                 "画像",
	})
}

// 通过全局的i18n获取翻译
func Get(lang, key string, args ...Options) string {
	return Def.Get(lang, key, args...)
}

func SetValue(lang, key, value string) {
	Def.SetValue(lang, key, value)
}

func SetLangValues(lang string, values map[string]string) {
	Def.SetLangValues(lang, values)
}
