import { routing } from "../../utils/routing"

const centPerSec = 1

function formatDuration(sec: number) {
    const padString = (n: number) => 
        n < 10 ? '0'+n.toFixed(0) : n.toFixed(0)
    
    const h = Math.floor(sec / 3600)
    sec -= 3600 * h
    const m = Math.floor(sec / 60)
    sec -= 60 * m
    const s = Math.floor(sec)
    // return `${padString(h)}:${padString(m)}:${padString(s)}`
    return `${padString(h)}:${padString(m)}:${padString(s)}`
}

function formatFee(cents: number) {
    return (cents/100).toFixed(2)
}

// pages/driving/driving.ts
Page({

    timer: undefined as number | undefined,
    /**
     * 页面的初始数据
     */
    data: {
        location: {
            latitude: 0,
            longitude: 0,
        },
        scale: 12,
        elapsed: '00:00:00',
        fee: '0.00',
        markers: [
            {
                iconPath: "/resources/car.png",
                id: 0,
                latitude: 0,
                longitude: 0,
                width: 20,
                height: 20,
            },
        ],
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(opt: Record<'trip_id', string>) {
        const o : routing.DrivingOpts = opt
        console.log('current trip', o.trip_id)
        this.setupLocationUpdator()
        this.setupTimer()
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
        wx.stopLocationUpdate()
        if(this.timer) {
            clearInterval(this.timer)
        }
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

    setupLocationUpdator() {
        wx.startLocationUpdate({
            fail: console.error
        })
        wx.onLocationChange(loc => {
            this.setData({
                location: {
                    latitude: loc.latitude,
                    longitude: loc.longitude,
                }
            })
        })
        
    },

    setupTimer() {
        let elapsedSec = 0
        let cents = 0
        this.timer = setInterval(() => {
            elapsedSec++
            cents++
            this.setData({
                elapsed: formatDuration(elapsedSec),
                fee: formatFee(cents),
            })
        }, 1000)
    },

    onEndTripTap() {
        wx.redirectTo({
            url: routing.mytrips()
        })
    }
})