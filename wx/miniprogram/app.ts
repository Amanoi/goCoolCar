import { getSetting, getUserInfo } from "./utils/util"
let resolveUserInfo: (value: WechatMiniprogram.UserInfo | PromiseLike<WechatMiniprogram.UserInfo>) => void
let rejectUserInfo: (reason?: any) => void

// app.ts
App<IAppOption>({
  globalData: {
    userInfo: new Promise((resolve, reject) => {
      resolveUserInfo = resolve
      rejectUserInfo = reject
    }),
  },
  async onLaunch() {
    // 登录
    wx.login({
      success: res => {
        console.log(res.code)
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      },
    })
    // async /wait
    try {
      const setting = await getSetting()
      if (setting.authSetting['scope.userInfo']) {
        const userInfoRes = await getUserInfo()
        resolveUserInfo(userInfoRes.userInfo)
      }
    } catch (err) {
      rejectUserInfo(err)
    }

  },
  resoveUserInfo(userInfo: WechatMiniprogram.UserInfo) {
    resolveUserInfo(userInfo)
  },
})