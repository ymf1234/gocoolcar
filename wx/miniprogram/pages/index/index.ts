Page({
  isPageShowing: false,
  data: {
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: false,
      subKey: '',
      layerStyle: 1,
      enableZoom: true,
      enableScroll: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },
    is3D: false,
    isOverLooking: false,
    location: {
      latitude: 31,
      longitude: 130
    },
    scale: 10,
    markers: [
      {
        iconPath: "/resources/car.png",
        id: 0,
        latitude: 23.099994,
        longitude: 113.342520,
        width: 50,
        height: 50
      },
      {
        iconPath: "/resources/car.png",
        id: 1,
        latitude: 23.099994,
        longitude: 114.342520,
        width: 50,
        height: 50
      },
    ],
    
  },

  onLoad: function () { 
    // 页面渲染后 执行
    this.onHide()
  },

  // 扫码
  onScanClicked() {
    wx.scanCode({
      success: () => {
        wx.navigateTo({
          url: '/pages/register/register'
        })
      },
      fail: res => {
        console.log('res', res)
      }
    })
  },

  // 显示
  onShow() {
    this.isPageShowing = true
  },

  // 不显示 
  onHide() {
    this.isPageShowing = false
  },

  // 获取当前位置
  onMyLocationTap() {
    wx.getLocation({
      type: 'gcj02',
      success: res => {
        console.log("latitude", res.latitude)
        console.log("longitude", res.longitude)
        this.setData({
          location: {
            latitude: res.latitude,
            longitude: res.longitude
          }
        })
      },
      fail: () => {
        wx.showToast({
          icon: "none",
          title: '请前往设置页面授权'
        })
      }
    })
  },


  // 移动
  moveCars() {
    console.log("移动")
    const map = wx.createMapContext("map");
    const dest = {
        latitude: 23.099994,
        longitude: 113.342520,
    }
    const moveCar = () => {
      dest.latitude += 0.1
      dest.longitude += 0.1
      map.translateMarker({
        destination: { 
          latitude: dest.latitude,
          longitude: dest.longitude,
        },
        markerId: 0,
        autoRotate: false,
        rotate: 0,
        duration: 5000,
        animationEnd: () => {
          if (this.isPageShowing) {
            moveCar()
          }
        }
      })
    }

    moveCar()
  }
      
})
