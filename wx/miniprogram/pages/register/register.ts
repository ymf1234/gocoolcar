import { ProfileService } from "../../service/profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { padString } from "../../utils/format"
import { routing } from "../../utils/routing"

function formatTime(millis: number) {
  const dt = new Date(millis)
  const y = dt.getFullYear()
  const m = dt.getMonth() + 1
  const d = dt.getDate()

  return `${padString(y)}-${padString(m)}-${padString(d)}`
}
// pages/register/register.ts
Page({
  redirectURL: '',
  /**
   * 页面的初始数据
   */
  data: {
    licImgURL: '',
    genders: ['未知', '男', '女', '其他'],
    gendersIndex: 0,
    birthDate: '1990-01-01',
    licNo: '',
    name: '',
    state: rental.v1.IdentityStatus[rental.v1.IdentityStatus.UNSUBMITTED]
  },

  renderProfile(p: rental.v1.IProfile) {
    this.setData({
      licNo: p.identity?.licNumber || '',
      name: p.identity?.name || '',
      gendersIndex: p.identity?.gender || 0,
      birthDate: formatTime(p.identity?.birthDateMillis || 0),
      state: rental.v1.IdentityStatus[p.identityStatus || 0], 
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt: Record<'redirect', string>) {
    const o: routing.RegisterOpts = opt
    if(o.redirect) {
      this.redirectURL = decodeURIComponent(o.redirect)
    }

  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  },

  onUploadLic() {
    wx.chooseImage({
      success: res => {
        if(res.tempFilePaths.length > 0 ) {
          this.setData({
            licImgURL: res.tempFilePaths[0]
          })
          // TODO: upload image
          setTimeout(() => {
            this.setData({
              licNo: '12312312312',
              name : 'aaa',
              gendersIndex: 1,
              birthDate: '2000-12-12'
            })
          }, 1000)
        }
      }
    })
  },

  // 性别事件
  onGenderChange(e: any) {
    this.setData({
      gendersIndex: e.detail.value,
    })
  },

  // 时间
  onBirthDateChange(e: any) {
    this.setData({
      birthDate: e.detail.value,
    })
  },

  // 表单提交
  onSubmit() {
    ProfileService.submitProfile({
      licNumber: this.data.licNo,
      name: this.data.name,
      gender: this.data.gendersIndex,
      birthDateMillis: Date.parse(this.data.birthDate)
    }).then(p => this.renderProfile(p))
  },

  // 重新审核
  onResubmit() {
    ProfileService.clearProfile().then(p => this.renderProfile(p))
  },

  onLicVerified() {
    this.setData({
      state: 'VERIFIED',
    })

    if(this.redirectURL) {
      wx.redirectTo({
        url: this.redirectURL
      })
    }
  }
})