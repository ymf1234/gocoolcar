import { IAppOption } from "../../appoption"
import { ProfileService } from "../../service/profile";
import { rental } from "../../service/proto_gen/rental/rental_pb";
import { TripService } from "../../service/trip";
import { routing } from "../../utils/routing";

Page({
  avatarURL: '',
  isPageShowing: false,
  data: {
    showCancel: false,
    showModal: false,
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

  //  onLoad: function () { 
  //   // 页面渲染后 执行
  //   this.onHide()
  // },

  async onLoad() {
    // 页面渲染后 执行
    this.onHide()

    const userInfo = await getApp<IAppOption>().globalData.userInfo;
    if(userInfo) {
      this.setData({
        avatarURL: userInfo?.avatarUrl,
      })
    }

  },

  // 扫码
  async onScanTap() {
    const trips = await TripService.getTrips(rental.v1.TripStatus.IN_PROGRESS)
    if ((trips.trips?.length || 0) > 0) {
      await this.selectComponent('#tripModal').showModal()
      wx.navigateTo({
        url: routing.drving({
          trip_id: trips.trips![0].id!,
        }),
      })
      return
    }
    wx.scanCode({
      success: async () => {
        // TODO: get car id from scan result
        const carID='60af01e5a21ead3dccbcd1d8'
        const lockURL = routing.lock({
          car_id: carID,
        })
        const prof = await ProfileService.getProfile()
        if (prof.identityStatus === rental.v1.IdentityStatus.VERIFIED) {
          wx.navigateTo({
            url: lockURL,
          })
        } else {
          await this.selectComponent('#licModal').showModal()
          wx.navigateTo({
            url: routing.register({
              redirectURL: lockURL,
            })
          })
        }
      },
      fail: console.error,
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
  },

  onMyTripsTap() { 
    wx.navigateTo({
      url: routing.mytrips()
    })
  },

  onModalOK() {
    console.log('ok clicked')
    const carID='car123'
    // const redirectURL=`/pages/lock/lock?car_id=${carID}`
    const redirectURL = routing.lock({
      car_id: carID
    })
    wx.navigateTo({
      // url: `/pages/register/register?redirect=${encodeURIComponent(redirectURL)}`
      url: routing.register({redirectURL:redirectURL})
    })
  },
      
})
