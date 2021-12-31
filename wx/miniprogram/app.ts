import { IAppOption } from "./appoption"
import { coolcar } from "./service/proto_gen/trip_pb"
import { getSetting, getUserInfo } from "./utils/wxapi"

let resolveUserInfo: (value: WechatMiniprogram.UserInfo | PromiseLike<WechatMiniprogram.UserInfo>) => void
let rejectUserInfo: (reason?: any) => void

// app.ts
App<IAppOption>({
  globalData: {
    userInfo: new Promise((resolve, reject) => {
      resolveUserInfo = resolve
      rejectUserInfo = reject
    })
  },

  async onLaunch() {
    wx.request({
      url: "http://localhost:8080/trip/trip123",
      method: 'GET',
      success: res => {
        const getTripResp = coolcar.GetTripResponse.fromObject(res.data as object)
        console.log(getTripResp)
      },
      fail: console.error
    })
    // 登录
    wx.login({
      success: res => {
        
      }
    })
  },

  resolveUserInfo(userInfo: WechatMiniprogram.UserInfo) {
    resolveUserInfo(userInfo)
  }
})