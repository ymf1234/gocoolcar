import camelcaseKeys from "camelcase-keys"
import { IAppOption } from "./appoption"
import { auth } from "./service/proto_gen/auth/auth_pb"

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
    // wx.request({
    //   url: "http://localhost:8080/trip/trip123",
    //   method: 'GET',
    //   success: res => {
    //     const getTripResp = coolcar.GetTripResponse.fromObject(
    //       camelcaseKeys(res.data as object, {
    //         deep: true
    //       }))
    //     console.log(getTripResp)
    //   },
    //   fail: console.error
    // })
    // 登录
    wx.login({
      success: res => {
        console.log(res)
        wx.request({
          url: 'http://localhost:8080/v1/auth/login',
          method: 'POST',
          data: {
            code: res.code,
          } as auth.v1.ILoginRequest,
          success: res => {
            const loginResp: auth.v1.ILoginResponse = auth.v1.LoginResponse.fromObject(
              camelcaseKeys(res.data as object),
            )
              
          
            console.log(loginResp)
            wx.request({
              url: 'http://localhost:8080/v1/trip',
              method: 'POST',
              data: {
                start: 'abc',
              },
              header: {
                authorization: 'Bearer ' + loginResp.accessToken
              }
            })
          },
          fail: console.error,
        })
      }
    })
  },

  resolveUserInfo(userInfo: WechatMiniprogram.UserInfo) {
    resolveUserInfo(userInfo)
  }
})