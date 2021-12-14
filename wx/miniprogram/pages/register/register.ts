// pages/register/register.ts
Page({

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
    state: 'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED'
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {

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
    this.setData({
      state: 'PENDING'
    })
    setTimeout(() => {
      this.onLicVerified()
    }, 3000);
  },

  // 重新审核
  onResubmit() {
    this.setData({
      state: 'UNSUBMITTED',
      licImgURL: ''
    })
  },

  onLicVerified() {
    this.setData({
      state: 'VERIFIED',
    })
    wx.redirectTo({
      url: '/pages/lock/lock'
    })
  }
})