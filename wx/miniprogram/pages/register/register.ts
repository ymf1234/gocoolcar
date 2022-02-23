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
  profileRefresher: 0,
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
    this.clearProfileRefresher()
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
          const data = wx.getFileSystemManager().readFileSync(res.tempFilePaths[0])
          wx.request({
            method: 'PUT',
            url: 'https://coolcar-1253590403.cos.ap-shanghai.myqcloud.com',
            data,
            success: console.log,
            fail: console.error,
          })
        }
      }
    })
  },

  // 性别事件
  onGenderChange(e: any) {
    this.setData({
      gendersIndex: parseInt(e.detail.value),
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
    }).then(p => {
      this.renderProfile(p)
      this.scheduleProfileRefresher()
    })
  },

  scheduleProfileRefresher() {
    this.profileRefresher = setInterval(() => {
      ProfileService.getProfile().then(p => {
        this.renderProfile(p)
        if (p.identityStatus !== rental.v1.IdentityStatus.PENDING) {
          this.clearProfileRefresher()
        }

        if (p.identityStatus === rental.v1.IdentityStatus.VERIFIED) {
          this.onLicVerified()
        }
      })
    }, 1000)
  },


  clearProfileRefresher() {
    if (this.profileRefresher) {
      clearInterval(this.profileRefresher)
      this.profileRefresher = 0
    }
  },

  // 重新审核
  onResubmit() {
    ProfileService.clearProfile().then(p => this.renderProfile(p))
  },

  onLicVerified() {
    if(this.redirectURL) {
      wx.redirectTo({
        url: this.redirectURL
      })
    }
  }
})